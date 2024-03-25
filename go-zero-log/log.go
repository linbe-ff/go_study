package go_zero_log

import (
	"context"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

func InitLog(cfg logx.LogConf) {

	cfg.Path = "D:\\code\\log"
	_ = conf.FillDefault(&cfg)
	cfg.Mode = "file"
	logc.MustSetup(cfg)
}

func logTest() {
	var cfg logx.LogConf
	InitLog(cfg)
	defer logc.Close()
	logc.Info(context.Background(), "hello world")
}
