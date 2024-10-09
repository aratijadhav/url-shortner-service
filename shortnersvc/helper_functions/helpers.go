package helperfunctions

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var Shorturls = make(map[string]string)
var Originalurls = make(map[string]map[string]int)

type URLCount struct {
	Originalurl string
	URL         string
	Count       int
}

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

func sortOriginalUrls() map[string]int {
	var sortedURLs []URLCount
	var mostvisitedurl = make(map[string]int)

	for originalurl, short_urls := range Originalurls {
		for url, count := range short_urls {
			sortedURLs = append(sortedURLs, URLCount{Originalurl: originalurl, URL: url, Count: count})
		}
	}
	sort.Slice(sortedURLs, func(i, j int) bool {
		return sortedURLs[i].Count > sortedURLs[j].Count
	})

	fmt.Println("Sorted URLs:")

	for idx := 0; idx < len(sortedURLs); idx++ {
		if idx == 3 {
			break
		}
		item := sortedURLs[idx]
		mostvisitedurl[item.Originalurl] = item.Count
		fmt.Printf("originalurl: %s, Count: %d\n", item.Originalurl, item.Count)

	}
	return mostvisitedurl

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

func GetMostVisited() map[string]int {
	return sortOriginalUrls()
}

func GetAllURL() map[string]int {
	allurl := make(map[string]int)
	for originalurl, short_urls := range Originalurls {
		for _, count := range short_urls {
			allurl[originalurl] = count
		}
	}
	return allurl

}
