package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func mainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "")
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("html/*")
	router.Static("/stylesheet", "./stylesheet")
	router.StaticFile("favicon.ico", "./resources/favicon.ico")
	router.Static("images", "./images")
	router.Static("html", "./html")
	router.GET("/", mainPage)
	router.Run(":8081")
}
