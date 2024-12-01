package controllers

import (
	"backend/global"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
)

func LikeArticle(ctx *gin.Context) {
	articleID := ctx.Param("id")
	likeKey := "article:" + articleID + ":likes"
	if err := global.RedisDB.Incr(likeKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully liked article"})
}

func GetArticleLikes(ctx *gin.Context) {
	articleID := ctx.Param("id")
	likekey := "article:" + articleID + ":likes"
	likes, err := global.RedisDB.Get(likekey).Result()
	if err == redis.Nil {
		likes = "0"
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"likes": likes})

}
