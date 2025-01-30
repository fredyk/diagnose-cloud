package entities

import (
	_ "embed"

	"github.com/fredyk/diagnose-cloud/server/core/domain/entities"
	"github.com/fredyk/westack-go/v2/model"
)

type Patient entities.Patient

func NewPatient() model.Controller {
	return &Patient{}
}
