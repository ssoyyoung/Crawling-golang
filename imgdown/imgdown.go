package imgdown

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/ssoyyoung.p/Crawling-golang/utils"
)

// ImgDownloading func
func ImgDownloading(imgList []string, productPath string) {
	fileName := "imgFiles/error.txt"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	utils.CheckErr(err)

	fmt.Println("starting")

	for _, imgURL := range imgList {
		imgName := "imgFiles/" + productPath + "/" + utils.SplitData(imgURL, "/", 1)
		url := "https://benito.co.kr" + imgURL
		fmt.Println(">>", imgName, url)

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
					if _, err := file.WriteString(url + "\n"); err != nil {
						fmt.Println(">>>Writing error!", url)
					}
				}
			}
			wg.Done()
		}(url)
	}
	wg.Wait()
}

var wg sync.WaitGroup

// func downloading(errorFile *os.File, url string, imgName string) {
// 	wg.Add(1)
// 	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
// 	res, err := http.Get(url)

// 	if err != nil {
// 		fmt.Println("http get error: ", url, "-", err)
// 		if _, err := errorFile.WriteString(url + "\n"); err != nil {
// 			fmt.Println(">>>Writing error!", url)
// 		}
// 	} else {
// 		output, err := os.Create(imgName)
// 		if err != nil {
// 			fmt.Println("Error while createing", imgName, "-", err)
// 		}
// 		_, err = io.Copy(output, res.Body)
// 		output.Close()
// 		res.Body.Close()

// 		if err != nil {
// 			fmt.Println("Error while Downloading", url, "-", err)
// 			if _, err := errorFile.WriteString(url + "\n"); err != nil {
// 				fmt.Println(">>>Writing error!", url)
// 			}
// 		}
// 	}
// 	wg.Done()
// }
