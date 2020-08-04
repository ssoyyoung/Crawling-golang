package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	debug := true

	baseURL := "https://benito.co.kr"
	categoryURL := getCategory(baseURL)

	for idx, cateURL := range categoryURL {
		crawlLink := baseURL + cateURL
		if debug && idx >= 1 {
			break
		}
		fmt.Println(crawlLink)
		getProductPage(crawlLink)

	}
}

func getCategory(baseURL string) []string {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	categoryURL := []string{}
	doc.Find(".gnb > div.wrapper > ul > li > a").Each(func(i int, s *goquery.Selection) {
		link, success := s.Attr("href")

		if success == true && strings.Contains(link, "cate_no") == true {
			categoryURL = append(categoryURL, link)
		}
	})

	return categoryURL
}

func getProductPage(baseURL string) []string {
	pages := getPagesCount(baseURL)
	productPageList := []string{}

	for i := 1; i < pages+1; i++ {
		productURL := baseURL + "&page=" + strconv.Itoa(i)

		productItemList := getProductList(productURL)

		for idx, productURL := range productItemList {
			if idx >= 2 {
				continue
			}
			imgList := getAllImageInProduct(productURL)
			fmt.Println(imgList)
		}
	}

	return productPageList
}

func getAllImageInProduct(productURL string) []string {
	fmt.Println(productURL)
	res, err := http.Get(productURL)
	checkCode(res)
	checkErr(err)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	imgList := []string{}

	doc.Find("#prdDetail div img").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")

		if strings.Contains(src, "intop") == false {
			imgList = append(imgList, src)
		}
	})

	// doc.Find("#prdDetail > div > div > img").Each(func(i int, s *goquery.Selection) {
	// 	fmt.Println(s.Attr("src"))
	// })

	// doc.Find("#prdDetail > div > div > p > img").Each(func(i int, s *goquery.Selection) {
	// 	fmt.Println(s.Attr("src"))
	// })

	// doc.Find("#prdDetail > div > div > div > p > img").Each(func(i int, s *goquery.Selection) {
	// 	fmt.Println(s.Attr("src"))
	// })

	return imgList
}

func getProductList(productURL string) []string {
	res, err := http.Get(productURL)
	checkCode(res)
	checkErr(err)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	productItemList := []string{}

	doc.Find(".thumbnail > a").Each(func(i int, s *goquery.Selection) {
		//s.Find("xans-record- > thumbnail > a")
		link, success := s.Attr("href")
		if success == true {
			link = "https://benito.co.kr" + link
			productItemList = append(productItemList, link)
		}
	})

	return productItemList
}

func getPagesCount(baseURL string) int {
	pages := 0
	res, err := http.Get(baseURL)
	checkCode(res)
	checkErr(err)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find("a.last").Each(func(i int, s *goquery.Selection) {
		link, success := s.Attr("href")
		if success == true {
			num := strings.Split(link, "page=")
			if num[len(num)-1] == "#none" {
				pages = 1
			} else {
				pages, err = strconv.Atoi(num[len(num)-1])
				checkErr(err)
			}
		}
	})

	return pages
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
