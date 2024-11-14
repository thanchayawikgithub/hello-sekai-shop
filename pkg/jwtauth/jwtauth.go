package jwtauth

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/metadata"
)

type (
	AuthFactory interface {
		SignToken() string
	}

	Claims struct {
		PlayerID string `json:"player_id"`
		RoleCode int    `json:"role_code"`
	}

	AuthMapClaims struct {
		*Claims
		jwt.RegisteredClaims
	}

	authConcrete struct {
		Secret []byte
		Claims *AuthMapClaims `json:"claims"`
	}

	accessToken struct {
		*authConcrete
	}

	refreshToken struct {
		*authConcrete
	}

	apiKey struct {
		*authConcrete
	}
)

var (
	apiKeyInstance string
	once           sync.Once
)

func (a *authConcrete) SignToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.Claims)

	tokenString, _ := token.SignedString(a.Secret)

	return tokenString
}

// t is in seconds
func jwtTimeDurationCal(t int64) *jwt.NumericDate {
	return jwt.NewNumericDate(now().Add(time.Duration(t) * time.Second))
}

func jwtTimeRepeatAdapter(t int64) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Unix(t, 0))
}

func NewAccessToken(secret string, expriesAt int64, claims *Claims) AuthFactory {
	return &accessToken{
		authConcrete: &authConcrete{
			Secret: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "hello-sekai.com",
					Subject:   "access-token",
					Audience:  []string{"hello-sekai.com"},
					ExpiresAt: jwtTimeDurationCal(expriesAt),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}
}

func NewRefreshToken(secret string, expriesAt int64, claims *Claims) AuthFactory {
	return &refreshToken{
		authConcrete: &authConcrete{
			Secret: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "hello-sekai.com",
					Subject:   "refresh-token",
					Audience:  []string{"hello-sekai.com"},
					ExpiresAt: jwtTimeDurationCal(expriesAt),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}
}

func ReloadToken(secret string, expriesAt int64, claims *Claims) string {
	token := &refreshToken{
		authConcrete: &authConcrete{
			Secret: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "hello-sekai.com",
					Subject:   "refresh-token",
					Audience:  []string{"hello-sekai.com"},
					ExpiresAt: jwtTimeRepeatAdapter(expriesAt),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}

	return token.SignToken()
}

func NewApiKey(secret string) AuthFactory {
	return &apiKey{
		authConcrete: &authConcrete{
			Secret: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: &Claims{},
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "hello-sekai.com",
					Subject:   "api-key",
					Audience:  []string{"hello-sekai.com"},
					ExpiresAt: jwtTimeDurationCal(31560000),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}
}

func ParseToken(secret string, tokenString string) (*AuthMapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthMapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error: unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errors.New("error: invalid token format")
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("error: token expired")
		}

		return nil, errors.New("error: invalid token")
	}

	if claims, ok := token.Claims.(*AuthMapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("error: invalid token claims")
}

func SetApiKey(secret string) {
	once.Do(func() {
		apiKeyInstance = NewApiKey(secret).SignToken()
	})
}

func SetApiKeyInContext(ctx *context.Context) {
	*ctx = metadata.NewOutgoingContext(*ctx, metadata.Pairs("auth", apiKeyInstance))
}

func now() time.Time {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return time.Now().In(loc)
}
