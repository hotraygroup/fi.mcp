package logic

import (
	"context"
	"fmt"
	"time"

	"fi.mcp/internal/types"
	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/mcp"
)

var incomeDescription = map[string]string{
	"report_date":                    "报告日期",
	"report_name":                    "报告名称",
	"ctime":                          "创建时间",
	"net_profit":                     "净利润",
	"net_profit_atsopc":              "归母净利润",
	"total_revenue":                  "总营收",
	"op":                             "营业利润",
	"income_from_chg_in_fv":          "公允价值变动收益",
	"invest_incomes_from_rr":         "投资收益",
	"invest_income":                  "投资收益",
	"exchg_gain":                     "汇兑收益",
	"operating_taxes_and_surcharge":  "营业税金及附加",
	"asset_impairment_loss":          "资产减值损失",
	"non_operating_income":           "营业外收入",
	"non_operating_payout":           "营业外支出",
	"profit_total_amt":               "利润总额",
	"minority_gal":                   "少数股东损益",
	"basic_eps":                      "基本每股收益",
	"dlt_earnings_per_share":         "稀释每股收益",
	"othr_compre_income_atoopc":      "归母其他综合收益",
	"othr_compre_income_atms":        "少数股东其他综合收益",
	"total_compre_income":            "综合收益总额",
	"total_compre_income_atsopc":     "归母综合收益总额",
	"total_compre_income_atms":       "少数股东综合收益总额",
	"othr_compre_income":             "其他综合收益",
	"net_profit_after_nrgal_atsolc":  "扣非归母净利润",
	"income_tax_expenses":            "所得税费用",
	"credit_impairment_loss":         "信用减值损失",
	"revenue":                        "营业收入",
	"operating_costs":                "营业成本",
	"operating_cost":                 "营业成本",
	"sales_fee":                      "销售费用",
	"manage_fee":                     "管理费用",
	"financing_expenses":             "财务费用",
	"rad_cost":                       "研发费用",
	"finance_cost_interest_fee":      "财务费用-利息支出",
	"finance_cost_interest_income":   "财务费用-利息收入",
	"asset_disposal_income":          "资产处置收益",
	"other_income":                   "其他收益",
	"noncurrent_assets_dispose_gain": "非流动资产处置利得",
	"noncurrent_asset_disposal_loss": "非流动资产处置损失",
	"net_profit_bi":                  "净利润(含少数股东损益)",
	"continous_operating_np":         "持续经营净利润",
	"amount":                         "数值",
	"growth_rate":                    "增长率",
}

func newIncomeTool(_mcp *MCP) mcp.Tool {

	var incomeTool = mcp.Tool{
		Name:        "income",
		Description: "获取公司利润表",
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
				Count  string `json:"count"`
			}

			if err := mcp.ParseArguments(params, &req); err != nil {
				return nil, fmt.Errorf("failed to parse params: %w", err)
			}

			income := types.Income{}
			var items []map[string]interface{}

			url := fmt.Sprintf(_mcp.svcCtx.Config.DataSource.Snowball.IncomeURL, req.Symbol, req.Count, time.Now().UnixMilli())
			_mcp.Infof("url: %s", url)

			client := resty.New()
			setHeader(_mcp.svcCtx.Config.DataSource.UserAgent, _mcp.svcCtx.Config.DataSource.Snowball.IndexURL, client)
			_, err := client.R().SetResult(&income).Get(url)

			if err != nil {
				return nil, err
			}

			for _, item := range income.Data.List {
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
						_mcp.Infof("unknown type: %T, value: %v", v, v)
					}
				}
				items = append(items, new)
			}

			_mcp.Infof("data item count: %d", len(income.Data.List))
			return map[string]any{
				"columns": incomeDescription,
				"items":   items,
			}, nil

		},
	}
	return incomeTool
}
