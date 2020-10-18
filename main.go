package main

import (
	//"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	io "vinh.crawler/ultilities"
	//"github.com/steelx/extractlinks"
)

func init() {  
	println("main package initialized")
}

func MakeRequest( url string) string {
	response, err := http.Get(url)
	defer response.Body.Close()
	io.CheckError(err)

	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}

func main() {
	fmt.Printf("Type using url: ")
	var usingUrl string
	fmt.Scanln(&usingUrl) 

	body := MakeRequest (usingUrl)
	fmt.Printf(body)

}

