package handler

import (
	"fmt"
	"net/http"
)

func ItemC(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ItemC")
}
