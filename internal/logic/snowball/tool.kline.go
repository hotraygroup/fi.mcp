package snowball

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"fi.mcp/internal/types"
	"github.com/zeromicro/go-zero/mcp"
)

// KlineArgs 定义 kline 工具的输入参数
type KlineArgs struct {
	Symbol string `json:"symbol" jsonschema:"股票代码"`
	Period string `json:"period,omitempty" jsonschema:"周期，默认值:day"`
	Count  string `json:"count,omitempty" jsonschema:"每次获取数量，默认值:284"`
	Days   string `json:"days,omitempty" jsonschema:"最近多少天，默认值:360"`
}

type KlineItem struct {
	Timestamp string  `json:"timestamp"`
	Open      float64 `json:"open"`
	Close     float64 `json:"close"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Volume    float64 `json:"volume"`
	Amount    float64 `json:"amount"`
}

var klineDescription = map[string]string{
	"timestamp": "时间",
	"open":      "开盘价",
	"close":     "收盘价",
	"high":      "最高价",
	"low":       "最低价",
	"volume":    "成交量",
	"amount":    "成交额",
}

func NewKlineTool(_mcp types.MCPProvider) (*mcp.Tool, func(context.Context, *mcp.CallToolRequest, KlineArgs) (*mcp.CallToolResult, any, error)) {
	tool := &mcp.Tool{
		Name:        "kline",
		Description: "用股票代码获取对应的K线数据",
	}

	handler := func(ctx context.Context, req *mcp.CallToolRequest, args KlineArgs) (*mcp.CallToolResult, any, error) {
		if args.Symbol == "" {
			return nil, nil, fmt.Errorf("symbol is required")
		}

		if args.Period == "" {
			args.Period = "day"
		}

		if args.Count == "" {
			args.Count = "284"
		}

		if args.Days == "" {
			args.Days = "360"
		}

		lastDay := "0000-00-00"

		nDays, _ := strconv.Atoi(args.Days)
		if nDays <= 0 {
			nDays = 360
		}

		kline := types.Kline{}
		var items []*KlineItem
		now := time.Now()
		maxTimeStamp := now.UnixMilli()

		client := NewClientWithConfig(&_mcp.GetServiceContext().Config)
		setHeader(_mcp.GetServiceContext().Config.DataSource.UserAgent, _mcp.GetServiceContext().Config.DataSource.Snowball.IndexURL, _mcp.GetServiceContext().Config.DataSource.Snowball.CookieURL, &_mcp.GetServiceContext().Config, client)
		finished := false
		for {
			url := fmt.Sprintf(_mcp.GetServiceContext().Config.DataSource.Snowball.KlineURL, args.Symbol, maxTimeStamp, args.Period, args.Count)
			_mcp.GetLogger().Infof("url: %s", url)
			_, err := client.R().SetResult(&kline).Get(url)
			if err != nil {
				return nil, nil, err
			}

			if kline.ErrorCode != 0 {
				return nil, nil, fmt.Errorf("request error, code: %d", kline.ErrorCode)
			}

			_mcp.GetLogger().Infof("data item count: %d", len(kline.Data.Item))

			if len(kline.Data.Item) == 0 {
				break
			}

			for i := len(kline.Data.Item) - 1; i >= 0; i-- {

				t := time.Unix(int64(kline.Data.Item[i][0])/1000, 0)

				currentDay := t.Format("2006-01-02")
				if currentDay != lastDay {
					lastDay = currentDay
					nDays -= 1
					if nDays < 0 {
						_mcp.GetLogger().Infof("currentDay: %s, finished", currentDay)
						finished = true
						break
					}
				}

				item := &KlineItem{
					Timestamp: t.Format(time.DateTime),
					Volume:    kline.Data.Item[i][1],
					Open:      kline.Data.Item[i][2],
					High:      kline.Data.Item[i][3],
					Low:       kline.Data.Item[i][4],
					Close:     kline.Data.Item[i][5],
					Amount:    kline.Data.Item[i][9],
				}
				items = append([]*KlineItem{item}, items...)

				maxTimeStamp = int64(kline.Data.Item[i][0]) - 1
			}

			if finished {
				break
			}
		}

		result := map[string]any{
			"columns": klineDescription,
			"items":   items,
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