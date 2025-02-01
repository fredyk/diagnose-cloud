package ports

import (
	"time"

	"github.com/fredyk/diagnose-cloud/server/core/domain/entities"
)

type DiagnosePort interface {
	GetDiagnoses(patientName string, patientDob time.Time) ([]entities.Diagnose, error)
}
