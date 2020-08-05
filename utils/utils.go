package utils

import (
	"log"
	"net/http"
)

// CheckErr func
func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// CheckCode func
func CheckCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Status code err: %d %s", res.StatusCode, res.Status)
	}
}
