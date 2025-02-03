package services

import (
	"time"

	"github.com/fredyk/diagnose-cloud/server/adapters/external-consumer-1/entities"
	"github.com/fredyk/diagnose-cloud/server/adapters/external-consumer-1/mappers"
	"github.com/fredyk/diagnose-cloud/server/adapters/external-consumer-1/repositories"
)

type GetDiagnosesRequest struct {
	PatientName string    `query:"patient_name"`
	DOB         time.Time `query:"dob"`
}

func (r GetDiagnosesRequest) GetPatientName() string {
	return r.PatientName
}

func (r GetDiagnosesRequest) GetPatientDob() time.Time {
	return r.DOB
}

type ExternalDiagnoseService struct {
	repository repositories.DiagnosesRepository
}

func (s *ExternalDiagnoseService) GetDiagnoses(req GetDiagnosesRequest) ([]entities.ExternalDiagnoseDto, error) {
	return s.repository.GetDiagnosesByNameOrDOB(req.PatientName, req.DOB)
}

func (s *ExternalDiagnoseService) CreateDiagnoses(req entities.ExternalDiagnoseDto) (entities.ExternalDiagnoseDto, error) {
	created, err := s.repository.CreateDiagnose(req)
	if err != nil {
		return entities.ExternalDiagnoseDto{}, err
	}
	return mappers.MapModelInstanceToDiagnoseDTO(created)
}

type ExternalDiagnoseServiceOptions struct {
	Repository repositories.DiagnosesRepository
}

func NewExternalDiagnoseService(options ExternalDiagnoseServiceOptions) *ExternalDiagnoseService {
	return &ExternalDiagnoseService{
		repository: options.Repository,
	}
}
