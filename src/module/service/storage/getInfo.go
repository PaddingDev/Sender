package storage

import (
	"github.com/PaddingDEV/Sender/module/service/storage/model"
	"github.com/PaddingDEV/Sender/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func getFileInfoHandler(c *gin.Context) {
	uuid := c.GetString(model.UuidHeader)
	token := c.GetString(model.TokenHeader)
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
	if time.Now().After(info.ExpiredAt) {
		utils.HttpReturnWithErrAndAbort(
			c,
			http.StatusBadRequest,
			"expired!")
		_ = utils.RemoveFile(getFileStorePath(uuid))
	}
	if info.Token != token {
		utils.HttpReturnWithErrAndAbort(
			c,
			http.StatusUnauthorized,
			"Wrong token")
	}
	info.Token = ""
	c.JSON(http.StatusOK, info)
}
