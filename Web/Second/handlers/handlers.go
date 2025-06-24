package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Created_at time.Time `json:"created_at"`
}

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	router.GET("/feed", func(c *gin.Context) {
		posts, err := getPosts(db)
		fmt.Println(posts)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"error": err.Error(),
			})
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main page",
			"posts": posts,
		})
	})

	router.GET("/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create.html", gin.H{
			"title":   "Post for creating",
			"content": "First title",
		})
	})

	router.POST("/posts", func(c *gin.Context) {
		title := c.PostForm("title")
		content := c.PostForm("content")

		_, err := db.Exec(`INSERT INTO posts
						(title, content)
						VALUES (?, ?)
						 `, title, content)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"error": err.Error(),
			})
			return
		}
		c.Redirect(http.StatusFound, "/")
	})

	router.GET("/about_me", func(c *gin.Context) {
		c.HTML(200, "about_me.html", gin.H{
			"title": "Обо мне",
		})
	})

	router.GET("/test", func(c *gin.Context) {
		testData := gin.H{
			"title": "First page",
			"items": []string{"hello world", "second hello world", "third"},
		}
		c.HTML(200, "test.html", testData)
	})

	router.GET("/contacts", func(c *gin.Context) {
		data := gin.H{
			"title": "Контакты",
		}
		c.HTML(200, "contacts.html", data)
	})
}

func getPosts(db *sql.DB) ([]Post, error) {
	rows, err := db.Query("SELECT id, title, content FROM posts")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.ID, &p.Title, &p.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}
