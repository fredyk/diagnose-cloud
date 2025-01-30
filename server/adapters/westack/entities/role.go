package entities

import (
	_ "embed"
	"time"

	"github.com/fredyk/westack-go/v2/model"
)

type Role struct {
	Id       string    `json:"id,omitempty"`
	Created  time.Time `json:"created,omitempty"`
	Modified time.Time `json:"modified,omitempty"`
	Name     string    `json:"name,omitempty"`
}

func NewRole() model.Controller {
	return &Role{}
}
