package akshare

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"fi.mcp/internal/types"
	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/mcp"
)

// StockZhAHistArgs 定义 stock_zh_a_hist 工具的输入参数
type StockZhAHistArgs struct {
	Symbol string `json:"symbol" jsonschema:"股票代码, 如600001"`
	Period string `json:"period,omitempty" jsonschema:"周期（daily,weekly,monthly）"`
	Count  string `json:"count,omitempty" jsonschema:"获取数据的数量，默认值:284"`
}

type historyItem struct {
	NAMING_FAILED   string  `json:"日期"`
	NAMING_FAILED0  string  `json:"股票代码"`
	NAMING_FAILED1  float64 `json:"开盘"`
	NAMING_FAILED2  float64 `json:"收盘"`
	NAMING_FAILED3  float64 `json:"最高"`
	NAMING_FAILED4  float64 `json:"最低"`
	NAMING_FAILED5  int     `json:"成交量"`
	NAMING_FAILED6  float64 `json:"成交额"`
	NAMING_FAILED7  float64 `json:"振幅"`
	NAMING_FAILED8  float64 `json:"涨跌幅"`
	NAMING_FAILED9  float64 `json:"涨跌额"`
	NAMING_FAILED10 float64 `json:"换手率"`
}

func NewStockZhAHistTool(_mcp types.MCPProvider) (*mcp.Tool, func(context.Context, *mcp.CallToolRequest, StockZhAHistArgs) (*mcp.CallToolResult, any, error)) {
	tool := &mcp.Tool{
		Name:        "stock_zh_a_hist",
		Description: "东方财富-沪深京 A 股日频率数据; 历史数据按日频率更新, 当日收盘价请在收盘后获取",
	}

	handler := func(ctx context.Context, req *mcp.CallToolRequest, args StockZhAHistArgs) (*mcp.CallToolResult, any, error) {
		if args.Count == "" {
			args.Count = "284"
		}

		count, err := strconv.Atoi(args.Count)
		if err != nil || count <= 0 {
			count = 284
		}

		now := time.Now()

		endDate := now.Format("20060102")
		startDate := now.AddDate(0, 0, -1*count).Format("20060102")

		client := resty.New()
		url := fmt.Sprintf("%s/api/public/stock_zh_a_hist?symbol=%s&period=%s&start_date=%s&end_date=%s&adjust=qfq",
			_mcp.GetServiceContext().Config.DataSource.Akshare.Host,
			args.Symbol,
			args.Period,
			startDate,
			endDate,
		)

		_mcp.GetLogger().Infof("url: %s", url)

		var history []historyItem
		_, err = client.R().SetResult(&history).Get(url)
		if err != nil {
			return nil, nil, err
		}

		jsonBytes, err := json.Marshal(map[string]any{"items": history})
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