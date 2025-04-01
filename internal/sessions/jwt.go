package sessions

import (
	"financial/internal/models"

	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Token string

type AuthenticationRepository interface {
	CreateToken(user models.User) (Token, error)
	// check if token is valid and returns users id
	IsAuthenticated(token Token) (uint, bool)
}

type authenticationRepository struct {
	secretKey []byte
}

func NewAuthenticationRepository(secretKey string) AuthenticationRepository {
	return &authenticationRepository{
		secretKey: []byte(secretKey),
	}
}

func (r *authenticationRepository) CreateToken(user models.User) (Token, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"iss": "financial-app",
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(r.secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return Token(signedToken), nil
}

func (r *authenticationRepository) IsAuthenticated(token Token) (uint, bool) {
	parsedToken, err := jwt.Parse(string(token), func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return r.secretKey, nil
	})

	if err != nil || !parsedToken.Valid {
		return 0, false
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return 0, false
			}
		} else {
			return 0, false
		}

		if sub, ok := claims["sub"].(float64); ok {
			return uint(sub), true
		}
	}

	return 0, false
}
