package utils

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"

	"github.com/shivajee98/aamishrit/internal/model"
)

func FetchClerkPublicKey(jwksUrl string, kid string) (*rsa.PublicKey, error) {
	resp, err := http.Get(jwksUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jwks model.JWKS
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return nil, err
	}

	for _, key := range jwks.Keys {
		if key.Kid == kid {
			nBytes, err := base64.RawURLEncoding.DecodeString(key.N)
			if err != nil {
				return nil, err
			}

			eBytes, err := base64.RawURLEncoding.DecodeString(key.E)
			if err != nil {
				return nil, err
			}

			e := 0
			for _, b := range eBytes {
				e = e<<8 + int(b)
			}

			pubKey := &rsa.PublicKey{
				N: new(big.Int).SetBytes(nBytes),
				E: e,
			}
			return pubKey, nil
		}
	}

	return nil, errors.New("public key not found for the given kid")
}
