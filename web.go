package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed assets/* html/* images/*
var f embed.FS

func main() {
	router := gin.Default()
	templates := template.Must(template.New("").ParseFS(f, "html/*"))
	router.SetHTMLTemplate(templates)
	router.GET("/images/*filepath", func(c *gin.Context) {
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
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
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
