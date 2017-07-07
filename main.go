package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robertoduessmann/goEinstein/controller"
)

func main() {
	route := mux.NewRouter()
	route.Path("/einsten/{mass}").Methods(http.MethodGet).HandlerFunc(controller.Energy)

	if err := http.ListenAndServe(":3000", route); err != nil {
		log.Fatal(err)
	}
}
