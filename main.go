package main

import (
	"gim"
	"net/http"
)

func main() {
	r := gim.New()
	r.GET("/", func(c *gim.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gim.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.GET("/hello/:name", func(c *gim.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gim.Context) {
		c.JSON(http.StatusOK, gim.H{"filepath": c.Param("filepath")})
	})
	r.POST("/login", func(c *gim.Context) {
		c.JSON(http.StatusOK, gim.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
