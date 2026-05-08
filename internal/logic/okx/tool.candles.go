package okx

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

// CandlesArgs 定义 candles 工具的输入参数
type CandlesArgs struct {
	InstId string `json:"instId" jsonschema:"交易对，如 BTC-USDT"`
	Bar    string `json:"bar,omitempty" jsonschema:"周期, 如 1m,3m,5m,15m,30m,1H,2H,4H,6H,12H,1D,1W,1M,3M，默认值:1D"`
	Count  string `json:"count,omitempty" jsonschema:"获取数据的数量，默认值:284"`
}

type candlesItem struct {
	Timestamp string  `json:"timestamp"`
	Open      float64 `json:"open"`
	Close     float64 `json:"close"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Volume    float64 `json:"volume"`
	Amount    float64 `json:"amount"`
}

var candlesItemDescription = map[string]string{
	"timestamp": "时间",
	"open":      "开盘价",
	"close":     "收盘价",
	"high":      "最高价",
	"low":       "最低价",
	"volume":    "成交量",
	"amount":    "成交额",
}

func NewCandlesTool(_mcp types.MCPProvider) (*mcp.Tool, func(context.Context, *mcp.CallToolRequest, CandlesArgs) (*mcp.CallToolResult, any, error)) {
	tool := &mcp.Tool{
		Name:        "candles",
		Description: "用交易对获取加密货币的K线数据",
	}

	handler := func(ctx context.Context, req *mcp.CallToolRequest, args CandlesArgs) (*mcp.CallToolResult, any, error) {
		if args.Bar == "" {
			args.Bar = "1D"
		}
		if args.Count == "" {
			args.Count = "284"
		}

		count, err := strconv.Atoi(args.Count)
		if err != nil || count <= 0 {
			count = 284
		}

		limit := 300
		var items []*candlesItem
		client := resty.New()
		client.SetTimeout(5 * time.Second)
		start := time.Now().UnixMilli()

		for {
			url := fmt.Sprintf(_mcp.GetServiceContext().Config.DataSource.OKX.Host+_mcp.GetServiceContext().Config.DataSource.OKX.Candles, args.InstId, args.Bar, limit, start)
			var response types.Candles
			_, err := client.R().SetResult(&response).Get(url)

			if err != nil || response.Code != "0" {
				_mcp.GetLogger().Errorf("get candles %s failed: %s", url, err)
				return nil, nil, err
			}

			if len(response.Data) == 0 {
				_mcp.GetLogger().Errorf("get candles %s failed: %s", url, "empty")
				break
			}

			_mcp.GetLogger().Infof("get candles success, %s-->%s", response.Data[len(response.Data)-1][0], response.Data[0][0])

			for i := range response.Data {
				ts, _ := strconv.ParseInt(response.Data[i][0], 10, 64)

				if ts < start {
					start = ts
				}

				o, err := strconv.ParseFloat(response.Data[i][1], 64)
				if err != nil {
					_mcp.GetLogger().Errorf("parse open failed: %s", err)
					return nil, nil, err
				}

				h, err := strconv.ParseFloat(response.Data[i][2], 64)
				if err != nil {
					_mcp.GetLogger().Errorf("parse high failed: %s", err)
					return nil, nil, err
				}

				l, err := strconv.ParseFloat(response.Data[i][3], 64)
				if err != nil {
					_mcp.GetLogger().Errorf("parse low failed: %s", err)
					return nil, nil, err
				}

				c, err := strconv.ParseFloat(response.Data[i][4], 64)
				if err != nil {
					_mcp.GetLogger().Errorf("parse close failed: %s", err)
					return nil, nil, err
				}

				v, err := strconv.ParseFloat(response.Data[i][5], 64)
				if err != nil {
					_mcp.GetLogger().Errorf("parse volume failed: %s", err)
					return nil, nil, err
				}

				a, err := strconv.ParseFloat(response.Data[i][6], 64)
				if err != nil {
					_mcp.GetLogger().Errorf("parse amount failed: %s", err)
					return nil, nil, err
				}

				items = append(items, &candlesItem{
					Timestamp: time.Unix(ts/1000, 0).Format(time.DateTime),
					Open:      o,
					Close:     c,
					High:      h,
					Low:       l,
					Volume:    v,
					Amount:    a,
				})

			}
			if len(response.Data) < limit {
				_mcp.GetLogger().Infof("get candles %s finished", url)
				break
			}
			if len(items) >= count {
				_mcp.GetLogger().Infof("get candles %s finished", url)
				break
			}
		}

		result := map[string]any{
			"columns": candlesItemDescription,
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