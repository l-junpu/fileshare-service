package main

import (
	"fileshare-service/internal/api"
	"fmt"
	"net/http"
)

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Testing post request")
}

func main() {
	api.InitializeApiRegistration()
	http.ListenAndServe(":8080", nil)
}
