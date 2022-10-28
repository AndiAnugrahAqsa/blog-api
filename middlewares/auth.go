package middlewares

import (
	"errors"
	"mini-project/models"
	"mini-project/util"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	ID     int `json:"id"`
	RoleID int `json:"role_id"`
	jwt.StandardClaims
}

func GenerateToken(user models.User) (string, error) {
	claims := &jwtCustomClaims{
		user.ID,
		user.RoleID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(util.GetConfig("JWT_SECRET_KEY")))

	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "fail to generate token")
	}

	return tokenString, nil
}

func GetJWTSecretKey(token *jwt.Token) (interface{}, error) {

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("Can not map the token")
	}

	if claims["role_id"] != float64(1) {
		return nil, errors.New("Can not access this path")
	}

	return []byte(util.GetConfig("JWT_SECRET_KEY")), nil
}
