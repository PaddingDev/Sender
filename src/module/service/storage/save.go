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
		utils.HttpAbort(c, http.StatusBadRequest, "out of range")
		return
	}

	expireAt := time.Now()
	expireAt = expireAt.Add(time.Duration(expVal) * cfg.GetExpTimeUnit())

	if utils.HttpAbortIfNotNil(err, c, http.StatusUnauthorized, "No file is received") {
		return
	}

	uuid := getFileUuid()
	path := getFilePath(uuid)
	isCreated := ensurePathExist(path)
	if utils.HttpAbortIf(!isCreated, c, http.StatusBadGateway, "No idea") {
		return
	}

	infoJson, _ := model.CreateFileInfoJson(file.Filename, token, expireAt)
	err = utils.WriteToFile(getFileInfoPath(uuid), infoJson)
	if utils.HttpAbortIfNotNil(err, c, http.StatusInternalServerError, "failed to save info") {
		_ = os.RemoveAll(path)
		return
	}

	if err = c.SaveUploadedFile(file, getFileStorePath(uuid)); err != nil {
		_ = os.RemoveAll(path)
		utils.HttpAbort(c, http.StatusInternalServerError, "Save failed")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
