package storage

import (
	"github.com/PaddingDEV/Sender/module/function/storage/model"
	"github.com/PaddingDEV/Sender/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SaveFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")

	// TODO:
	expireAt := time.Now()
	expireAt = expireAt.Add(15 * time.Minute)

	if err != nil {
		utils.HttpReturnWithErrAndAbort(c, http.StatusUnauthorized, "No file is received")
	}

	uuid := getFileUuid()
	path := getFilePath(uuid)
	isCreated := ensurePathExist(path)
	if !isCreated {
		utils.HttpReturnWithErrAndAbort(c, http.StatusBadGateway, "No idea")
	}

	infoJson, _ := model.CreateFileInfoJson(file.Filename, "", expireAt)
	_ = utils.WriteToFile(getFileInfoPath(uuid), infoJson)

	if err := c.SaveUploadedFile(file, getFileStorePath(uuid)); err != nil {
		// TODO: remove
		utils.HttpReturnWithErrAndAbort(c,
			http.StatusInternalServerError,
			"Save failed")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
