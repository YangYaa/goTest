package router

import (
	"github.com/gin-gonic/gin"
	"goTest/gin/model"
)

func HandGetMsg(c *gin.Context) {
	var v model.DbModel = new(model.TestDb)
	model.GetHandler(c, v)
}

func HandPutMsg(c *gin.Context) {
	var v model.DbModel = new(model.TestDb)
	model.AddHandler(c, v)
}

func HandPostMsg(c *gin.Context) {
	var v model.DbModel = new(model.TestDb)
	model.UpdateHandler(c, v)
}

func HandDeleteMsg(c *gin.Context) {
	var v model.DbModel = new(model.TestDb)
	model.DeleteHandler(c, v)
}
