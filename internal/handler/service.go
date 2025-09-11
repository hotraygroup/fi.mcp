package handler

import (
	"context"

	"fi.mcp/internal/logic"
	"fi.mcp/internal/svc"
	"github.com/zeromicro/go-zero/core/service"
)

func Register(svcCtx *svc.ServiceContext, group *service.ServiceGroup) {

	group.Add(logic.NewMCPLogic(context.Background(), svcCtx))

	group.Start()

}
