package storage

import (
	"github.com/PaddingDEV/Sender/module/function/storage/model"
	"github.com/PaddingDEV/Sender/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetFileHandler(c *gin.Context) {
	uuid := c.GetString("uuid")
	token := c.GetString("token")
	if uuid == "" {
		utils.HttpReturnWithErrAndAbort(
			c,
			http.StatusBadRequest,
			"need uuid")
		return
	}
	path := getFilePath(uuid)
	if !isPathExists(path) {
		utils.HttpReturnWithErrAndAbort(
			c,
			http.StatusBadRequest,
			"No file")
	}
	infoBs, err := utils.ReadFileToByte(path + "info.json")
	if err != nil {
		utils.HttpReturnWithErrAndAbort(
			c,
			http.StatusBadRequest,
			"no info")
	}
	info := model.FileInfo{}
	err = utils.FromJsonTo(infoBs, &info)
	if err != nil {
		utils.HttpReturnWithErrAndAbort(
			c,
			http.StatusBadRequest,
			"failed to load info")
	}
	if info.Token != token {
		utils.HttpReturnWithErrAndAbort(
			c,
			http.StatusUnauthorized,
			"Wrong token")
	}
	c.FileAttachment(path+info.OriginFileName, info.OriginFileName)
}
