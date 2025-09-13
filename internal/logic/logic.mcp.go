package logic

import (
	"context"

	"fi.mcp/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/mcp"
)

type MCP struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	mcpServer mcp.McpServer
}

func NewMCPLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MCP {

	// 创建 MCP 服务器
	server := mcp.NewMcpServer(svcCtx.Config.McpConf)

	// 注册 mcp tool
	server.RegisterTool(newSuggestTool(svcCtx))
	server.RegisterTool(newKlineTool(svcCtx))
	server.RegisterTool(newIndicatorTool(svcCtx))
	server.RegisterTool(newIncomeTool(svcCtx))
	server.RegisterTool(newBalanceTool(svcCtx))
	server.RegisterTool(newCashFlowTool(svcCtx))

	return &MCP{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		mcpServer: server,
	}
}

func (l *MCP) Start() {
	l.Info("mcp server start...")
	go l.mcpServer.Start()
}

func (l *MCP) Stop() {
	l.mcpServer.Stop()
	l.Info("mcp server stop...")
}
