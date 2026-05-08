package akshare

import (
	"context"
	"encoding/json"
	"fmt"

	"fi.mcp/internal/types"
	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/mcp"
)

// StockIndividualFundFlowArgs 定义 stock_individual_fund_flow 工具的输入参数
type StockIndividualFundFlowArgs struct {
	Stock string `json:"stock" jsonschema:"股票代码, 如600001"`
}

func NewStockIndividualFundFlowTool(_mcp types.MCPProvider) (*mcp.Tool, func(context.Context, *mcp.CallToolRequest, StockIndividualFundFlowArgs) (*mcp.CallToolResult, any, error)) {
	tool := &mcp.Tool{
		Name:        "stock_individual_fund_flow",
		Description: "东方财富网-数据中心-个股资金流向",
	}

	handler := func(ctx context.Context, req *mcp.CallToolRequest, args StockIndividualFundFlowArgs) (*mcp.CallToolResult, any, error) {
		client := resty.New()
		url := fmt.Sprintf("%s/api/public/stock_individual_fund_flow?stock=%s",
			_mcp.GetServiceContext().Config.DataSource.Akshare.Host,
			args.Stock,
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