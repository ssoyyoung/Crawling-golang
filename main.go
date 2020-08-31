package main

import (
	img "github.com/ssoyyoung.p/Crawling-golang/imgdown"
	parsing "github.com/ssoyyoung.p/Crawling-golang/parsing"
)

func main() {

	baseURL := "https://benito.co.kr"
	categoryURL := parsing.GetCategory(baseURL)

	var resultList []string
	c1 := make(chan string)

	for _, cateURL := range categoryURL {
		go img.FullProcess(cateURL, c1)
	}

	for i := 0; i < len(categoryURL); i++ {
		result := <-c1
		//fmt.Println("[DONE] crawling category", result)
		resultList = append(resultList, result)

	}
}
