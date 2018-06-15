package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"net/url"
	"net"
	"time"
)

func HttpGet(){
	resp, err := http.Get("http://www.baidu.com")
	if err != nil{
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err.Error())
	}

	fmt.Println(string(body))
}

func HttpPost(){
	resp, err := http.Post("http://www.baidu.com",
		"application/x-www-form-urlencode",
			strings.NewReader("name=joly"))
	if err != nil{
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func HttpPostForm(){
	resp,err := http.PostForm("http://www.baidu.com",
		url.Values{"key":{"a","b"}, "id":{"123"}})
	if err != nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(string(body))
}



func Gouzao() {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * 30,
		Transport: netTransport,
	}

	// get
	response, _ := netClient.Get("http://www.golangnote.com/")
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}

}


func HttpDo(){
	client := &http.Client{}

	req, err := http.NewRequest("POST","http://www.baidu.com", strings.NewReader("name= joly"))
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	req.Header.Set("Cookie","name=joly")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body,err:= ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func main(){
	HttpDo()
}