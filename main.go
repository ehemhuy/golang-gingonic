package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var router = gin.Default()
	godotenv.Load()
	port := os.Getenv("PORT")
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")

	})
	router.Run(":" + port)
	fmt.Printf("ok")
}
