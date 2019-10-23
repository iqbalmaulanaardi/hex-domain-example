package rest

import (
	"net/http"

	"github.com/iqbalmaulanaardi/hex-domain-example/hex-domain/pkg/listing"
	"github.com/iqbalmaulanaardi/hex-domain-example/hex-domain/pkg/utils/responses"
	"github.com/julienschmidt/httprouter"
)

func GetUsersHandler(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		list, err := s.GetUsersService()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
		}
		responses.JSON(w, http.StatusOK, list)
	}
}
