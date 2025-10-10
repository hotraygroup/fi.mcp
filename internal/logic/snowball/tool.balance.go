package snowball

import (
	"context"
	"fmt"
	"time"

	"fi.mcp/internal/types"
	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/mcp"
)

var balanceDescription = map[string]string{
	"report_date":                   "报告日期",
	"report_name":                   "报告名称",
	"ctime":                         "创建时间",
	"total_assets":                  "总资产",
	"total_liab":                    "总负债",
	"asset_liab_ratio":              "资产负债率",
	"total_quity_atsopc":            "归母股东权益",
	"tradable_fnncl_assets":         "交易性金融资产",
	"interest_receivable":           "应收利息",
	"saleable_finacial_assets":      "可供出售金融资产",
	"held_to_maturity_invest":       "持有至到期投资",
	"fixed_asset":                   "固定资产",
	"intangible_assets":             "无形资产",
	"construction_in_process":       "在建工程",
	"dt_assets":                     "递延所得税资产",
	"tradable_fnncl_liab":           "交易性金融负债",
	"payroll_payable":               "应付职工薪酬",
	"tax_payable":                   "应交税费",
	"estimated_liab":                "预计负债",
	"dt_liab":                       "递延所得税负债",
	"bond_payable":                  "应付债券",
	"shares":                        "股本",
	"capital_reserve":               "资本公积",
	"earned_surplus":                "盈余公积",
	"undstrbtd_profit":              "未分配利润",
	"minority_equity":               "少数股东权益",
	"total_holders_equity":          "股东权益合计",
	"total_liab_and_holders_equity": "负债和股东权益总计",
	"lt_equity_invest":              "长期股权投资",
	"derivative_fnncl_liab":         "衍生金融负债",
	"general_risk_provision":        "一般风险准备",
	"frgn_currency_convert_diff":    "外币报表折算差额",
	"goodwill":                      "商誉",
	"invest_property":               "投资性房地产",
	"interest_payable":              "应付利息",
	"treasury_stock":                "库存股",
	"othr_compre_income":            "其他综合收益",
	"othr_equity_instruments":       "其他权益工具",
	"currency_funds":                "货币资金",
	"bills_receivable":              "应收票据",
	"account_receivable":            "应收账款",
	"pre_payment":                   "预付款项",
	"dividend_receivable":           "应收股利",
	"othr_receivables":              "其他应收款",
	"inventory":                     "存货",
	"nca_due_within_one_year":       "一年内到期的非流动资产",
	"othr_current_assets":           "其他流动资产",
	"current_assets_si":             "流动资产特殊项目",
	"total_current_assets":          "流动资产合计",
	"lt_receivable":                 "长期应收款",
	"dev_expenditure":               "开发支出",
	"lt_deferred_expense":           "长期待摊费用",
	"othr_noncurrent_assets":        "其他非流动资产",
	"noncurrent_assets_si":          "非流动资产特殊项目",
	"total_noncurrent_assets":       "非流动资产合计",
	"st_loan":                       "短期借款",
	"bill_payable":                  "应付票据",
	"accounts_payable":              "应付账款",
	"pre_receivable":                "预收款项",
	"dividend_payable":              "应付股利",
	"othr_payables":                 "其他应付款",
	"noncurrent_liab_due_in1y":      "一年内到期的非流动负债",
	"current_liab_si":               "流动负债特殊项目",
	"total_current_liab":            "流动负债合计",
	"lt_loan":                       "长期借款",
	"lt_payable":                    "长期应付款",
	"special_payable":               "专项应付款",
	"othr_non_current_liab":         "其他非流动负债",
	"noncurrent_liab_si":            "非流动负债特殊项目",
	"total_noncurrent_liab":         "非流动负债合计",
	"salable_financial_assets":      "可供出售金融资产",
	"othr_current_liab":             "其他流动负债",
	"ar_and_br":                     "应收票据及应收账款",
	"contractual_assets":            "合同资产",
	"bp_and_ap":                     "应付票据及应付账款",
	"contract_liabilities":          "合同负债",
	"to_sale_asset":                 "持有待售资产",
	"other_eq_ins_invest":           "其他权益工具投资",
	"other_illiquid_fnncl_assets":   "其他非流动金融资产",
	"fixed_asset_sum":               "固定资产合计",
	"fixed_assets_disposal":         "固定资产清理",
	"construction_in_process_sum":   "在建工程合计",
	"project_goods_and_material":    "工程物资",
	"productive_biological_assets":  "生产性生物资产",
	"oil_and_gas_asset":             "油气资产",
	"to_sale_debt":                  "持有待售负债",
	"lt_payable_sum":                "长期应付款合计",
	"noncurrent_liab_di":            "非流动负债调整项目",
	"perpetual_bond":                "永续债",
	"special_reserve":               "专项储备",
	"amount":                        "数值",
	"growth_rate":                   "增长率",
}

func NewBalanceTool(_mcp types.MCPProvider) mcp.Tool {

	var balanceTool = mcp.Tool{
		Name:        "balance",
		Description: "获取公司资产负债表",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"symbol": map[string]any{
					"type":        "string",
					"description": "股票代码",
				},
				"count": map[string]any{
					"type":        "string",
					"description": "每次获取数量",
					"default":     "6",
				},
			},
			Required: []string{"symbol"},
		},
		Handler: func(ctx context.Context, params map[string]any) (any, error) {
			var req struct {
				Symbol string `json:"symbol"`
				Count  string `json:"count,omitempty"`
			}

			if err := mcp.ParseArguments(params, &req); err != nil {
				return nil, fmt.Errorf("failed to parse params: %w", err)
			}

			if req.Count == "" {
				req.Count = "6"
			}

			balance := types.Balance{}
			var items []map[string]interface{}

			url := fmt.Sprintf(_mcp.GetServiceContext().Config.DataSource.Snowball.BalanceURL, req.Symbol, req.Count, time.Now().UnixMilli())
			_mcp.GetLogger().Infof("url: %s", url)

			client := resty.New()
			setHeader(_mcp.GetServiceContext().Config.DataSource.UserAgent, _mcp.GetServiceContext().Config.DataSource.Snowball.IndexURL, _mcp.GetServiceContext().Config.DataSource.Snowball.CookieURL, client)
			_, err := client.R().SetResult(&balance).Get(url)

			if err != nil {
				return nil, err
			}

			for _, item := range balance.Data.List {
				new := make(map[string]interface{})
				for k, v := range item {
					if v == nil {
						continue
					}
					switch v := v.(type) {
					case float64, string, int64:
						new[k] = v
					case []interface{}:
						if len(v) == 0 || v[0] == nil {
							continue
						}
						if len(v) > 1 {
							if v[0] == nil || v[1] == nil {
								continue
							}
							new[k] = struct {
								Amount     float64 `json:"amount"`
								GrowthRate float64 `json:"growth_rate"`
							}{
								Amount:     v[0].(float64),
								GrowthRate: v[1].(float64),
							}
						} else if len(v) == 1 {
							new[k] = v[0]
						}
					default:
						_mcp.GetLogger().Infof("unknown type: %T, value: %v", v, v)
					}
				}
				items = append(items, new)
			}

			_mcp.GetLogger().Infof("data item count: %d", len(balance.Data.List))
			return map[string]any{
				"columns": balanceDescription,
				"items":   items,
			}, nil

		},
	}
	return balanceTool
}
