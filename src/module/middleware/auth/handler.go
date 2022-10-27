package auth

import (
	"github.com/PaddingDEV/Sender/model"
	"github.com/PaddingDEV/Sender/module/cfg"
	"github.com/PaddingDEV/Sender/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenMidAuth() func(ctx *gin.Context) {
	switch cfg.GetCfg().AuthType {
	case model.BasicAuth:
		return gin.BasicAuth(map[string]string(*cfg.GetCfg().BasicAuthList))
	case model.None:
		return noneAuth
	case model.OnlyDebugMod:
		return onlyDebugModAuth
	default:
		return noneAuth
	}
}

func noneAuth(c *gin.Context) {
	c.Next()
}

func onlyDebugModAuth(c *gin.Context) {
	if cfg.GetCfg().IsDebug {
		c.Next()
		return
	}
	utils.HttpReturnWithErrAndAbort(c, http.StatusUnauthorized, "Unauthorised!")
}
