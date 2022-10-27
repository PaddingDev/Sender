package storage

import (
	"github.com/PaddingDEV/Sender/module/cfg"
	"github.com/PaddingDEV/Sender/module/service/storage/model"
	"github.com/PaddingDEV/Sender/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func uploadFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	token := c.GetString(model.TokenHeader)
	expVal := c.GetInt(model.ExpireAtHeader)
	if expVal > cfg.GetCfg().MaxExpireTime || expVal < cfg.GetCfg().MinExpireTime {
		utils.HttpReturnWithErrAndAbort(c, http.StatusBadRequest, "out of range")
		return
	}

	expireAt := time.Now()
	expireAt = expireAt.Add(time.Duration(expVal) * cfg.GetExpTimeUnit())

	if err != nil {
		utils.HttpReturnWithErrAndAbort(c, http.StatusUnauthorized, "No file is received")
	}

	uuid := getFileUuid()
	path := getFilePath(uuid)
	isCreated := ensurePathExist(path)
	if !isCreated {
		utils.HttpReturnWithErrAndAbort(c, http.StatusBadGateway, "No idea")
	}

	infoJson, _ := model.CreateFileInfoJson(file.Filename, token, expireAt)
	err = utils.WriteToFile(getFileInfoPath(uuid), infoJson)
	if err != nil {
		_ = os.RemoveAll(path)
		utils.HttpReturnWithErrAndAbort(c,
			http.StatusInternalServerError,
			"failed to save info")
		return
	}

	if err = c.SaveUploadedFile(file, getFileStorePath(uuid)); err != nil {
		_ = os.RemoveAll(path)
		utils.HttpReturnWithErrAndAbort(c,
			http.StatusInternalServerError,
			"Save failed")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}