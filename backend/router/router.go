package router

import (
	"backend/controllers"
	"backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	api := r.Group("/api")
	api.GET("/exchangeRates", controllers.GetExchangeRate)
	api.Use(middlewares.AuthMiddleware())
	{
		api.POST("/exchangeRates", controllers.CreateExchangeRate)
		api.POST("/articles", controllers.CreateArticle)
		api.GET("/articles", controllers.GetArticles)
		api.GET("/articles/:id", controllers.GetArticlesByID)
		api.POST("/articles/:id/like", controllers.LikeArticle)
		api.GET("articles/:id/like", controllers.GetArticleLikes)
	}
	return r
}
