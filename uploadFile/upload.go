package uploadFile

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

/*
显示欢迎页upload.html
*/
func welcome(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../uploadFile/upload.html")
	t.Execute(rw, nil)
}

/*
文件上传
*/
func upload(rw http.ResponseWriter, r *http.Request) {
	//获取普通表单数据
	username := r.FormValue("username")
	fmt.Println(username)
	//获取文件流,第三个返回值是错误对象
	file, header, _ := r.FormFile("photo")
	//读取文件流为[]byte
	b, _ := ioutil.ReadAll(file)
	//把文件保存到指定位置
	ioutil.WriteFile("/home/newfile.png", b, 0777)
	//输出上传时文件名
	fmt.Println("上传文件名:", header.Filename)
}

func UpLoadFile() {
	server := http.Server{Addr: "192.168.50.137:8080"}

	http.HandleFunc("/", welcome)
	http.HandleFunc("/upload", upload)

	server.ListenAndServe()
}
