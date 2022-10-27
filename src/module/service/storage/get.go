package storage

import (
	"github.com/PaddingDEV/Sender/module/service/storage/model"
	"github.com/gin-gonic/gin"
)

func getFileHandler(c *gin.Context) {
	uuid := c.GetString(model.UuidHeader)
	info, isAborted := getFileInfo(c)
	if isAborted {
		return
	}
	c.FileAttachment(getFileStorePath(uuid), info.OriginFileName)
}
