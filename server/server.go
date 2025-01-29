package server

import (
	"github.com/fredyk/diagnose-cloud/common/models"
	"github.com/fredyk/diagnose-cloud/server/core/infra"
	"github.com/fredyk/westack-go/v2/westack"
)

func InitServer() error {

	app := westack.New(westack.Options{})

	app.Boot(
		westack.BootOptions{RegisterControllers: models.RegisterControllers},
		infra.SetupExternalPatients,
	)

	return app.Start()
}
