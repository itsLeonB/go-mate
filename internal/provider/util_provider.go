package provider

import (
	"github.com/itsLeonB/go-mate/internal/config"
	"github.com/itsLeonB/go-mate/internal/util"
)

type Utils struct {
	Hash util.Hash
	JWT  util.JWT
}

func ProvideUtils(configs *config.Auth) *Utils {
	return &Utils{
		Hash: util.NewHashBcrypt(10),
		JWT:  util.NewJWTProviderHS256(configs),
	}
}
