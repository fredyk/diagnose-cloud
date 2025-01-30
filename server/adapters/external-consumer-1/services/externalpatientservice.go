package services

import "github.com/fredyk/diagnose-cloud/server/adapters/external-consumer-1/entities"

type GetPatientsRequest struct{}
type GetDiagnosesRequest struct {
	PatientID string `json:"patient_id"`
}

type ExternalService struct {
}

func (s *ExternalService) GetPatients(req GetPatientsRequest) ([]entities.ExternalPatient, error) {
	return []entities.ExternalPatient{}, nil
}

func (s *ExternalService) GetDiagnoses(req GetDiagnosesRequest) ([]entities.ExternalDiagnose, error) {
	return []entities.ExternalDiagnose{}, nil
}

type Options struct {
}

func NewExternalService(options Options) *ExternalService {
	return &ExternalService{}
}
