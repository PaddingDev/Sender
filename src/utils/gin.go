package utils

import (
	"github.com/gin-gonic/gin"
)

func HttpReturnWithErrAndAbort(c *gin.Context, httpCode int, msg string) {
	c.AbortWithStatusJSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  msg,
	})
}
