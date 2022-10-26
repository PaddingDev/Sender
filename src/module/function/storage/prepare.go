package storage

import (
	"github.com/PaddingDEV/Sender/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func genMustHasHeaderHandler(header string, c *gin.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		h, exists := c.Get(header)
		if !exists {
			utils.HttpReturnWithErrAndAbort(c,
				http.StatusBadRequest,
				"header loss")
			return
		}
		c.Set(header, h)
		c.Next()
	}
}

func genShouldHasHeaderHandler(header string, initV any, c *gin.Context) func(*gin.Context) {
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
