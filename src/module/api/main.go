package api

import (
	"github.com/PaddingDEV/Sender/utils"
	"github.com/gin-gonic/gin"
)

func StartApiListening(listenAddr string) {
	r := gin.New()

	r.Use(
		gin.Recovery(),
	)

	// TODO:
	_ = r.Group("/")

	err := r.Run(listenAddr)
	utils.PanicIfNotNil(err, "GIN: Failed to listen %v\n")
}
