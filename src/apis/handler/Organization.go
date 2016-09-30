package handler

import (
	"fmt"
	"net/http"
)

func OrganizationR(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OrganizationR")
}

func OrganizationL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OrganizationL")
}

func OrganizationS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OrganizationS")
}
