package models

import (
	_ "embed"
	"github.com/fredyk/westack-go/v2/model"
	"time"
)

type App struct {
	Id       string    `json:"id,omitempty"`
	Created  time.Time `json:"created,omitempty"`
	Modified time.Time `json:"modified,omitempty"`
	Name     string    `json:"name,omitempty"`
}

func NewApp() model.Controller {
	return &App{}
}
