package handler

import (
	"fmt"
	"net/http"
)

func DeviceC(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeviceC")
}
