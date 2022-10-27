package size

import (
	"github.com/PaddingDEV/Sender/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SizeLimiterHandler(maxBytes int64) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		lengthS := c.GetHeader("Content-Length")
		if lengthS != "" {
			length64, err := strconv.ParseInt(lengthS, 10, 64)
			if err != nil && length64 > maxBytes {
				utils.HttpAbort(c, http.StatusBadRequest, "over size!")
				return
			}
		}
		var w http.ResponseWriter = c.Writer
		c.Request.Body = http.MaxBytesReader(w, c.Request.Body, maxBytes)
		c.Next()
	}
}
