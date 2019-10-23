package rest

import (
	"net/http"

	"github.com/iqbalmaulanaardi/hex-domain-example/hex-domain/pkg/listing"
	"github.com/julienschmidt/httprouter"
)

func Handler(l listing.Service) http.Handler {
	router := httprouter.New()
	router.GET("/users", GetUsersHandler(l))
	return router
}
