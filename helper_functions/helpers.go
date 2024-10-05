package helperfunctions

import (
	"fmt"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type OriginalUrlData struct {
	Shorturl1 string
	Count     int
}

var Shorturls = make(map[string]string)
var Originalurls = make(map[string]OriginalUrlData)

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

	if shortURL, ok := Originalurls[originalurl]; ok {
		Originalurls[originalurl] = OriginalUrlData{
			Shorturl1: shortURL.Shorturl1,
			Count:     shortURL.Count + 1,
		}
		fmt.Printf("%s URL is already shortened, shortened URL: %s", originalurl, shortURL.Shorturl1)
		fmt.Println()
		fmt.Println(Originalurls)
		return shortURL.Shorturl1, true
	} else {
		genString := randStringBytes(5)
		Originalurls[originalurl] = OriginalUrlData{
			Shorturl1: genString,
			Count:     1,
		}
		Shorturls[genString] = originalurl
		fmt.Println()
		fmt.Println(Originalurls)
		return genString, false
	}

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
