package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Загружаем все шаблоны из папки templates
	templ := template.Must(template.New("").ParseGlob("./web/templates/*.html"))
	if templ == nil {
		log.Fatal("Failed to load templates")
	}
	router.SetHTMLTemplate(templ)
	router.SetFuncMap(template.FuncMap{})

	// Главная страница
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Home",
			"content": "Welcome to the homepage!",
		})
	})

	// Страница First
	router.GET("/first", func(c *gin.Context) {
		c.HTML(http.StatusOK, "first.html", gin.H{
			"title": "First Page",
		})
	})

	// Страница Second
	router.GET("/second", func(c *gin.Context) {
		c.HTML(http.StatusOK, "second.html", gin.H{
			"title": "Second Page",
		})
	})

	// Запускаем сервер на порту 8080
	router.Run(":8080")
}
