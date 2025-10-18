package akshare

import (
	"context"
	"fmt"

	"fi.mcp/internal/types"
	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/mcp"
)

// stock_individual_fund_flow
func NewStockIndividualFundFlowTool(_mcp types.MCPProvider) mcp.Tool {
	var tool = mcp.Tool{
		Name:        "stock_individual_fund_flow",
		Description: "东方财富网-数据中心-个股资金流向",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"stock": map[string]any{
					"type":        "string",
					"description": "股票代码, 如600001",
				},
			},
			Required: []string{"stock"},
		},
		Handler: func(ctx context.Context, params map[string]any) (any, error) {
			var req struct {
				Stock string `json:"stock"`
			}

			if err := mcp.ParseArguments(params, &req); err != nil {
				return nil, fmt.Errorf("failed to parse params: %w", err)
			}

			client := resty.New()
			url := fmt.Sprintf("%s/api/public/stock_individual_fund_flow?stock=%s",
				_mcp.GetServiceContext().Config.DataSource.Akshare.Host,
				req.Stock,
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
