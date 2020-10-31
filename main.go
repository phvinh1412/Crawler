package main

import (
	//"crypto/tls"
	"fmt"
	//"io/ioutil"
	//"net/http"
	io "vinh.crawler/ultilities"
	scraper "vinh.crawler/scraper"
	//"github.com/steelx/extractlinks"
)

func init() {  
	println("main package initialized")
}

func main() {
	url := io.GetText()
	body := scraper.GetListing (url)
	fmt.Printf(body)

}