package jwt

import (
	"time"

	jwtToken "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWT struct{}

type jwtCustomClaims struct {
	UserId string
	Role   string
	jwt.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{}
}

func (j *JWT) GenerateToken(userId, role string) (tokenString string, err error) {

	claims := &jwtCustomClaims{
		UserId: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 8)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte("secret"))

	return
}

func (j *JWT) GetUserIdFromToken(ctx echo.Context) (userId, role string, err error) {

	token := ctx.Get("user").(*jwtToken.Token)
	claims := token.Claims.(jwtToken.MapClaims)
	userId = claims["UserId"].(string)
	role = claims["Role"].(string)

	return
}
