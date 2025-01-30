package infra

import (
	"log"

	"github.com/fredyk/diagnose-cloud/server/adapters/external-consumer-1/entities"
	"github.com/fredyk/diagnose-cloud/server/adapters/external-consumer-1/services"
	"github.com/fredyk/diagnose-cloud/server/adapters/westack/utils"
	"github.com/fredyk/westack-go/v2/model"
	"github.com/fredyk/westack-go/v2/westack"
)

func SetupCloud(wstApp *westack.WeStack) {

	// cloud := app.NewDiagnoseCloud(app.Options{})
	registeredModels, err := utils.FindRegisteredModels(wstApp)
	if err != nil {
		log.Fatalf("Failed to find registered models: %v", err)
	}

	// Westack comes with multiple built-in operations that can be disabled
	utils.DisableAccountDefaultOperations(registeredModels.AccountModel)
	utils.DisablePatientDefaultOperations(registeredModels.PatientModel)
	utils.DisableDiagnoseDefaultOperations(registeredModels.DiagnoseModel)

	// Bind remote operations for External Service
	externalService := services.NewExternalService(services.Options{})

	model.BindRemoteOperationWithOptions[services.GetPatientsRequest, []entities.ExternalPatient](registeredModels.ExternalPatient, externalService.GetPatients,
		model.RemoteOptions().
			WithName("Consumer1GetPatients").
			WithPath("/consumer-1/get-patients").
			WithVerb("get"))

	model.BindRemoteOperationWithOptions(registeredModels.DiagnoseModel, externalService.GetDiagnoses,
		model.RemoteOptions().
			WithName("Consumer1GetDiagnoses").
			WithPath("/consumer-1/get-diagnoses").
			WithVerb("get"))

}
