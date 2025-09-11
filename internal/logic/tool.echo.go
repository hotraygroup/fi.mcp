package logic

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/mcp"
)

// 简单的回显工具
var echoTool = mcp.Tool{
	Name:        "echo",
	Description: "回显用户提供的消息",
	InputSchema: mcp.InputSchema{
		Properties: map[string]any{
			"message": map[string]any{
				"type":        "string",
				"description": "要回显的消息",
			},
			"prefix": map[string]any{
				"type":        "string",
				"description": "可选的前缀，添加到回显消息前",
				"default":     "Echo: ",
			},
		},
		Required: []string{"message"},
	},
	Handler: func(ctx context.Context, params map[string]any) (any, error) {
		var req struct {
			Message string `json:"message"`
			Prefix  string `json:"prefix,omitempty"`
		}

		if err := mcp.ParseArguments(params, &req); err != nil {
			return nil, fmt.Errorf("failed to parse params: %w", err)
		}

		prefix := "Echo: "
		if len(req.Prefix) > 0 {
			prefix = req.Prefix
		}

		return prefix + req.Message, nil
	},
}
