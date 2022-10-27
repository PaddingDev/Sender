package storage

import (
	"github.com/PaddingDEV/Sender/module/middleware/auth"
	"github.com/PaddingDEV/Sender/module/middleware/param"
	"github.com/PaddingDEV/Sender/module/service/storage/model"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	g := r.Group("/files")
	{
		g.Use(auth.GenMidAuth())
		g.POST("/upload",
			param.MustHasHeaderHandler(model.ExpireAtHeader),
			uploadFileHandler)
		g.GET("/download",
			param.MustHasQueryHandlerAndSet(model.UuidHeader, model.UuidHeader),
			param.ShouldHasHeaderHandler(model.TokenHeader, ""),
			getFileHandler)
		g.GET("/info",
			param.MustHasQueryHandlerAndSet(model.UuidHeader, model.UuidHeader),
			param.ShouldHasHeaderHandler(model.TokenHeader, ""),
			getFileInfoHandler)
	}
}
