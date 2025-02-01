package server

import (
	"github.com/fredyk/diagnose-cloud/server/adapters/westack/entities"
	"github.com/fredyk/diagnose-cloud/server/infra"
	"github.com/fredyk/westack-go/v2/westack"
)

var App *westack.WeStack

func InitServer() error {

	App = westack.New(westack.Options{})

	App.Boot(
		westack.BootOptions{RegisterControllers: entities.RegisterControllers},
		infra.SetupCloud,
	)

	return App.Start()
}
