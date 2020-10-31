package crawler

import(
	"fmt"
	"net/http"
	"crypto/tls"

	io "vinh.crawler/ultilities"
	"github.com/steelx/extractlinks"

)

func ExtractLinks(usingUrl string){

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

	io.CheckError(err)
	defer response.Body.Close()

	links, err := extractlinks.All(response.Body)
	io.CheckError(err)
	for i, link := range links{
		fmt.Printf("Index: %v\tLink: %v\n",i, link)
	}
}