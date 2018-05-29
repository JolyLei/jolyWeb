package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"os"
	"log"
	"io"
)

func GetHandler(c *gin.Context){
	value, exist := c.GetQuery("key")
	if !exist{
		value = "the key is not exist!"
	}
	c.Data(http.StatusOK,"text/plain",[]byte(fmt.Sprintf("get succeed\n %s",value)))
	return
}

func PostHandler(c *gin.Context){

	id := c.PostForm("id")
	name := c.PostForm("name")
	c.JSON(http.StatusOK, gin.H{
		"status_code":http.StatusOK,
		"id" : id,
		"name" : name,
	})
	return
}

func PutHandler(c *gin.Context){
	c.Data(http.StatusOK, "text/plain", []byte("put success"))
	return
}

func DeleteHandler(c *gin.Context){
	c.Data(http.StatusOK, "text/plain", []byte("delete succeed\n"))
	return
}

func UploadFile(c *gin.Context){
	name := c.PostForm("name")
	fmt.Println(name)
	file, header, err := c.Request.FormFile("upload")
	if err != nil{
		c.String(http.StatusBadRequest,"Bad request")
		return
	}
	filename := header.Filename
	fmt.Println(file,err,filename)

	out,err := os.Create(filename)
	if err != nil{
		log.Fatal(err)
	}
	defer out.Close()
	_,err = io.Copy(out, file)
	if err != nil{
		log.Fatal(err)
	}
	c.String(http.StatusCreated, "upload successful")
}