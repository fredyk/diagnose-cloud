package services

import (
	"time"

	"github.com/fredyk/diagnose-cloud/server/adapters/external-consumer-1/mappers"
	"github.com/fredyk/diagnose-cloud/server/adapters/external-consumer-1/utils"
	"github.com/fredyk/diagnose-cloud/server/core/domain/entities"
)

// DiagnoseServiceAdapter is a struct that implements the DiagnosePort interface
type DiagnoseServiceAdapter struct {
	service *ExternalDiagnoseService
}

func (s *DiagnoseServiceAdapter) GetDiagnoses(patientName string, patientDob time.Time) ([]entities.Diagnose, error) {
	req := GetDiagnosesRequest{
		PatientName: patientName,
		DOB:         patientDob,
	}
	diagnoses, err := s.service.GetDiagnoses(req)
	if err != nil {
		return nil, err
	}
	return utils.MapItems(diagnoses, mappers.MapExternalDiagnoseToDiagnose), nil
}

// NewDiagnoseServiceAdapter creates a new instance of DiagnoseServiceAdapter
func NewDiagnoseServiceAdapter(service *ExternalDiagnoseService) *DiagnoseServiceAdapter {
	return &DiagnoseServiceAdapter{
		service: service,
	}
}
