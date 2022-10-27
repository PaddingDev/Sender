package storage

import (
	"github.com/PaddingDEV/Sender/module/middleware/auth"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	g := r.Group("/files")
	{
		g.Use(auth.GenMidAuth())
		g.POST("/upload", uploadFileHandler)
		g.GET("/download", getFileHandler)
		g.GET("/info", getFileInfoHandler)
	}
}
