package login

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"github.com/monime-lab/gwater/cryptos"
)

const (
	tokenSignatureECDSAPublicKey = `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE8nRYOpaQAI2RoEUsOabcHGHO2ZHK
6xRHYp01HAfNJ0Ljllx+qDyU8NGY/kDOwGP6G+tgC+0/wTybBO1gt0K31w==
-----END PUBLIC KEY-----`
)

type tokenVerifier struct {
}

func (v tokenVerifier) verify(token SecuredToken) error {
	signatureBytes, err := base64.StdEncoding.DecodeString(token.Signature)
	if err != nil {
		return err
	}
	publicKey, err := v.loadECDSAPublicKey()
	if err != nil {
		return err
	}
	hash := sha256.New()
	hash.Write(token.bytes)
	digest := hash.Sum(nil)
	if ecdsa.VerifyASN1(publicKey, digest, signatureBytes) {
		return nil
	}
	return errors.New("authentication token cannot be verified")
}

func (v tokenVerifier) loadECDSAPublicKey() (*ecdsa.PublicKey, error) {
	key, err := cryptos.ParsePublicKey([]byte(tokenSignatureECDSAPublicKey))
	if err != nil {
		return nil, err
	}
	return key.(*ecdsa.PublicKey), nil
}
