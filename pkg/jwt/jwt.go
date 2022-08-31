package jwt

import (
	"fmt"
	"time"

	jwtLib "github.com/golang-jwt/jwt/v4"
)

type JwtHelper struct {
	Config struct {
		SecretKey       string
		TokenExpiration int
	}
}

func (jwt *JwtHelper) TokenFromUser(userId int) (string, error) {
	now := time.Now()
	expiredAt := now.Add(time.Second * time.Duration(jwt.Config.TokenExpiration))
	token := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, jwtLib.MapClaims{
		"iat":     now.Unix(),
		"exp":     expiredAt.Unix(),
		"user_id": userId,
	})

	tokenString, err := token.SignedString([]byte(jwt.Config.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jwt *JwtHelper) VerifyToken(token string) (jwtLib.MapClaims, bool) {
	jwtToken, err := jwtLib.Parse(token, func(token *jwtLib.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtLib.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwt.Config.SecretKey), nil
	})

	if err != nil {
		return nil, false
	}
	return jwtToken.Claims.(jwtLib.MapClaims), err == nil
}
