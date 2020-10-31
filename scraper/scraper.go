package scraper

import(
	
	"io/ioutil"
	"net/http"
	io "vinh.crawler/ultilities"

)

func GetListing (url string) string{
	response, err := http.Get(url)
	io.CheckError(err)
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}