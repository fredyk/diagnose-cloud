package entities

import (
	_ "embed"
	"time"

	"github.com/fredyk/westack-go/v2/model"
)

type Account struct {
	Id       string    `json:"id,omitempty"`
	Created  time.Time `json:"created,omitempty"`
	Modified time.Time `json:"modified,omitempty"`
	Password string    `json:"password,omitempty"`
	Email    string    `json:"email,omitempty"`
}

func NewAccount() model.Controller {
	return &Account{}
}
