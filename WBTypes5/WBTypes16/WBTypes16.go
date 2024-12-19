package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CookieTool() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get cookie
		if cookie, err := c.Cookie("WB"); err == nil {
			if cookie == "ok" {
				c.Next()
				return
			}
		}

		// Cookie verification failed
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden with no cookie"})
		c.Abort()
	}
}

func main() {
	route := gin.Default()

	// Встроенное middleware для логирования запросов
	// Gin по умолчанию включает это в `gin.Default()`
	//route.Use(gin.Logger())

	// Встроенное middleware для восстановления от паники
	// Gin по умолчанию включает это в `gin.Default()`
	//route.Use(gin.Recovery())

	route.GET("/login", func(c *gin.Context) {
		// Set cookie {"WB": "ok" }, maxAge 20 seconds.
		c.SetCookie("WB", "ok", 20, "/", "localhost", false, true)
		c.String(200, "Login success!")
	})

	route.GET("/home", CookieTool(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "WB home page"})
	})

	route.Run(":8081")
}
