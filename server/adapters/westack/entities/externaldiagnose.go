package entities

import (
	_ "embed"

	"github.com/fredyk/diagnose-cloud/server/adapters/external-consumer-1/entities"
	"github.com/fredyk/westack-go/v2/model"
)

type ExternalDiagnose entities.ExternalDiagnoseDto

func NewExternalDiagnose() model.Controller {
	return &ExternalDiagnose{}
}
