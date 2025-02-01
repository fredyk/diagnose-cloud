package mappers

import (
	"time"

	externalentities "github.com/fredyk/diagnose-cloud/server/adapters/external-consumer-1/entities"
	"github.com/fredyk/diagnose-cloud/server/core/domain/entities"
	"github.com/fredyk/westack-go/v2/model"
)

func MapExternalDiagnoseToDiagnose(diagnose externalentities.ExternalDiagnoseDto) (entities.Diagnose, error) {
	return entities.Diagnose{
		Id:   diagnose.Id,
		Date: diagnose.Date,
		// PatientID: 			// TODO: Fetch patient ID
		// TODO: Fetch diagnose
		// Diagnose: 			item.Diagnose,
		Prescription: diagnose.Prescription,
	}, nil
}

func MapModelInstanceToDiagnoseDTO(diagnose model.Instance) (dto externalentities.ExternalDiagnoseDto, err error) {
	dateAsString := diagnose.GetString("date")
	var date time.Time
	if dateAsString != "" {
		date, err = time.Parse(time.RFC3339, dateAsString)
		if err != nil {
			return dto, err
		}
	}
	return externalentities.ExternalDiagnoseDto{
		Id:           diagnose.GetString("id"),
		PatientName:  diagnose.GetOne("patient").GetString("name"),
		DiagnoseId:   diagnose.GetString("diagnose_id"),
		Date:         date,
		Prescription: diagnose.GetString("prescription"),
	}, nil
}
