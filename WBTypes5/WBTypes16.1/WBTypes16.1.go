package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Custom logging middleware
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Логируем информацию о запросе
		log.Printf("Received request: %s %s", c.Request.Method, c.Request.URL.Path)

		// Обрабатываем запрос
		c.Next()

		// Логируем информацию об ответе
		log.Printf("Sent response: %d", c.Writer.Status())
	}
}

// CookieTool  middleware для проверки наличия и значения cookie
func CookieTool() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем cookie
		if cookie, err := c.Cookie("WB"); err == nil {
			if cookie == "ok" {
				c.Next()
				return
			}
		}

		// Проверка cookie не удалась
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden with no cookie"})
		c.Abort()
	}
}

func main() {
	// Создаем новый экземпляр Gin с логированием
	route := gin.New()

	// Добавляем кастомное middleware для логирования
	route.Use(LoggingMiddleware())

	// Определяем маршрут для входа и установки cookie
	route.GET("/login", func(c *gin.Context) {
		// Устанавливаем cookie {"WB": "ok" }, maxAge 20 секунд
		c.SetCookie("WB", "ok", 20, "/", "localhost", false, true)
		c.String(200, "Login success!")
	})

	// Определяем маршрут для домашней страницы с проверкой cookie
	route.GET("/home", CookieTool(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "WB home page"})
	})

	// Запускаем сервер на порту 8081
	route.Run(":8081")
}
