package main

import (
	"fmt"
	"log"
	"net/http"

	img "github.com/ssoyyoung.p/Crawling-golang/imgdown"
	parsing "github.com/ssoyyoung.p/Crawling-golang/parsing"
	utils "github.com/ssoyyoung.p/Crawling-golang/utils"
)

func main() {
	debug := true
	baseURL := "https://benito.co.kr"
	categoryURL := parsing.GetCategory(baseURL)

	for idx, cateURL := range categoryURL {
		fmt.Println(cateURL)
		crawlLink := baseURL + cateURL
		cateN := utils.SplitData(crawlLink, "=")
		utils.CreateDir(cateN)

		if debug && idx >= 1 {
			break
		}

		resultList := parsing.GetProductPage(crawlLink)
		for i, result := range resultList {
			if debug && i >= 1 {
				break
			}
			fmt.Println(result.URL)
			fmt.Println(len(result.ImgList))
			img.ImgDownloading(result.ImgList, cateN)
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
