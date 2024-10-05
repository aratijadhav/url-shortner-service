package routeutils

import (
	"github.com/aratijadhav/url-shortner-service/rest"
	"github.com/aratijadhav/url-shortner-service/routes"
)

func CreateRoute() HandlerStruct {

	var HS = HandlerStruct{
		Url:         routes.API + routes.CREATEURL + "/{urlname}",
		HandlerName: rest.CreateShortURL,
		Methods:     []string{"GET"},
	}

	return HS

}
