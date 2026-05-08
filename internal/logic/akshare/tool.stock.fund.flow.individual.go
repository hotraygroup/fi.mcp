package akshare

import (
	"context"
	"encoding/json"
	"fmt"

	"fi.mcp/internal/types"
	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/mcp"
)

// StockFundFlowIndividualArgs 定义 stock_fund_flow_individual 工具的输入参数
type StockFundFlowIndividualArgs struct {
	Symbol string `json:"symbol" jsonschema:"类型（即时,3日排行,5日排行,10日排行,20日排行）"`
}

func NewStockFundFlowIndividualTool(_mcp types.MCPProvider) (*mcp.Tool, func(context.Context, *mcp.CallToolRequest, StockFundFlowIndividualArgs) (*mcp.CallToolResult, any, error)) {
	tool := &mcp.Tool{
		Name:        "stock_fund_flow_individual",
		Description: "同花顺-数据中心-资金流向-个股资金流",
	}

	handler := func(ctx context.Context, req *mcp.CallToolRequest, args StockFundFlowIndividualArgs) (*mcp.CallToolResult, any, error) {
		client := resty.New()
		url := fmt.Sprintf("%s/api/public/stock_fund_flow_individual?symbol=%s",
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