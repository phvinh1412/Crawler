package crawler

import(
	"fmt"
	"net/http"
	"net/url"
	"crypto/tls"
	io "vinh.crawler/ultilities"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type CrawlResult struct {
	URL		string
	Title	string
	H1		string
}

type Parser interface {
	ParsePage(*goquery.Document) CrawlResult
}

func extractLinks(doc *goquery.Document) []string{
	foundUrls := []string{}
	if doc!= nil {
		doc.Find("a").Each(func(i int, s *goquery.Selection){
			res, _ := s.Attr("href")
			foundUrls = append(foundUrls, res)
		})
		return foundUrls
	}
	fmt.Printf("No fucking link\n")
	return foundUrls
}

func resolveRelative(baseURL string, hrefs []string) []string {
	internalUrls := []string{}
 
	for _, href := range hrefs {
		if strings.HasPrefix(href, baseURL) {
			internalUrls = append(internalUrls, href)
		}
 
		if strings.HasPrefix(href, "/") {
			resolvedURL := fmt.Sprintf("%s%s", baseURL, href)
			internalUrls = append(internalUrls, resolvedURL)
		}
	}

	return internalUrls
}

func getRequest (url string) (*http.Response, error) {
	tls_config := &tls.Config{
		InsecureSkipVerify: true,
	}
	transport := &http.Transport{
		TLSClientConfig: tls_config,
	}
	client := &http.Client{
		Transport: transport,
	}
	response, err := client.Get(url)

	io.CheckError(err)
	defer response.Body.Close()
	return response, err
}

func crawlPage(baseURL, targetURL string, parser Parser, token chan struct{}) ([]string, CrawlResult) {
 
	token <- struct{}{}
	fmt.Println("Requesting: ", targetURL)
	resp, _ := getRequest(targetURL)
	<-token
 
	doc, _ := goquery.NewDocumentFromResponse(resp)
	pageResults := parser.ParsePage(doc)
	links := extractLinks(doc)
	foundUrls := resolveRelative(baseURL, links)
 
	return foundUrls, pageResults
}

func parseStartURL(str string) string {
	parsed, _ := url.Parse(str)
	return fmt.Sprintf("%s://%s", parsed.Scheme,parsed.Host)
}

func Crawl(startURL string, parser Parser, concurrency int) []CrawlResult {
	results := []CrawlResult{}
	worklist := make(chan []string)
	var n int
	n++
	var tokens = make(chan struct{}, concurrency)
	go func() { worklist <- []string{startURL} }()
	seen := make(map[string]bool)
	baseDomain := parseStartURL(startURL)
 
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(baseDomain, link string, parser Parser, token chan struct{}) {
					foundLinks, pageResults := crawlPage(baseDomain, link, parser, token)
					results = append(results, pageResults)
					if foundLinks != nil {
						worklist <- foundLinks
					}
				}(baseDomain, link, parser, tokens)
			}
		}
	}
	return results
}