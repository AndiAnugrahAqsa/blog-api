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

var whitelist []string = make([]string, 5)

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

func GetJWTSecretKeyForAdmin(token *jwt.Token) (interface{}, error) {

	claims, ok := token.Claims.(jwt.MapClaims)

	if isListed := CheckToken(token.Raw); !isListed {
		return nil, errors.New("invalid or expired jwt")
	}

	if !ok {
		return nil, errors.New("invalid or expired jwt")
	}

	if claims["role_id"] != float64(1) {
		return nil, errors.New("invalid or expired jwt")
	}

	return []byte(util.GetConfig("JWT_SECRET_KEY")), nil
}

func GetJWTSecretKeyForUser(token *jwt.Token) (interface{}, error) {
	if isListed := CheckToken(token.Raw); !isListed {
		return nil, errors.New("invalid or expired jwt")
	}

	return []byte(util.GetConfig("JWT_SECRET_KEY")), nil
}

func AddTokenInWhiteList(token string) {
	whitelist = append(whitelist, token)
}

func CheckToken(token string) bool {
	for _, tkn := range whitelist {
		if tkn == token {
			return true
		}
	}

	return false
}

func Logout(token string) bool {
	for idx, tkn := range whitelist {
		if tkn == token {
			whitelist = append(whitelist[:idx], whitelist[idx+1:]...)
		}
	}

	return true
}
