package services

import "github.com/fredyk/diagnose-cloud/server/adapters/external-consumer-1/entities"

type GetDiagnosesRequest struct {
	PatientName string `json:"patient_name"`
	DOB         string `json:"dob"`
}

type ExternalDiagnoseService struct {
}

func (s *ExternalDiagnoseService) GetDiagnoses(req GetDiagnosesRequest) ([]entities.ExternalDiagnose, error) {
	return []entities.ExternalDiagnose{}, nil
}

type ExternalDiagnoseServiceOptions struct {
}

func NewExternalDiagnoseService(options ExternalDiagnoseServiceOptions) *ExternalDiagnoseService {
	return &ExternalDiagnoseService{}
}
