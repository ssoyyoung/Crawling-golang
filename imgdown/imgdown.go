package imgdown

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	parsing "github.com/ssoyyoung.p/Crawling-golang/parsing"
	"github.com/ssoyyoung.p/Crawling-golang/utils"
)

// FullProcess func
func FullProcess(cateURL string, c1 chan<- string) {

	baseURL := "https://benito.co.kr"
	debug := false
	fmt.Println("Starting Crawling", cateURL)
	crawlLink := baseURL + cateURL
	cateN := utils.SplitData(crawlLink, "=", 1)
	utils.CreateDir(cateN)

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
		ImgDownloading(result.ImgList, productPath)
	}
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

	c1 <- cateURL

}

// ImgDownloading func
func ImgDownloading(imgList []string, productPath string) {
	var wg sync.WaitGroup

	fileName := "imgFiles/error.txt"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	utils.CheckErr(err)

	for _, imgURL := range imgList {
		imgName := "imgFiles/" + productPath + "/" + utils.SplitData(imgURL, "/", 1)
		url := "https://benito.co.kr" + imgURL

		go func(url string) {
			wg.Add(1)
			http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
			res, err := http.Get(url)

			if err != nil {
				fmt.Println("http get error: ", url, "-", err)
				if _, err := file.WriteString(productPath + ", " + url + "\n"); err != nil {
					fmt.Println(">>>Writing error!", url)
				}
			} else {
				output, err := os.Create(imgName)
				if err != nil {
					fmt.Println("Error while createing", imgName, "-", err)
				}
				_, err = io.Copy(output, res.Body)
				output.Close()
				res.Body.Close()

				if err != nil {
					fmt.Println("Error while Downloading", url, "-", err)
					if _, err := file.WriteString(productPath + ", " + url + "\n"); err != nil {
						fmt.Println(">>>Writing error!", url)
					}
				}
			}
			wg.Done()
		}(url)
	}
	wg.Wait()
	fmt.Println("[DONE] downloaded img", productPath, ">>", len(imgList))
}
