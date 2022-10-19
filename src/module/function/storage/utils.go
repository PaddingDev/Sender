package storage

import "github.com/google/uuid"

func getFileName() string {
	// TODO
	return uuid.New().String()
}

func calcFilePath() string {
	// TODO
	return "/some/path/on/server/"
}
