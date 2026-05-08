package akshare

import (
	"context"
	"encoding/json"
	"fmt"

	"fi.mcp/internal/types"
	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/mcp"
)

// StockFundFlowBigDealArgs 定义 stock_fund_flow_big_deal 工具的输入参数（无参数）
type StockFundFlowBigDealArgs struct{}

func NewStockFundFlowBigDealTool(_mcp types.MCPProvider) (*mcp.Tool, func(context.Context, *mcp.CallToolRequest, StockFundFlowBigDealArgs) (*mcp.CallToolResult, any, error)) {
	tool := &mcp.Tool{
		Name:        "stock_fund_flow_big_deal",
		Description: "同花顺-数据中心-资金流向-大单追踪",
	}

	handler := func(ctx context.Context, req *mcp.CallToolRequest, args StockFundFlowBigDealArgs) (*mcp.CallToolResult, any, error) {
		client := resty.New()
		url := fmt.Sprintf("%s/api/public/stock_fund_flow_big_deal",
			_mcp.GetServiceContext().Config.DataSource.Akshare.Host,
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