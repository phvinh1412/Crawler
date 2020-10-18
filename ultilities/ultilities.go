package ultilities

import (
	"fmt"
	"crypto/tls"
	"net/http"
	"os"
	"github.com/steelx/extractlinks"
)

func init() {  
    fmt.Println("Ultilities package initialized")
}

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

	CheckError(err)
	defer response.Body.Close()

	links, err := extractlinks.All(response.Body)
	CheckError(err)
	for i, link := range links{
		fmt.Printf("Index: %v\tLink: %v\n",i, link)
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Print("Error: ", err)
		os.Exit(1)
	}
}