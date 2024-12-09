package util

import "github.com/golang-jwt/jwt/v5"

type Hash interface {
	Hash(val string) (string, error)
	CheckHash(hash, val string) (bool, error)
}

type JWT interface {
	CreateToken(data map[string]any) (string, error)
	VerifyToken(token string) (*JWTClaims, error)
}

type JWTClaims struct {
	jwt.RegisteredClaims
	Data map[string]any `json:"data"`
}
