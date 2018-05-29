package main

import(
	"github.com/gin-gonic/gin"
	"joly/routers"
	"fmt"
	"net/http"
)


func main(){
	gin.SetMode(gin.DebugMode)	//全局设置为开发环境
	router := gin.Default()		//获得路由实例

	//添加中间件
	router.Use(Middleware)
	//注册接口
	router.GET("/simple/server/get", routers.GetHandler)
	router.POST("/simple/server/post", routers.PostHandler)
	router.PUT("/simple/server/put", routers.PutHandler)
	router.DELETE("/simple/server/delete", routers.DeleteHandler)
	router.POST("/simple/server/upload",routers.UploadFile)

	//监听端口
	http.ListenAndServe(":8099",router)

}


func Middleware(c *gin.Context){
	fmt.Println("this is a middleware!")
}