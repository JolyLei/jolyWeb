package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main_1(){
	router := gin.Default()
	router.GET("/welcome", func(c *gin.Context) {
		c.String(200,"hello here")
	})

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick","anonymous")

		c.JSON(http.StatusOK,gin.H{
			"status":gin.H{
				"status_code":http.StatusOK,
				"status": "ok",
			},
			"message":message,
			"nick": nick,
		})


	})

	router.Run(":8888")
}