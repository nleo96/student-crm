package utils

import (
	"net/http"
	"open-crm-api/errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(user int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["usr"] = user
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	var signedToken string
	var err error

	signedToken, err = token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

	return signedToken, errors.ErrTokenSigning.Wrap(err)
}

func GenerateSessionCookie(token string) *http.Cookie {
	return &http.Cookie{
		Name:     "session_cookie",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
}
