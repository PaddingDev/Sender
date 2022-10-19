package model

import (
	"github.com/PaddingDEV/Sender/utils"
)

type FileInfo struct {
	OriginFileName string `json:"name"`
	Token          string `json:"token"`
}

func CreateFileInfoJson(filename, token string) (string, error) {
	i := FileInfo{
		OriginFileName: filename,
		Token:          token,
	}
	return utils.ToJson(i)
}
