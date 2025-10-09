package logic

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"fi.mcp/internal/types"
	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/mcp"
)

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

func newKlineTool(_mcp *MCP) mcp.Tool {

	var klineTool = mcp.Tool{
		Name:        "kline",
		Description: "用股票代码获取对应的K线数据",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"symbol": map[string]any{
					"type":        "string",
					"description": "股票代码",
				},
				"period": map[string]any{
					"type":        "string",
					"description": "周期",
					"default":     "day",
				},
				"count": map[string]any{
					"type":        "string",
					"description": "每次获取数量",
					"default":     "284",
				},
				"days": map[string]any{
					"type":        "string",
					"description": "最近多少天",
					"default":     "360",
				},
			},
			Required: []string{"symbol"},
		},
		Handler: func(ctx context.Context, params map[string]any) (any, error) {
			var req struct {
				Symbol string `json:"symbol"`
				Period string `json:"period,omitempty"`
				Count  string `json:"count,omitempty"`
				Days   string `json:"days,omitempty"`
			}

			if err := mcp.ParseArguments(params, &req); err != nil {
				return nil, fmt.Errorf("failed to parse params: %w", err)
			}
			if req.Symbol == "" {
				return nil, fmt.Errorf("symbol is required")
			}

			if req.Period == "" {
				req.Period = "day"
			}

			if req.Count == "" {
				req.Count = "284"
			}

			if req.Days == "" {
				req.Days = "360"
			}

			lastDay := "0000-00-00"

			nDays, _ := strconv.Atoi(req.Days)
			if nDays <= 0 {
				nDays = 360
			}

			kline := types.Kline{}
			var items []*KlineItem
			now := time.Now()
			maxTimeStamp := now.UnixMilli()

			client := resty.New()
			setHeader(_mcp.svcCtx.Config.DataSource.UserAgent, _mcp.svcCtx.Config.DataSource.Snowball.IndexURL, _mcp.svcCtx.Config.DataSource.Snowball.CookieURL, client)
			finished := false
			for {
				url := fmt.Sprintf(_mcp.svcCtx.Config.DataSource.Snowball.KlineURL, req.Symbol, maxTimeStamp, req.Period, req.Count)
				_mcp.Infof("url: %s", url)
				_, err := client.R().SetResult(&kline).Get(url)
				if err != nil {
					return nil, err
				}

				if kline.ErrorCode != 0 {
					return nil, fmt.Errorf("request error, code: %d", kline.ErrorCode)
				}

				_mcp.Infof("data item count: %d", len(kline.Data.Item))

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
							_mcp.Infof("currentDay: %s, finished", currentDay)
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

			return map[string]any{
				"columns": klineDescription,
				"items":   items,
			}, nil
		},
	}
	return klineTool
}
