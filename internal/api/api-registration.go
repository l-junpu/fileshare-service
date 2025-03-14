package api

import (
	"fileshare-service/internal/api/middleware"
	"net/http"
)

func InitializeApiRegistration() {
	http.HandleFunc("/api/register", middleware.HandleRegistration)
	http.HandleFunc("/api/login", middleware.HandleLogin)
	http.HandleFunc("/api/logout", middleware.HandleLogout)
	http.HandleFunc("/api/user/profile", middleware.HandleRetrieveBasicInfo)
}
