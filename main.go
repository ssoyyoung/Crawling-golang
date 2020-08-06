package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	img "github.com/ssoyyoung.p/Crawling-golang/imgdown"
	parsing "github.com/ssoyyoung.p/Crawling-golang/parsing"
	utils "github.com/ssoyyoung.p/Crawling-golang/utils"
)

func main() {

	debug := false
	baseURL := "https://benito.co.kr"
	categoryURL := parsing.GetCategory(baseURL)

	for idx, cateURL := range categoryURL {
		fmt.Println(cateURL)
		crawlLink := baseURL + cateURL
		cateN := utils.SplitData(crawlLink, "=", 1)
		utils.CreateDir(cateN)

		if debug && idx >= 1 {
			break
		}

		start := time.Now()
		resultList := parsing.GOgetProductPage(crawlLink)
		fmt.Println(crawlLink, ">>", len(resultList))

		for i, result := range resultList {
			if debug && i >= 1 {
				break
			}
			productNum := utils.SplitData(result.URL, "/", 6)
			productPath := cateN + "/" + productNum
			utils.CreateDir(productPath)
			img.ImgDownloading(result.ImgList, productPath)
		}
		elapsed := time.Since(start)
		log.Printf("Binomial took %s", elapsed)
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
