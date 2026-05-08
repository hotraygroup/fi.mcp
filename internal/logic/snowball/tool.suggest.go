package snowball

import (
	"context"
	"fmt"
	"net/url"

	"fi.mcp/internal/types"
	"github.com/zeromicro/go-zero/mcp"
)

// SuggestArgs 定义 suggest 工具的输入参数
type SuggestArgs struct {
	Company string `json:"company" jsonschema:"公司名称或简称"`
}

func NewSuggestTool(_mcp types.MCPProvider) (*mcp.Tool, func(context.Context, *mcp.CallToolRequest, SuggestArgs) (*mcp.CallToolResult, any, error)) {
	tool := &mcp.Tool{
		Name:        "suggest",
		Description: "由公司名称或简称获取股票代码",
	}

	handler := func(ctx context.Context, req *mcp.CallToolRequest, args SuggestArgs) (*mcp.CallToolResult, any, error) {
		suggest := types.Suggest{}

		u := fmt.Sprintf(_mcp.GetServiceContext().Config.DataSource.Snowball.SuggestURL, url.QueryEscape(args.Company))

		client := NewClientWithConfig(&_mcp.GetServiceContext().Config)
		setHeader(_mcp.GetServiceContext().Config.DataSource.UserAgent, _mcp.GetServiceContext().Config.DataSource.Snowball.IndexURL, _mcp.GetServiceContext().Config.DataSource.Snowball.CookieURL, &_mcp.GetServiceContext().Config, client)

		resp, err := client.R().SetResult(&suggest).Get(u)
		_mcp.GetLogger().Infof("url: %s, body: %s", u, resp.String())

		if err != nil {
			return nil, nil, fmt.Errorf("request error: %w", err)
		}

		_mcp.GetLogger().Infof("suggest: %+v", suggest)

		symbol := ""

		if suggest.Code == 200 && len(suggest.Data) > 0 {
			symbol = suggest.Data[0].Code
		} else {
			return nil, nil, fmt.Errorf("internal error: not found")
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: symbol},
			},
		}, nil, nil
	}

	return tool, handler
}