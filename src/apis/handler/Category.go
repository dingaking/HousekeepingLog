package handler

import (
	"fmt"
	"net/http"
)

func CategoryC(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CategoryC")
}
