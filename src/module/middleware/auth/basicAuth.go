package auth

import "github.com/gin-gonic/gin"

func GenBasicAuth(m map[string]string) func(*gin.Context) {
	return gin.BasicAuth(m)
}
