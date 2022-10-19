package storage

import (
	"github.com/PaddingDEV/Sender/module/function/storage/model"
	"github.com/PaddingDEV/Sender/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		utils.HttpReturnWithErrAndAbort(c, http.StatusUnauthorized, "No file is received")
	}

	uuid := getFileUuid()
	path := getFilePath(uuid)
	ensurePathExist(path)

	infoJson, _ := model.CreateFileInfoJson(file.Filename, "")
	utils.WriteToFile(path+"info.json", infoJson)

	if err := c.SaveUploadedFile(file, path+file.Filename); err != nil {
		utils.HttpReturnWithErrAndAbort(c,
			http.StatusInternalServerError,
			"Save failed")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
