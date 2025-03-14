package middleware

import (
	"encoding/json"
	"fileshare-service/internal/database"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func HandleRegistration(w http.ResponseWriter, r *http.Request) error {
	var details database.UserRegistration
	if err := json.NewDecoder(r.Body).Decode(&details); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	fmt.Println("Username: ", details.Username)
	fmt.Println("Password: ", details.Password)
	fmt.Println("Confirm Password: ", details.ConfirmPassword)

	return nil
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {

}

func HandleLogout(w http.ResponseWriter, r *http.Request) {

}

func HandleRetrieveBasicInfo(w http.ResponseWriter, r *http.Request) {

}
