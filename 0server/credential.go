package server

import (
	"encoding/json"
	"io"
	"os"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CredArr []Credentials

func (p *CredArr) FromJSON(r io.Reader) error {
	en := json.NewDecoder(r)
	return en.Decode(p)
}

func GetCred() *CredArr {

	f, _ := os.Open("credential.json")
	var creds CredArr
	creds.FromJSON(f)
	return &creds
}
