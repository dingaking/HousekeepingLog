package handler

import (
	"fmt"
	"net/http"
)

func CapitalC(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CapitalC")
}

func CapitalR(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CapitalR")
}

func CapitalU(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CapitalU")
}

func CapitalD(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CapitalD")
}

func CapitalL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CapitalL")
}

func CapitalS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CapitalS")
}
