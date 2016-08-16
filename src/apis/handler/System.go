package handler

import (
	"fmt"
	"net/http"
)

func SystemR(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SystemR")
}
