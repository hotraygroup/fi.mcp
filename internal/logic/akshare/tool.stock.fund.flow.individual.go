package akshare

import (
	"context"
	"fmt"

	"fi.mcp/internal/types"
	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/mcp"
)

// stock_fund_flow_individual
func NewStockFundFlowIndividualTool(_mcp types.MCPProvider) mcp.Tool {
	var tool = mcp.Tool{
		Name:        "stock_fund_flow_individual",
		Description: "同花顺-数据中心-资金流向-个股资金流",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"symbol": map[string]any{
					"type":        "string",
					"description": "类型",
					"default":     "即时",
					"enum": []string{
						"即时", "3日排行", "5日排行", "10日排行", "20日排行",
					},
				},
			},
			Required: []string{"symbol"},
		},
		Handler: func(ctx context.Context, params map[string]any) (any, error) {
			var req struct {
				Symbol string `json:"symbol,omitempty"`
			}

			if err := mcp.ParseArguments(params, &req); err != nil {
				return nil, fmt.Errorf("failed to parse params: %w", err)
			}

			client := resty.New()
			url := fmt.Sprintf("%s/api/public/stock_fund_flow_individual?symbol=%s",
				_mcp.GetServiceContext().Config.DataSource.Akshare.Host,
				req.Symbol,
			)

			_mcp.GetLogger().Infof("url: %s", url)

			var result []map[string]any

			_, err := client.R().SetResult(&result).Get(url)
			if err != nil {
				return nil, err
			}
			return map[string]any{
				"items": result,
			}, nil
		},
	}
	return tool
}
