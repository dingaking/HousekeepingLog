package handler

import (
	"fmt"
	"net/http"
)

func SystemC(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SystemC")
}
