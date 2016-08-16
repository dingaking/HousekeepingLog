package handler

import (
	"fmt"
	"net/http"
)

func CapitalC(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CapitalC")
}
