package utils

import (
	"errors"
	"net/http"
	"strings"

	"github.com/JuanDiegoCastellanos/rest-ws/models"
	"github.com/JuanDiegoCastellanos/rest-ws/server"
	"github.com/golang-jwt/jwt"
)

func TokenExtractor(s server.Server, r *http.Request, w http.ResponseWriter) (*models.AppClaims, error) {
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JWTSecret), nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return nil, errors.New("unauthorized")
	}
	claims, ok := token.Claims.(*models.AppClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token is not valid")
}
