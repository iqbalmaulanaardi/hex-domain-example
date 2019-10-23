package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iqbalmaulanaardi/hex-domain-example/hex-domain/pkg/http/rest"
	"github.com/iqbalmaulanaardi/hex-domain-example/hex-domain/pkg/listing"
	"github.com/iqbalmaulanaardi/hex-domain-example/hex-domain/pkg/storage/database"
)

func main() {
	var lister listing.Service
	s, _ := database.NewStorage()
	lister = listing.NewService(s)
	router := rest.Handler(lister)

	fmt.Println("The local server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
