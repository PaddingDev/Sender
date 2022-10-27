package storage

import (
	"github.com/PaddingDEV/Sender/module/service/storage/model"
	"github.com/PaddingDEV/Sender/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func getFileInfoHandler(c *gin.Context) {
	info, needRet := getFileInfo(c)
	if needRet {
		return
	}
	info.Token = ""
	c.JSON(200, info)
}

func getFileInfo(c *gin.Context) (info model.FileInfo, isAborted bool) {
	isAborted = true
	uuid := c.GetString(model.UuidHeader)
	token := c.GetString(model.TokenHeader)
	if utils.HttpAbortIf(uuid == "", c,
		http.StatusBadRequest,
		"need uuid") {
		return
	}

	path := getFilePath(uuid)
	if utils.HttpAbortIf(!isPathExists(path), c,
		http.StatusBadRequest,
		"No file") {
		return
	}

	infoBs, err := utils.ReadFileToByte(path + "info.json")
	if utils.HttpAbortIfNotNil(err, c,
		http.StatusBadRequest,
		"no info") {
		return
	}

	err = utils.FromJsonTo(infoBs, &info)
	if utils.HttpAbortIfNotNil(err, c,
		http.StatusBadRequest,
		"failed to load info") {
		return
	}

	if utils.HttpAbortIf(
		time.Now().After(info.ExpiredAt),
		c, http.StatusBadRequest,
		"expired!") {
		_ = utils.RemoveFile(getFileStorePath(uuid))
		return
	}

	if utils.HttpAbortIf(
		info.Token != token,
		c, http.StatusUnauthorized,
		"Wrong token") {
		return
	}
	return info, false
}
