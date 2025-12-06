package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(getSecret())

func getSecret() string {
	if s := os.Getenv("JWT_SECRET"); s != "" {
		return s
	}
	return "dev-secret"
}

// GenerateToken creates a signed JWT with user and org claims.
func GenerateToken(userID string, orgID string, ttl time.Duration) (string, error) {
	if ttl == 0 {
		ttl = 24 * time.Hour
	}
	claims := jwt.MapClaims{
		"sub": userID,
		"org": orgID,
		"exp": time.Now().Add(ttl).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken validates and returns claims as a map.
func ParseToken(tokenStr string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		out := map[string]interface{}{}
		for k, v := range claims {
			out[k] = v
		}
		return out, nil
	}
	return nil, errors.New("invalid claims")
}
