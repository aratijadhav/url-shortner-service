package routeutils

import (
	"github.com/aratijadhav/url-shortner-service/rest"
	"github.com/aratijadhav/url-shortner-service/routes"
)

func MostvisitedURLRoute() HandlerStruct {
	var HS = HandlerStruct{
		Url:         routes.API + routes.GETMOSTVISITED,
		HandlerName: rest.MostvisitedURL,
		Methods:     []string{"GET"},
	}
	return HS
}

func CreateRoute() HandlerStruct {
	var HS = HandlerStruct{
		Url:         routes.API + routes.CREATEURL + "/{urlname}",
		HandlerName: rest.CreateShortURL,
		Methods:     []string{"GET"},
	}
	return HS
}

func RedirectOriginalURL() HandlerStruct {
	var HS = HandlerStruct{
		Url:         "/{shortpath}",
		HandlerName: rest.RedirectOriginalURL,
		Methods:     []string{"GET"},
	}
	return HS
}
