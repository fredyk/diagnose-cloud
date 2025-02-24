package repositories

import (
	"time"

	"github.com/fredyk/diagnose-cloud/server/adapters/external-consumer-1/entities"
	"github.com/fredyk/diagnose-cloud/server/adapters/external-consumer-1/mappers"
	"github.com/fredyk/diagnose-cloud/server/adapters/external-consumer-1/utils"
	wst "github.com/fredyk/westack-go/v2/common"
	"github.com/fredyk/westack-go/v2/model"
)

type DiagnosesRepository interface {
	GetDiagnosesByNameOrDOB(patientName string, dob time.Time) ([]entities.ExternalDiagnoseDto, error)
	CreateDiagnose(diagnose entities.ExternalDiagnoseDto) (model.Instance, error)
}

type DiagnosesRepositoryImpl struct {
	PersistedDiagnoseModel *model.StatefulModel
	systemContext          *model.EventContext
}

func (r *DiagnosesRepositoryImpl) GetDiagnosesByNameOrDOB(patientName string, dob time.Time) ([]entities.ExternalDiagnoseDto, error) {
	where := buildWhereAndOrPatientNameDob(patientName, dob)
	diagnosesCursor := r.PersistedDiagnoseModel.FindMany(&wst.Filter{Where: &where, Include: &wst.Include{{Relation: "patient"}}}, r.systemContext)

	diagnoses, err := diagnosesCursor.All()
	if err != nil {
		return nil, err
	}

	return utils.MapItemsWithError(diagnoses, mappers.MapModelInstanceToDiagnoseDTO)
}

func (r *DiagnosesRepositoryImpl) CreateDiagnose(diagnose entities.ExternalDiagnoseDto) (model.Instance, error) {
	// TODO: Create the patient first
	return r.PersistedDiagnoseModel.Create(wst.M{
		"diagnose": diagnose.Diagnose,
	}, r.systemContext)
}

type DiagnosesRepositoryOptions struct {
	PersitedDiagnoseModel *model.StatefulModel
}

func NewDiagnosesRepository(options DiagnosesRepositoryOptions) DiagnosesRepository {
	systemContext := &model.EventContext{Bearer: &model.BearerToken{Account: &model.BearerAccount{System: true}}}
	return &DiagnosesRepositoryImpl{
		systemContext:          systemContext,
		PersistedDiagnoseModel: options.PersitedDiagnoseModel,
	}
}
