package handler

import (
	"fmt"
	"net/http"
)

func DeviceC(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeviceC")
}

func DeviceR(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeviceR")
}

func DeviceU(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeviceU")
}

func DeviceD(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeviceD")
}

func DeviceL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeviceL")
}

func DeviceS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeviceS")
}
