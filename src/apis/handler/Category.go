package handler

import (
	"fmt"
	"net/http"
)

func CategoryC(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CategoryC")
}

func CategoryR(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CategoryR")
}

func CategoryU(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CategoryU")
}

func CategoryD(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CategoryD")
}

func CategoryL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CategoryL")
}

func CategoryS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CategoryS")
}
