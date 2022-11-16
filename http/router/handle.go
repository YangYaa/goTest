package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func HandPutMsg(c *gin.Context) {
	//var v model.DbModel = new(model.TestDb)
	//model.AddHandler(c, v)
	err := c.ShouldBindJSON(&TestPutName)
	if err != nil {
		fmt.Println("The PUT function not have json", err)
		return
	}
	fmt.Println("The TestPutName is ", TestPutName)
}
