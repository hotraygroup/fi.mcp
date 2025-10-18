package akshare

import (
	"context"
	"fmt"
	"strconv"

	"fi.mcp/internal/types"
	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/mcp"
)

// stock_financial_cash_ths
func NewStockFinancialCashThsTool(_mcp types.MCPProvider) mcp.Tool {
	var tool = mcp.Tool{
		Name:        "stock_financial_cash_ths",
		Description: "同花顺-财务指标-现金流量表",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"symbol": map[string]any{
					"type":        "string",
					"description": "股票代码, 如600001",
				},
				"count": map[string]any{
					"type":        "string",
					"description": "获取数据的数量",
					"default":     "12",
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
				req.Count = "12"
			}

			count, err := strconv.Atoi(req.Count)
			if err != nil || count <= 0 {
				count = 12
			}

			client := resty.New()
			url := fmt.Sprintf("%s/api/public/stock_financial_cash_ths?symbol=%s",
				_mcp.GetServiceContext().Config.DataSource.Akshare.Host,
				req.Symbol,
			)

			_mcp.GetLogger().Infof("url: %s", url)

			var result []map[string]any

			_, err = client.R().SetResult(&result).Get(url)
			if err != nil {
				return nil, err
			}

			if len(result) > count {
				result = result[:count]
			}

			return map[string]any{
				"items": result,
			}, nil
		},
	}
	return tool
}
