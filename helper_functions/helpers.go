package helperfunctions

import (
	"fmt"
	"math/rand"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type OriginalUrlData struct {
	Shorturl1 string
	Count     int
}

var Shorturls = make(map[string]string)

// var Originalurls = make(map[string]OriginalUrlData)
var Originalurls = make(map[string]map[string]int)

/*
Local function
*/
func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func checkIfOriginalUrlIsAlreadyShorten(originalurl string) (string, bool) {
	genString := randStringBytes(5)
	var short_url = genString
	if shortURL, ok := Originalurls[originalurl]; ok {
		for url, val := range shortURL {
			short_url = url
			Originalurls[originalurl][url] = val + 1
			fmt.Printf("%s URL is already shortened, shortened URL: %s", originalurl, url)
		}
		return short_url, true
	} else {
		Originalurls[originalurl] = make(map[string]int)
		Originalurls[originalurl][genString] = 1
	}

	Shorturls[genString] = originalurl
	return short_url, false

}

/*
Global Functions
*/
func GenerateShortUrl(url string) (string, bool) {
	shortURL, isExists := checkIfOriginalUrlIsAlreadyShorten(url)
	if isExists {
		return shortURL, false
	}

	return shortURL, true
}

func GetExternalUrl(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	fmt.Println(url)
	return url
}
