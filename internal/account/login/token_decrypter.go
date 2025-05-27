package login

import (
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/nacl/box"
)

type tokenDecrypter struct {
}

func (d tokenDecrypter) decrypt(token SecuredToken, privateKey *[32]byte) ([]byte, error) {
	secret, err := base64.StdEncoding.DecodeString(token.Secret)
	if err != nil {
		return nil, err
	}
	var senderPublicKey [32]byte
	copy(senderPublicKey[:], secret[:32])

	var nonce [24]byte
	copy(nonce[:], secret[32:64])

	cipher := secret[56:]
	tokenBytes, ok := box.Open(nil, cipher, &nonce, &senderPublicKey, privateKey)
	if !ok {
		return nil, errors.New("failed to decrypt secured token")
	}
	return tokenBytes, nil
}
