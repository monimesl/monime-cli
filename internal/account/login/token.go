package login

import (
	"encoding/json"
	"time"
)

type SecuredToken struct {
	Secret    string
	Signature string
	bytes     []byte
}

func (st SecuredToken) extractToken() (Token, error) {
	token := Token{}
	if err := json.Unmarshal(st.bytes, &token); err != nil {
		return token, err
	}
	return token, nil
}

type Token struct {
	Id      string `json:"id"`
	State   string `json:"state"`
	Account struct {
		Id    string `json:"id"`
		Alias string `json:"alias"`
	} `json:"account"`
	CreateTime time.Time `json:"createTime"`
}
