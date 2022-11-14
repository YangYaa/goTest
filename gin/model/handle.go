package model

import "github.com/gin-gonic/gin"

func AddHandler(c *gin.Context, profile DbModel) {
	c.ShouldBindJSON(profile)
}
