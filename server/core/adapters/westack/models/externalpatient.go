package models

import (
	_ "embed"

	"github.com/fredyk/diagnose-cloud/server/core/domain/entities"
	"github.com/fredyk/westack-go/v2/model"
)

type ExternalPatient entities.ExternalPatient

func NewExternalPatient() model.Controller {
	return &ExternalPatient{}
}
