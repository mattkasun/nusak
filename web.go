package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed: images/dontwin.png
var icon embed.FS

//go:embed assets/* html/* images/*
var f embed.FS

func main() {
	router := gin.Default()
	templates := template.Must(template.New("").ParseFS(f, "html/*"))
	router.SetHTMLTemplate(templates)
	router.StaticFS("/favicon.ico", http.FS(icon))
	router.GET("/images/*filepath", func(c *gin.Context) {
		c.FileFromFS(c.Request.URL.Path, http.FS(f))
	})
	router.GET("/html/*filepath", func(c *gin.Context) {
		c.FileFromFS(c.Request.URL.Path, http.FS(f))
	})
	router.GET("/assets/*filepath", func(c *gin.Context) {
		c.FileFromFS(c.Request.URL.Path, http.FS(f))
	})

	router.GET("/", mainPage)
	router.GET("/about", about)
	router.GET("/projects", projects)
	router.GET("/contact", contact)
	router.POST("/email", email)
	router.Run(":8081")
}

func mainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "main", "")
}

func about(c *gin.Context) {
	c.HTML(http.StatusOK, "about", "")
}

func projects(c *gin.Context) {
	c.HTML(http.StatusOK, "projects", "")
}

func contact(c *gin.Context) {
	c.HTML(http.StatusOK, "contact", "")
}
