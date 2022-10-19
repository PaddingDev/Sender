package storage

import (
	"github.com/PaddingDEV/Sender/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func SaveFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		utils.HttpReturnWithErrAndAbort(c, http.StatusUnauthorized, "No file is received")
	}

	extension := filepath.Ext(file.Filename)
	newFileName := getFileName() + extension

	if err := c.SaveUploadedFile(file, calcFilePath()+newFileName); err != nil {
		utils.HttpReturnWithErrAndAbort(c,
			http.StatusInternalServerError,
			"Save failed")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
