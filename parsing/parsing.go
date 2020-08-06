package parsing

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	m "github.com/ssoyyoung.p/Crawling-golang/models"
	"github.com/ssoyyoung.p/Crawling-golang/utils"
)

// GetCategory : [1] get all category
func GetCategory(baseURL string) []string {
	res, err := http.Get(baseURL)
	utils.CheckErr(err)
	utils.CheckCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	utils.CheckErr(err)

	categoryURL := []string{}
	doc.Find(".gnb > div.wrapper > ul > li > a").Each(func(i int, s *goquery.Selection) {
		link, success := s.Attr("href")

		if success == true && strings.Contains(link, "cate_no") == true {
			categoryURL = append(categoryURL, link)
		}
	})

	return categoryURL
}

// GOgetProductPage : [2] get product category
func GOgetProductPage(baseURL string) []m.Data {
	pages := GetPagesCount(baseURL)

	var resultList []m.Data
	c := make(chan m.Data)

	for i := 1; i < pages+1; i++ {
		productURL := baseURL + "&page=" + strconv.Itoa(i)
		productItemList := GetProductList(productURL)

		for _, productURL := range productItemList {
			go GOgetAllImageInProduct(productURL, c)
		}
		for i := 0; i < len(productItemList); i++ {
			result := <-c
			resultList = append(resultList, result)
		}
	}

	return resultList
}

// GetPagesCount : [2-1] get page count
func GetPagesCount(baseURL string) int {
	pages := 0
	res, err := http.Get(baseURL)
	utils.CheckCode(res)
	utils.CheckErr(err)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	utils.CheckErr(err)

	doc.Find("a.last").Each(func(i int, s *goquery.Selection) {
		link, success := s.Attr("href")
		if success == true {
			num := strings.Split(link, "page=")
			if num[len(num)-1] == "#none" {
				pages = 1
			} else {
				pages, err = strconv.Atoi(num[len(num)-1])
				utils.CheckErr(err)
			}
		}
	})

	return pages
}

// GetProductList : [2-2] get product list
func GetProductList(productURL string) []string {
	res, err := http.Get(productURL)
	utils.CheckCode(res)
	utils.CheckErr(err)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	utils.CheckErr(err)

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

// GOgetAllImageInProduct : [2-2] get all image in product
func GOgetAllImageInProduct(productURL string, c chan<- m.Data) {
	res, err := http.Get(productURL)
	utils.CheckCode(res)
	utils.CheckErr(err)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	utils.CheckErr(err)

	imgList := []string{}

	doc.Find("#prdDetail div img").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		if strings.Contains(src, "intop") == false {
			imgList = append(imgList, src)
		}
	})

	c <- m.Data{
		URL:     productURL,
		ImgList: imgList,
	}
}

/////////////// Not Use Goroutine ///////////////

// GetProductPage : [2] get product category
func GetProductPage(baseURL string) []m.Data {
	pages := GetPagesCount(baseURL)

	// 짧은 선언 사용불가
	var result m.Data
	var resultList []m.Data

	for i := 1; i < pages+1; i++ {
		productURL := baseURL + "&page=" + strconv.Itoa(i)

		productItemList := GetProductList(productURL)

		for _, productURL := range productItemList {
			result.URL = productURL
			result.ImgList = GetAllImageInProduct(productURL)
			resultList = append(resultList, result)
		}

	}

	return resultList
}

// GetAllImageInProduct : [2-2] get all image in product
func GetAllImageInProduct(productURL string) []string {
	res, err := http.Get(productURL)
	utils.CheckCode(res)
	utils.CheckErr(err)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	utils.CheckErr(err)

	imgList := []string{}

	doc.Find("#prdDetail div img").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		if strings.Contains(src, "intop") == false {
			imgList = append(imgList, src)
		}
	})

	return imgList
}
