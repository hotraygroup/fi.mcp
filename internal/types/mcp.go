package types

import (
	"context"

	"fi.mcp/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

// MCPProvider 定义 MCP 相关功能的接口
type MCPProvider interface {
	GetContext() context.Context
	GetServiceContext() *svc.ServiceContext // 或者具体的 *svc.ServiceContext 类型
	GetLogger() logx.Logger                 // 或者具体的 logger 类型
}
