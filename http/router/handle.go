package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goTest/gin/model"
	"net/http"
	"os"
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

func DownloadFileService(c *gin.Context) {
	fileDir := c.Query("fileDir")
	fileName := c.Query("fileName")
	//打开文件
	_, errByOpenFile := os.Open(fileDir + "/" + fileName)
	//非空处理
	fmt.Println(fileDir, fileName)
	if errByOpenFile != nil {
		c.Redirect(http.StatusFound, "/404")
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(fileDir + "/" + fileName)
	return
}
