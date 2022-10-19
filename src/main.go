package main

import (
	"github.com/PaddingDEV/Sender/module/api"
	"github.com/PaddingDEV/Sender/module/cfg"
	"github.com/gin-gonic/gin"
)

func main() {
	if !cfg.GetCfg().IsDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	api.StartApiListening(cfg.GetCfg().ListenAddr)
}
