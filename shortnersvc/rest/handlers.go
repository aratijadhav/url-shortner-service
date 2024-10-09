package rest

import (
	"fmt"
	"net/http"

	helperfunctions "github.com/aratijadhav/url-shortner-service/helper_functions"
	"github.com/gorilla/mux"
)

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	urlName := mux.Vars(r)["urlname"]
	shortUrl, isGenerated := helperfunctions.GenerateShortUrl(urlName)
	responseMsg := ""

	w.Header().Set("Content-Type", "application/txt")
	if isGenerated {
		w.WriteHeader(http.StatusCreated)
		responseMsg = fmt.Sprintf("Shortned URL for %s is: http://localhost:8080/%s", urlName, shortUrl)

	} else {
		w.WriteHeader(http.StatusFound)
		responseMsg = fmt.Sprintf("Shortned URL for %s already exists: http://localhost:8080/%s ", urlName, shortUrl)

	}
	w.Write([]byte(responseMsg))
}

func RedirectOriginalURL(w http.ResponseWriter, r *http.Request) {
	shorturlname := mux.Vars(r)["shortpath"]

	if original_url, ok := helperfunctions.Shorturlsl[shorturlname]; ok {
		urlinfoData := helperfunctions.URLInfo[original_url]
		urlinfoData.Visit += 1
		helperfunctions.URLInfo[original_url] = urlinfoData

		orignalurl := helperfunctions.GetExternalUrl(original_url)
		helperfunctions.GenerateMostVisited(original_url, helperfunctions.URLInfo[original_url].Visit)

		w.Header().Set("Content-Type", "application/txt")
		http.Redirect(w, r, orignalurl, http.StatusTemporaryRedirect)
	} else {
		w.Header().Set("Content-Type", "application/txt")
		w.WriteHeader(http.StatusBadRequest)
		errorMsg := "Error: Incorrect Short URL"
		w.Write([]byte(errorMsg))
	}

}

func MostvisitedURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/txt")

	result := helperfunctions.GetMostVisited()
	msg := fmt.Sprintf("MostvisitedURL are: %v", result)
	w.Write([]byte(msg))

}

// func GetAllURL(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/txt")

// 	allurl := helperfunctions.GetAllURL()
// 	msg := fmt.Sprintf("All URLs are: %v", allurl)
// 	w.Write([]byte(msg))

// }
