package server

import (
	"github.com/fredyk/diagnose-cloud/server/adapters/westack/entities"
	"github.com/fredyk/diagnose-cloud/server/infra"
	"github.com/fredyk/westack-go/v2/westack"
)

func InitServer() error {

	app := westack.New(westack.Options{})

	app.Boot(
		westack.BootOptions{RegisterControllers: entities.RegisterControllers},
		infra.SetupCloud,
	)

	return app.Start()
}
