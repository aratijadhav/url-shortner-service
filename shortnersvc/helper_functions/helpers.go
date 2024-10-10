package helperfunctions

import (
	"fmt"
	"math/rand"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type URLs struct {
	Shorturl string
	Visit    int
}

var URLInfo = make(map[string]URLs)
var mostvisited = make(map[string]int)
var Shorturlsl = make(map[string]string)

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func checkIfOriginalUrlIsAlreadyShorten(originalurl string) (string, bool) {
	genString := randStringBytes(5)
	if original_url, ok := URLInfo[originalurl]; ok {
		URLInfo[originalurl] = URLs{
			Shorturl: original_url.Shorturl,
			Visit:    original_url.Visit,
		}
		fmt.Printf("%s URL is already shortened, shortened URL: %s", originalurl, original_url.Shorturl)
		return original_url.Shorturl, true
	} else {
		URLInfo[originalurl] = URLs{
			Shorturl: genString,
			Visit:    0,
		}

	}
	Shorturlsl[genString] = originalurl
	return genString, false
}

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
	return url
}

func GenerateMostVisited(originalURL string, visits int) {

	var leastvisitedurl = originalURL
	var leastvisistedcount = visits
	if len(mostvisited) == 3 {
		alreadyMostVisited := false
		for url, visitCount := range mostvisited {
			if url == originalURL {
				mostvisited[originalURL] = visits
				alreadyMostVisited = true
			} else if visitCount < visits {
				leastvisitedurl = url
				leastvisistedcount = visitCount
			}
		}
		if leastvisistedcount < visits && !alreadyMostVisited {
			delete(mostvisited, leastvisitedurl)
			mostvisited[originalURL] = visits
		}
	} else {
		mostvisited[originalURL] = visits
	}
}

func GetMostVisited() map[string]int {
	return mostvisited
}
