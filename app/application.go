package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartAplication() {
	mapUrls()
	router.Run(":8080")
}
