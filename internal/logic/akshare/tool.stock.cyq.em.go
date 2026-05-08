package akshare

import (
	"context"
	"encoding/json"
	"fmt"

	"fi.mcp/internal/types"
	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/mcp"
)

// StockCyqEmArgs 定义 stock_cyq_em 工具的输入参数
type StockCyqEmArgs struct {
	Symbol string `json:"symbol" jsonschema:"股票代码, 如600001"`
}

func NewStockCyqEmTool(_mcp types.MCPProvider) (*mcp.Tool, func(context.Context, *mcp.CallToolRequest, StockCyqEmArgs) (*mcp.CallToolResult, any, error)) {
	tool := &mcp.Tool{
		Name:        "stock_cyq_em",
		Description: "东方财富网-行情中心-日K-筹码分布",
	}

	handler := func(ctx context.Context, req *mcp.CallToolRequest, args StockCyqEmArgs) (*mcp.CallToolResult, any, error) {
		client := resty.New()
		url := fmt.Sprintf("%s/api/public/stock_cyq_em?symbol=%s&adjust=qfq",
			_mcp.GetServiceContext().Config.DataSource.Akshare.Host,
			args.Symbol,
		)

		_mcp.GetLogger().Infof("url: %s", url)

		var result []map[string]any

		_, err := client.R().SetResult(&result).Get(url)
		if err != nil {
			return nil, nil, err
		}

		jsonBytes, err := json.Marshal(map[string]any{"items": result})
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