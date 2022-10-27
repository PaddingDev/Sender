package utils

import (
	"github.com/gin-gonic/gin"
)

func HttpAbort(c *gin.Context, httpCode int, msg string) {
	c.AbortWithStatusJSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  msg,
	})
}

func HttpAbortIfNotNil(err error, c *gin.Context, httpCode int, msg string) (needRet bool) {
	if err != nil {
		HttpAbort(c, httpCode, msg)
		return true
	}
	return false
}

func HttpAbortIf(trueValidate bool, c *gin.Context, httpCode int, msg string) (needRet bool) {
	if trueValidate {
		HttpAbort(c, httpCode, msg)
		return true
	}
	return false
}

func GetGinQuery(c *gin.Context, key string) (val string, exist bool) {
	return c.GetQuery(key)
}
