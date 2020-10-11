package main

import (
	"crypto/tls"
	"fmt"
	//"io/ioutil"
	"net/http"
	"os"
	//io "vinh.crawler/ultilities"
	"github.com/steelx/extractlinks"
)

func init() {  
	println("main package initialized")
}

func main() {
	
	usingUrl := "https://vnexpress.net"

	tls_config := &tls.Config{
		InsecureSkipVerify: true,
	}
	transport := &http.Transport{
		TLSClientConfig: tls_config,
	}
	client := &http.Client{
		Transport: transport,
	}

	response, err := client.Get(usingUrl)

	checkError(err)
	defer response.Body.Close()

	links, err := extractlinks.All(response.Body)
	checkError(err)
	for i, link := range links{
		fmt.Printf("Index: %v\tLink: %v\n",i, link)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Print("Error: ", err)
		os.Exit(1)
	}
}
