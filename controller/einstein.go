package controller

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Energy(w http.ResponseWriter, r *http.Request) {
	mass := getMass(r)
	fmt.Fprintf(w, "%f", calculate(mass))
}

func getMass(r *http.Request) float64 {
	mass := mux.Vars(r)["mass"]
	f, err := strconv.ParseFloat(mass, 64)
	if err != nil {
		fmt.Println("Erro na conversão")
	}
	return f
}

func calculate(mass float64) float64 {
	c := 3e8
	energy := math.Pow(c, 2) * mass
	return energy
}
