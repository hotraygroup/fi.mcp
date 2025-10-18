package logic

import (
	"context"

	"fi.mcp/internal/logic/akshare"
	"fi.mcp/internal/logic/okx"
	"fi.mcp/internal/logic/snowball"
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

// 确保 MCP 实现了 MCPProvider 接口
func (m *MCP) GetContext() context.Context {
	return m.ctx
}

func (m *MCP) GetServiceContext() *svc.ServiceContext {
	return m.svcCtx
}

func (m *MCP) GetLogger() logx.Logger {
	return m.Logger
}

func NewMCPLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MCP {

	// 创建 MCP 服务器
	server := mcp.NewMcpServer(svcCtx.Config.McpConf)

	_mcp := &MCP{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		mcpServer: server,
	}

	// 注册 snowball tool
	server.RegisterTool(snowball.NewSuggestTool(_mcp))
	server.RegisterTool(snowball.NewKlineTool(_mcp))
	server.RegisterTool(snowball.NewIndicatorTool(_mcp))
	server.RegisterTool(snowball.NewIncomeTool(_mcp))
	server.RegisterTool(snowball.NewBalanceTool(_mcp))
	server.RegisterTool(snowball.NewCashFlowTool(_mcp))

	// 注册 okx tool
	server.RegisterTool(okx.NewCandlesTool(_mcp))

	// 注册 akshare tool
	server.RegisterTool(akshare.NewStockZhAHistTool(_mcp))
	server.RegisterTool(akshare.NewStockCyqEmTool(_mcp))
	server.RegisterTool(akshare.NewStockIndividualFundFlowTool(_mcp))
	server.RegisterTool(akshare.NewStockFundFlowIndividualTool(_mcp))
	server.RegisterTool(akshare.NewStockFundFlowBigDealTool(_mcp))

	server.RegisterTool(akshare.NewStockFinancialDebtThsTool(_mcp))
	server.RegisterTool(akshare.NewStockFinancialBenefitThsTool(_mcp))
	server.RegisterTool(akshare.NewStockFinancialCashThsTool(_mcp))

	return _mcp
}

func (l *MCP) Start() {
	l.Info("mcp server start...")
	go l.mcpServer.Start()
}

func (l *MCP) Stop() {
	l.mcpServer.Stop()
	l.Info("mcp server stop...")
}
