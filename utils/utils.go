package utils

import (
	"log"
	"net/http"
	"os"
	"strings"
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

// SplitData func
func SplitData(original, split string) string {
	splitString := strings.Split(original, split)

	return splitString[len(splitString)-1]
}

// CreateDir func
func CreateDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir("imgFiles/"+path, os.FileMode(0775))
	}
}
