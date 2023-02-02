package uploadFile

import (
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"os"
)

var (
	uploadFileKey = "upload-key"
)

func GinUpload() {
	r := gin.Default()
	r.POST("/upload", uploadHandler)
	r.Run()
}
func uploadHandler(c *gin.Context) {
	header, err := c.FormFile(uploadFileKey)
	if err != nil {
		//ignore
	}
	dst := header.Filename
	// gin 简单做了封装,拷贝了文件流
	if err := SaveUploadedFile(header, dst); err != nil {
		// ignore
	}
}

// SaveUploadedFile uploads the form file to specific dst.
func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	//创建 dst 文件
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	// 拷贝文件
	_, err = io.Copy(out, src)
	return err
}
