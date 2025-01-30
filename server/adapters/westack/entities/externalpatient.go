package entities

import (
	_ "embed"

	"github.com/fredyk/diagnose-cloud/server/adapters/external-consumer-1/entities"
	"github.com/fredyk/westack-go/v2/model"
)

type ExternalPatient entities.ExternalPatient

func NewExternalPatient() model.Controller {
	return &ExternalPatient{}
}
