package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEnergy(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:3000/einsten/5", nil)
	if err != nil {
		fmt.Println("Error in line 14")
	}

	w := httptest.NewRecorder()
	Equation(w, req)

	Convey("Test access controller equation", t, func() {
		Convey("Connection Succes", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})

}

func TestCalculate(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:3000/einsten", nil)
	if err != nil {
		fmt.Println("Error in line 31")
	}

	w := httptest.NewRecorder()
	Energy(w, req)

	Convey("Test access controller Energy", t, func() {
		Convey("Connection Succes", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})

}
