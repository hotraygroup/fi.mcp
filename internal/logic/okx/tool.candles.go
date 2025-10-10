package okx

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"fi.mcp/internal/types"
	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/mcp"
)

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

// "/api/v5/market/candles?instId=%s&bar=%s&limit=300&after=%v"
/*
	"1m":  "1m",
	"3m":  "3m",
	"5m":  "5m",
	"15m": "15m",
	"30m": "30m",
	"1h":  "1H",
	"2h":  "2H",
	"4h":  "4H",
	"6h":  "6H",
	"12h": "12H",
	"1d":  "1D",
	"1w":  "1W",
	"1M":  "1M",
	"3M":  "3M",
*/
func NewCandlesTool(_mcp types.MCPProvider) mcp.Tool {
	var candlesTool = mcp.Tool{
		Name:        "candles",
		Description: "用交易对获取加密货币的K线数据",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"instId": map[string]any{
					"type":        "string",
					"description": "交易对，如 BTC-USDT",
				},
				"bar": map[string]any{
					"type":        "string",
					"description": "周期, 如 1m,3m,5m,15m,30m,1H,2H,4H,6H,12H,1D,1W,1M,3M",
					"default":     "1D",
				},
				"count": map[string]any{
					"type":        "string",
					"description": "获取数据的数量",
					"default":     "284",
				},
			},
			Required: []string{"instId"},
		},
		Handler: func(ctx context.Context, params map[string]any) (any, error) {
			var req struct {
				InstId string `json:"instId"`
				Bar    string `json:"bar,omitempty"`
				Count  string `json:"count,omitempty"`
			}

			if err := mcp.ParseArguments(params, &req); err != nil {
				return nil, fmt.Errorf("failed to parse params: %w", err)
			}

			if req.Bar == "" {
				req.Bar = "1D"
			}
			if req.Count == "" {
				req.Count = "284"
			}

			count, err := strconv.Atoi(req.Count)
			if err != nil || count <= 0 {
				count = 284
			}

			limit := 300
			var items []*candlesItem
			client := resty.New()
			client.SetTimeout(5 * time.Second)
			start := time.Now().UnixMilli()

			for {
				url := fmt.Sprintf(_mcp.GetServiceContext().Config.DataSource.OKX.Host+_mcp.GetServiceContext().Config.DataSource.OKX.Candles, req.InstId, req.Bar, limit, start)
				var response types.Candles
				_, err := client.R().SetResult(&response).Get(url)

				if err != nil || response.Code != "0" {
					_mcp.GetLogger().Errorf("get candles %s failed: %s", url, err)
					return nil, err
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
						return nil, err
					}

					h, err := strconv.ParseFloat(response.Data[i][2], 64)
					if err != nil {
						_mcp.GetLogger().Errorf("parse high failed: %s", err)
						return nil, err
					}

					l, err := strconv.ParseFloat(response.Data[i][3], 64)
					if err != nil {
						_mcp.GetLogger().Errorf("parse low failed: %s", err)
						return nil, err
					}

					c, err := strconv.ParseFloat(response.Data[i][4], 64)
					if err != nil {
						_mcp.GetLogger().Errorf("parse close failed: %s", err)
						return nil, err
					}

					v, err := strconv.ParseFloat(response.Data[i][5], 64)
					if err != nil {
						_mcp.GetLogger().Errorf("parse volume failed: %s", err)
						return nil, err
					}

					a, err := strconv.ParseFloat(response.Data[i][6], 64)
					if err != nil {
						_mcp.GetLogger().Errorf("parse amount failed: %s", err)
						return nil, err
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

			return map[string]any{
				"columns": candlesItemDescription,
				"items":   items,
			}, nil
		},
	}
	return candlesTool
}
