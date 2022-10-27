package param

import (
	"github.com/PaddingDEV/Sender/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MustHasQueryHandlerAndSet(key string, targetKey string) func(*gin.Context) {
	return func(c *gin.Context) {
		h, exists := c.GetQuery(key)
		if utils.HttpAbortIf(!exists, c, http.StatusBadRequest, "header lost") {
			return
		}
		c.Set(targetKey, h)
		c.Next()
	}
}

func ShouldHasQueryHandler(header string, targetKey string, initV string) func(*gin.Context) {
	return func(c *gin.Context) {
		h, exists := c.GetQuery(header)
		if !exists {
			c.Set(targetKey, initV)
		} else {
			c.Set(targetKey, h)
		}
		c.Next()
	}
}
