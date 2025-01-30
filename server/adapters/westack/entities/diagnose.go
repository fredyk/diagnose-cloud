package entities

import (
	_ "embed"

	"github.com/fredyk/diagnose-cloud/server/core/domain/entities"
	"github.com/fredyk/westack-go/v2/model"
)

type Diagnose entities.Diagnose

func NewDiagnose() model.Controller {
	return &Diagnose{}
}
