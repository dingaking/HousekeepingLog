package handler

import (
	"fmt"
	"net/http"
)

func ItemC(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ItemC")
}

func ItemR(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ItemR")
}

func ItemU(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ItemU")
}

func ItemD(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ItemD")
}

func ItemL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ItemL")
}

func ItemS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ItemS")
}
