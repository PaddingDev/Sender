package storage

import (
	"github.com/PaddingDEV/Sender/module/cfg"
	"github.com/PaddingDEV/Sender/module/middleware/auth"
	"github.com/PaddingDEV/Sender/module/middleware/param"
	"github.com/PaddingDEV/Sender/module/middleware/size"
	"github.com/PaddingDEV/Sender/module/service/storage/model"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	g := r.Group("/files")
	{
		g.Use(auth.GenMidAuth())
		g.POST("/upload",
			size.SizeLimiterHandler(cfg.GetCfg().MaxUploadBytes),
			param.MustHasHeaderHandler(model.ExpireAtHeader),
			uploadFileHandler)
	}

	g2 := r.Group("/files")
	{
		g2.Use(
			size.SizeLimiterHandler(cfg.GetCfg().MaxNormalRequestBytes),
			param.MustHasQueryHandlerAndSet(model.UuidHeader, model.UuidHeader),
			param.ShouldHasHeaderHandler(model.TokenHeader, ""))

		g2.GET("/download", getFileHandler)
		g2.GET("/info", getFileInfoHandler)
	}
}
