package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/register", func(c *gin.Context) {

	})

	router.POST("/login", func(c *gin.Context) {

	})

	router.POST("/exam", func(c *gin.Context) {

	})

	router.GET("/exams", func(c *gin.Context) {

	})

	router.POST("/generate-plan", func(c *gin.Context) {

	})

	router.POST("/studyplan", func(c *gin.Context) {

	})

}
