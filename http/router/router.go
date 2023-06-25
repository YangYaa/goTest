package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

func NewRoutes() Routes {
	routes := Routes{
		{
			Method:      "PUT",
			Pattern:     "/testRoute/test",
			HandlerFunc: HandPutMsg,
		},
		{
			Method:      "POST",
			Pattern:     "/testRoute/test",
			HandlerFunc: HandPostMsg,
		},
		{
			Method:      "DELETE",
			Pattern:     "/testRoute/test",
			HandlerFunc: HandDeleteMsg,
		},
		{
			Method:      "GET",
			Pattern:     "/downloadFiles",
			HandlerFunc: DownloadFileService,
		},
	}
	return routes
}

func NewRouter(routes Routes) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	//set gin log function
	path := "/home/gin.log"
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("The open gin.log failed")
	}
	gin.DefaultWriter = io.MultiWriter(f)
	router.Use(gin.Logger(), gin.Recovery())

	//range route arrary
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}
	return router
}
