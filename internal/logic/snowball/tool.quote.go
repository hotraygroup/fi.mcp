package snowball

import (
	"context"
	"fmt"

	"fi.mcp/internal/types"
	"github.com/zeromicro/go-zero/mcp"
)

var quoteDescription = map[string]string{
	"symbol":              "股票代码",
	"name":                "股票名称",
	"current":             "当前价",
	"chg":                 "涨跌额",
	"percent":             "涨跌幅",
	"last_close":          "昨收价",
	"open":                "开盘价",
	"high":                "最高价",
	"low":                 "最低价",
	"high52w":             "52周最高",
	"low52w":              "52周最低",
	"volume":              "成交量",
	"amount":              "成交额",
	"volume_ratio":        "量比",
	"turnover_rate":       "换手率",
	"pe_ttm":              "市盈率TTM",
	"pe_lyr":              "市盈率LYR",
	"pb":                  "市净率",
	"eps":                 "每股收益",
	"navps":               "每股净资产",
	"dividend":            "每股派息",
	"dividend_yield":      "股息率",
	"market_capital":      "总市值",
	"float_market_capital": "流通市值",
	"total_shares":        "总股本",
	"float_shares":        "流通股本",
	"amplitude":           "振幅",
	"limit_up":            "涨停价",
	"limit_down":          "跌停价",
}

func NewQuoteTool(_mcp types.MCPProvider) mcp.Tool {
	var tool = mcp.Tool{
		Name:        "quote",
		Description: "用股票代码获取对应的实时行情数据",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"symbol": map[string]any{
					"type":        "string",
					"description": "股票代码（如 SH600519）",
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
			if req.Symbol == "" {
				return nil, fmt.Errorf("symbol is required")
			}

			client := NewClientWithConfig(&_mcp.GetServiceContext().Config)
			setHeader(_mcp.GetServiceContext().Config.DataSource.UserAgent,
				_mcp.GetServiceContext().Config.DataSource.Snowball.IndexURL,
				_mcp.GetServiceContext().Config.DataSource.Snowball.CookieURL,
				&_mcp.GetServiceContext().Config, client)

			url := fmt.Sprintf(_mcp.GetServiceContext().Config.DataSource.Snowball.QuoteURL, req.Symbol)
			_mcp.GetLogger().Infof("quote url: %s", url)

			var quote types.Quote
			_, err := client.R().SetResult(&quote).Get(url)
			if err != nil {
				return nil, err
			}

			if quote.ErrorCode != 0 {
				return nil, fmt.Errorf("request error, code: %d, desc: %s",
					quote.ErrorCode, quote.ErrorDescription)
			}

			return map[string]any{
				"columns": quoteDescription,
				"quote":   quote.Data.Quote,
				"market":  quote.Data.Market,
				"others":  quote.Data.Others,
				"tags":    quote.Data.Tags,
			}, nil
		},
	}
	return tool
}