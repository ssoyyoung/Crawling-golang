package main

import (
	"fmt"
	"log"
	"net/http"

	parsing "github.com/ssoyyoung.p/Crawling-golang/parsing"
)

func main() {
	debug := true
	baseURL := "https://benito.co.kr"
	categoryURL := parsing.GetCategory(baseURL)

	for idx, cateURL := range categoryURL {
		crawlLink := baseURL + cateURL
		if debug && idx >= 1 {
			break
		}
		resultList := parsing.GetProductPage(crawlLink)
		for _, result := range resultList {
			fmt.Println(result.URL)
			fmt.Println(len(result.ImgList))
		}

	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Status code err: %d %s", res.StatusCode, res.Status)
	}
}
