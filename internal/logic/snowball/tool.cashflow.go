package snowball

import (
	"context"
	"fmt"
	"time"

	"fi.mcp/internal/types"
	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/mcp"
)

var cashFlowDescription = map[string]string{
	"report_date":                    "报告日期",
	"report_name":                    "报告名称",
	"ctime":                          "创建时间",
	"ncf_from_oa":                    "经营活动产生的现金流量净额",
	"ncf_from_ia":                    "投资活动产生的现金流量净额",
	"ncf_from_fa":                    "筹资活动产生的现金流量净额",
	"cash_received_of_othr_oa":       "收到其他与经营活动有关的现金",
	"sub_total_of_ci_from_oa":        "经营活动现金流入小计",
	"cash_paid_to_employee_etc":      "支付给职工以及为职工支付的现金",
	"payments_of_all_taxes":          "支付的各项税费",
	"othrcash_paid_relating_to_oa":   "支付其他与经营活动有关的现金",
	"sub_total_of_cos_from_oa":       "经营活动现金流出小计",
	"cash_received_of_dspsl_invest":  "收回投资收到的现金",
	"invest_income_cash_received":    "取得投资收益收到的现金",
	"net_cash_of_disposal_assets":    "处置固定资产、无形资产和其他长期资产收回的现金净额",
	"net_cash_of_disposal_branch":    "处置子公司及其他营业单位收到的现金净额",
	"cash_received_of_othr_ia":       "收到其他与投资活动有关的现金",
	"sub_total_of_ci_from_ia":        "投资活动现金流入小计",
	"invest_paid_cash":               "投资支付的现金",
	"cash_paid_for_assets":           "购建固定资产、无形资产和其他长期资产支付的现金",
	"othrcash_paid_relating_to_ia":   "支付其他与投资活动有关的现金",
	"sub_total_of_cos_from_ia":       "投资活动现金流出小计",
	"cash_received_of_absorb_invest": "吸收投资收到的现金",
	"cash_received_from_investor":    "公司吸收少数股东投资收到的现金",
	"cash_received_from_bond_issue":  "发行债券收到的现金",
	"cash_received_of_borrowing":     "取得借款收到的现金",
	"cash_received_of_othr_fa":       "收到其他与筹资活动有关的现金",
	"sub_total_of_ci_from_fa":        "筹资活动现金流入小计",
	"cash_pay_for_debt":              "偿还债务支付的现金",
	"cash_paid_of_distribution":      "分配股利、利润或偿付利息支付的现金",
	"branch_paid_to_minority_holder": "公司支付给少数股东的股利、利润",
	"othrcash_paid_relating_to_fa":   "支付其他与筹资活动有关的现金",
	"sub_total_of_cos_from_fa":       "筹资活动现金流出小计",
	"effect_of_exchange_chg_on_cce":  "汇率变动对现金及现金等价物的影响",
	"net_increase_in_cce":            "现金及现金等价物净增加额",
	"initial_balance_of_cce":         "期初现金及现金等价物余额",
	"final_balance_of_cce":           "期末现金及现金等价物余额",
	"cash_received_of_sales_service": "销售商品、提供劳务收到的现金",
	"refund_of_tax_and_levies":       "收到的税费返还",
	"goods_buy_and_service_cash_pay": "购买商品、接受劳务支付的现金",
	"net_cash_amt_from_branch":       "取得子公司及其他营业单位支付的现金净额",
}

func NewCashFlowTool(_mcp types.MCPProvider) mcp.Tool {

	var cashFlowTool = mcp.Tool{
		Name:        "cashflow",
		Description: "获取公司现金流量表",
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

			cashFlow := types.CashFlow{}
			var items []map[string]interface{}

			url := fmt.Sprintf(_mcp.GetServiceContext().Config.DataSource.Snowball.CashFlowURL, req.Symbol, req.Count, time.Now().UnixMilli())
			_mcp.GetLogger().Infof("url: %s", url)

			client := resty.New()
			setHeader(_mcp.GetServiceContext().Config.DataSource.UserAgent, _mcp.GetServiceContext().Config.DataSource.Snowball.IndexURL, _mcp.GetServiceContext().Config.DataSource.Snowball.CookieURL, client)
			_, err := client.R().SetResult(&cashFlow).Get(url)

			if err != nil {
				return nil, err
			}

			for _, item := range cashFlow.Data.List {
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

			_mcp.GetLogger().Infof("data item count: %d", len(cashFlow.Data.List))
			return map[string]any{
				"columns": cashFlowDescription,
				"items":   items,
			}, nil

		},
	}
	return cashFlowTool
}
