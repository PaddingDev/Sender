package param

import (
	"github.com/PaddingDEV/Sender/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MustHasHeaderHandler(header string) func(*gin.Context) {
	return func(c *gin.Context) {
		h, exists := c.Get(header)
		if !exists {
			utils.HttpReturnWithErrAndAbort(c,
				http.StatusBadRequest,
				"header lost")
			return
		}
		c.Set(header, h)
		c.Next()
	}
}

func ShouldHasHeaderHandler(header string, initV any) func(*gin.Context) {
	return func(c *gin.Context) {
		h, exists := c.Get(header)
		if !exists {
			c.Set(header, initV)
		} else {
			c.Set(header, h)
		}
		c.Next()
	}
}
