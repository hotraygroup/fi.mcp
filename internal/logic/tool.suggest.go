package logic

import (
	"context"
	"fmt"
	"net/url"

	"fi.mcp/internal/types"
	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/mcp"
)

func newSuggestTool(_mcp *MCP) mcp.Tool {
	var suggestTool = mcp.Tool{
		Name:        "suggest",
		Description: "由公司名称或简称获取股票代码",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"company": map[string]any{
					"type":        "string",
					"description": "公司名称或简称",
				},
			},
			Required: []string{"company"},
		},
		Handler: func(ctx context.Context, params map[string]any) (any, error) {
			var req struct {
				Company string `json:"company"`
			}

			if err := mcp.ParseArguments(params, &req); err != nil {
				return nil, fmt.Errorf("failed to parse params: %w", err)
			}

			suggest := types.Suggest{}

			url := fmt.Sprintf(_mcp.svcCtx.Config.DataSource.Snowball.SuggestURL, url.QueryEscape(req.Company))

			client := resty.New()
			setHeader(_mcp.svcCtx.Config.DataSource.UserAgent, _mcp.svcCtx.Config.DataSource.Snowball.IndexURL, client)

			resp, err := client.R().SetResult(&suggest).Get(url)
			_mcp.Infof("url: %s, body: %s", url, resp.String())

			if err != nil {
				return nil, fmt.Errorf("request error: %w", err)
			}

			_mcp.Infof("suggest: %+v", suggest)

			symbol := ""

			if suggest.Code == 200 && len(suggest.Data) > 0 {
				symbol = suggest.Data[0].Code
			} else {
				return nil, fmt.Errorf("internal error: not found")
			}

			return mcp.ToolResult{
				Type:    mcp.ContentTypeText,
				Content: symbol,
			}, nil
		},
	}
	return suggestTool
}
