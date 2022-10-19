package storage

import "github.com/google/uuid"

func getFileUuid() string {
	// TODO
	return uuid.New().String()
}

func getFilePath(uuid string) string {
	// TODO
	return "/some/path/on/server/" + uuid
}

func ensurePathExist(path string) (isCreatedNew bool) {
	// TODO
	return false
}
