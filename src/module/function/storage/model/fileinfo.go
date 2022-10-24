package model

import (
	"github.com/PaddingDEV/Sender/utils"
	"time"
)

type FileInfo struct {
	OriginFileName string    `json:"name"`
	Token          string    `json:"token"`
	ExpiredAt      time.Time `json:"expired_at"`
}

func CreateFileInfoJson(filename, token string, expiredAt time.Time) (string, error) {
	i := FileInfo{
		OriginFileName: filename,
		Token:          token,
		ExpiredAt:      expiredAt,
	}
	return utils.ToJson(i)
}
