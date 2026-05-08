package akshare

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"fi.mcp/internal/types"
	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/mcp"
)

// StockFinancialCashThsArgs 定义 stock_financial_cash_ths 工具的输入参数
type StockFinancialCashThsArgs struct {
	Symbol string `json:"symbol" jsonschema:"股票代码, 如600001"`
	Count  string `json:"count,omitempty" jsonschema:"获取数据的数量，默认值:12"`
}

func NewStockFinancialCashThsTool(_mcp types.MCPProvider) (*mcp.Tool, func(context.Context, *mcp.CallToolRequest, StockFinancialCashThsArgs) (*mcp.CallToolResult, any, error)) {
	tool := &mcp.Tool{
		Name:        "stock_financial_cash_ths",
		Description: "同花顺-财务指标-现金流量表",
	}

	handler := func(ctx context.Context, req *mcp.CallToolRequest, args StockFinancialCashThsArgs) (*mcp.CallToolResult, any, error) {
		if args.Count == "" {
			args.Count = "12"
		}

		count, err := strconv.Atoi(args.Count)
		if err != nil || count <= 0 {
			count = 12
		}

		client := resty.New()
		url := fmt.Sprintf("%s/api/public/stock_financial_cash_ths?symbol=%s",
			_mcp.GetServiceContext().Config.DataSource.Akshare.Host,
			args.Symbol,
		)

		_mcp.GetLogger().Infof("url: %s", url)

		var result []map[string]any

		_, err = client.R().SetResult(&result).Get(url)
		if err != nil {
			return nil, nil, err
		}

		if len(result) > count {
			result = result[:count]
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