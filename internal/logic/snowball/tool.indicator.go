package snowball

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"fi.mcp/internal/types"
	"github.com/zeromicro/go-zero/mcp"
)

// IndicatorArgs 定义 indicator 工具的输入参数
type IndicatorArgs struct {
	Symbol string `json:"symbol" jsonschema:"股票代码"`
	Count  string `json:"count,omitempty" jsonschema:"每次获取数量，默认值:6"`
}

var indicatorDescription = map[string]string{
	"report_date":                    "报告日期",
	"report_name":                    "报告名称",
	"ctime":                          "创建时间",
	"avg_roe":                        "平均ROE",
	"np_per_share":                   "每股净利润",
	"operate_cash_flow_ps":           "每股经营现金流",
	"basic_eps":                      "基本每股收益",
	"capital_reserve":                "资本公积",
	"undistri_profit_ps":             "每股未分配利润",
	"net_interest_of_total_assets":   "总资产净利率",
	"net_selling_rate":               "净销售率",
	"gross_selling_rate":             "毛销售率",
	"total_revenue":                  "总营收",
	"operating_income_yoy":           "营业收入同比",
	"net_profit_atsopc":              "归母净利润",
	"net_profit_atsopc_yoy":          "归母净利润同比",
	"net_profit_after_nrgal_atsolc":  "扣非归母净利润",
	"np_atsopc_nrgal_yoy":            "扣非归母净利润同比",
	"ore_dlt":                        "营运资金变动",
	"rop":                            "营业利润",
	"asset_liab_ratio":               "资产负债率",
	"current_ratio":                  "流动比率",
	"quick_ratio":                    "速动比率",
	"equity_multiplier":              "权益乘数",
	"equity_ratio":                   "权益比率",
	"holder_equity":                  "股东权益",
	"ncf_from_oa_to_total_liab":      "经营活动现金流/总负债",
	"inventory_turnover_days":        "存货周转天数",
	"receivable_turnover_days":       "应收账款周转天数",
	"accounts_payable_turnover_days": "应付账款周转天数",
	"cash_cycle":                     "现金周期",
	"operating_cycle":                "营运周期",
	"total_capital_turnover":         "总资产周转率",
	"inventory_turnover":             "存货周转率",
	"account_receivable_turnover":    "应收账款周转率",
	"accounts_payable_turnover":      "应付账款周转率",
	"current_asset_turnover_rate":    "流动资产周转率",
	"fixed_asset_turnover_ratio":     "固定资产周转率",
	"amount":                         "数值",
	"growth_rate":                    "增长率",
}

func NewIndicatorTool(_mcp types.MCPProvider) (*mcp.Tool, func(context.Context, *mcp.CallToolRequest, IndicatorArgs) (*mcp.CallToolResult, any, error)) {
	tool := &mcp.Tool{
		Name:        "indicator",
		Description: "获取公司主要财务指标",
	}

	handler := func(ctx context.Context, req *mcp.CallToolRequest, args IndicatorArgs) (*mcp.CallToolResult, any, error) {
		if args.Count == "" {
			args.Count = "6"
		}

		indicator := types.Indicator{}
		var items []map[string]interface{}

		url := fmt.Sprintf(_mcp.GetServiceContext().Config.DataSource.Snowball.IndicatorURL, args.Symbol, args.Count, time.Now().UnixMilli())
		_mcp.GetLogger().Infof("url: %s", url)

		client := NewClientWithConfig(&_mcp.GetServiceContext().Config)
		setHeader(_mcp.GetServiceContext().Config.DataSource.UserAgent, _mcp.GetServiceContext().Config.DataSource.Snowball.IndexURL, _mcp.GetServiceContext().Config.DataSource.Snowball.CookieURL, &_mcp.GetServiceContext().Config, client)
		_, err := client.R().SetResult(&indicator).Get(url)

		if err != nil {
			return nil, nil, err
		}

		for _, item := range indicator.Data.List {
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

		_mcp.GetLogger().Infof("data item count: %d", len(indicator.Data.List))

		result := map[string]any{
			"columns": indicatorDescription,
			"items":   items,
		}

		jsonBytes, err := json.Marshal(result)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to marshal result: %w", err)
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: string(jsonBytes)},
			},
		}, nil, nil
	}

	return tool, handler
}