package middlewares

import (
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

func GenerateToken(user models.User) string {
	claims := &jwtCustomClaims{
		user.ID,
		user.Role.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(util.GetConfig("JWT_SECRET_KEY")))

	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, "fail to generate token")
	}

	return tokenString
}
