package storage

import (
	"github.com/PaddingDEV/Sender/module/cfg"
	"github.com/google/uuid"
	"os"
)

func getFileUuid() string {
	return uuid.New().String()
}

func getFileStorePath(uuid string) string {
	return getFilePath(uuid) + "/file"
}

func getFileInfoPath(uuid string) string {
	return getFilePath(uuid) + "/info.json"
}

func getFilePath(uuid string) string {
	if len(uuid) < 2 {
	}
	return cfg.GetCfg().StorageBasePath + uuid[:2] + "/" + uuid
}

func ensurePathExist(path string) (isCreatedNew bool) {
	if isPathExists(path) {
		return false
	}
	_ = os.MkdirAll(path, os.ModeDir)
	return true
}

func isPathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
