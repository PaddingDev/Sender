package auth

import (
	"github.com/PaddingDEV/Sender/module/cfg"
	"github.com/PaddingDEV/Sender/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MidAuth(c *gin.Context) {
	if cfg.GetCfg().IsDebug {
		c.Next()
		return
	}
	utils.HttpReturnWithErrAndAbort(c, http.StatusUnauthorized, "Unauthorised!")
}
