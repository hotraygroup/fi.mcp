package snowball

import (
	"context"
	"encoding/json"
	"fmt"

	"fi.mcp/internal/types"
	"github.com/zeromicro/go-zero/mcp"
)

// QuoteArgs 定义 quote 工具的输入参数
type QuoteArgs struct {
	Symbol string `json:"symbol" jsonschema:"股票代码（如 SH600519）"`
}

var quoteDescription = map[string]string{
	"symbol":               "股票代码",
	"name":                 "股票名称",
	"current":              "当前价",
	"chg":                  "涨跌额",
	"percent":              "涨跌幅",
	"last_close":           "昨收价",
	"open":                 "开盘价",
	"high":                 "最高价",
	"low":                  "最低价",
	"high52w":              "52周最高",
	"low52w":               "52周最低",
	"volume":               "成交量",
	"amount":               "成交额",
	"volume_ratio":         "量比",
	"turnover_rate":        "换手率",
	"pe_ttm":               "市盈率TTM",
	"pe_lyr":               "市盈率LYR",
	"pb":                   "市净率",
	"eps":                  "每股收益",
	"navps":                "每股净资产",
	"dividend":             "每股派息",
	"dividend_yield":       "股息率",
	"market_capital":       "总市值",
	"float_market_capital": "流通市值",
	"total_shares":         "总股本",
	"float_shares":         "流通股本",
	"amplitude":            "振幅",
	"limit_up":             "涨停价",
	"limit_down":           "跌停价",
}

func NewQuoteTool(_mcp types.MCPProvider) (*mcp.Tool, func(context.Context, *mcp.CallToolRequest, QuoteArgs) (*mcp.CallToolResult, any, error)) {
	tool := &mcp.Tool{
		Name:        "quote",
		Description: "用股票代码获取对应的实时行情数据",
	}

	handler := func(ctx context.Context, req *mcp.CallToolRequest, args QuoteArgs) (*mcp.CallToolResult, any, error) {
		if args.Symbol == "" {
			return nil, nil, fmt.Errorf("symbol is required")
		}

		client := NewClientWithConfig(&_mcp.GetServiceContext().Config)
		setHeader(_mcp.GetServiceContext().Config.DataSource.UserAgent,
			_mcp.GetServiceContext().Config.DataSource.Snowball.IndexURL,
			_mcp.GetServiceContext().Config.DataSource.Snowball.CookieURL,
			&_mcp.GetServiceContext().Config, client)

		u := fmt.Sprintf(_mcp.GetServiceContext().Config.DataSource.Snowball.QuoteURL, args.Symbol)
		_mcp.GetLogger().Infof("quote url: %s", u)

		var quote types.Quote
		_, err := client.R().SetResult(&quote).Get(u)
		if err != nil {
			return nil, nil, err
		}

		if quote.ErrorCode != 0 {
			return nil, nil, fmt.Errorf("request error, code: %d, desc: %s",
				quote.ErrorCode, quote.ErrorDescription)
		}

		result := map[string]any{
			"columns": quoteDescription,
			"quote":   quote.Data.Quote,
			"market":  quote.Data.Market,
			"others":  quote.Data.Others,
			"tags":    quote.Data.Tags,
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