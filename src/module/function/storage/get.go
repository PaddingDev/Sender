package storage

import (
	"github.com/PaddingDEV/Sender/module/function/storage/model"
	"github.com/PaddingDEV/Sender/utils"
	"github.com/gin-gonic/gin"
)

func GetFileHandler(c *gin.Context) {
	uuid := c.GetString("uuid")
	token := c.GetString("token")
	if uuid == "" {

	}
	path := getFilePath(uuid)
	if !isPathExists(path) {

	}
	infoBs, err := utils.ReadFileToByte(path + "info.json")
	if err != nil {

	}
	info := model.FileInfo{}
	err = utils.FromJsonTo(infoBs, &info)
	if err != nil {

	}
	if info.Token != token {

	}
	c.FileAttachment(path+info.OriginFileName, info.OriginFileName)
}
