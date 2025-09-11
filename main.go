package main

import (
	"os"
	"os/signal"
	"syscall"

	"fi.mcp/internal/config"
	"fi.mcp/internal/handler"
	"fi.mcp/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

func main() {
	// 从 YAML 文件加载配置
	var c config.Config
	conf.MustLoad("etc/fi.mcp.yaml", &c)

	// 可选：禁用统计日志
	logx.DisableStat()

	svcCtx := svc.NewServiceContext(c)

	group := service.NewServiceGroup()

	handler.Register(svcCtx, group)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-ch
		logx.Infof("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			logx.Info("stop service group...")
			group.Stop()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
