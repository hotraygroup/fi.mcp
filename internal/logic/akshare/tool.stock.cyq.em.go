package akshare

import (
	"context"
	"fmt"

	"fi.mcp/internal/types"
	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/mcp"
)

// stock_cyq_em
func NewStockCyqEmTool(_mcp types.MCPProvider) mcp.Tool {
	var tool = mcp.Tool{
		Name:        "stock_cyq_em",
		Description: "东方财富网-行情中心-日K-筹码分布",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"symbol": map[string]any{
					"type":        "string",
					"description": "股票代码, 如600001",
				},
			},
			Required: []string{"symbol"},
		},
		Handler: func(ctx context.Context, params map[string]any) (any, error) {
			var req struct {
				Symbol string `json:"symbol"`
			}

			if err := mcp.ParseArguments(params, &req); err != nil {
				return nil, fmt.Errorf("failed to parse params: %w", err)
			}

			client := resty.New()
			url := fmt.Sprintf("%s/api/public/stock_cyq_em?symbol=%s&adjust=qfq",
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
