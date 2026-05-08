package logic

import (
	"context"

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
	if tool, handler := snowball.NewSuggestTool(_mcp); tool != nil {
		mcp.AddTool(server, tool, handler)
	}
	if tool, handler := snowball.NewKlineTool(_mcp); tool != nil {
		mcp.AddTool(server, tool, handler)
	}
	if tool, handler := snowball.NewIndicatorTool(_mcp); tool != nil {
		mcp.AddTool(server, tool, handler)
	}
	if tool, handler := snowball.NewIncomeTool(_mcp); tool != nil {
		mcp.AddTool(server, tool, handler)
	}
	if tool, handler := snowball.NewBalanceTool(_mcp); tool != nil {
		mcp.AddTool(server, tool, handler)
	}
	if tool, handler := snowball.NewCashFlowTool(_mcp); tool != nil {
		mcp.AddTool(server, tool, handler)
	}
	if tool, handler := snowball.NewQuoteTool(_mcp); tool != nil {
		mcp.AddTool(server, tool, handler)
	}

	// 注册 okx tool
	// if tool, handler := okx.NewCandlesTool(_mcp); tool != nil {
	// 	mcp.AddTool(server, tool, handler)
	// }

	// 注册 akshare tool
	// if tool, handler := akshare.NewStockZhAHistTool(_mcp); tool != nil {
	// 	mcp.AddTool(server, tool, handler)
	// }
	// if tool, handler := akshare.NewStockCyqEmTool(_mcp); tool != nil {
	// 	mcp.AddTool(server, tool, handler)
	// }
	// if tool, handler := akshare.NewStockIndividualFundFlowTool(_mcp); tool != nil {
	// 	mcp.AddTool(server, tool, handler)
	// }
	// if tool, handler := akshare.NewStockFundFlowIndividualTool(_mcp); tool != nil {
	// 	mcp.AddTool(server, tool, handler)
	// }
	// if tool, handler := akshare.NewStockFundFlowBigDealTool(_mcp); tool != nil {
	// 	mcp.AddTool(server, tool, handler)
	// }

	// if tool, handler := akshare.NewStockFinancialDebtThsTool(_mcp); tool != nil {
	// 	mcp.AddTool(server, tool, handler)
	// }
	// if tool, handler := akshare.NewStockFinancialBenefitThsTool(_mcp); tool != nil {
	// 	mcp.AddTool(server, tool, handler)
	// }
	// if tool, handler := akshare.NewStockFinancialCashThsTool(_mcp); tool != nil {
	// 	mcp.AddTool(server, tool, handler)
	// }

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