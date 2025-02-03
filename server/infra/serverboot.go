package infra

import (
	"fmt"
	"log"

	"github.com/fredyk/diagnose-cloud/server/adapters/external-consumer-1/repositories"
	"github.com/fredyk/diagnose-cloud/server/adapters/external-consumer-1/services"
	"github.com/fredyk/diagnose-cloud/server/adapters/westack/utils"
	"github.com/fredyk/diagnose-cloud/server/core/app"
	"github.com/fredyk/westack-go/v2/model"
	"github.com/fredyk/westack-go/v2/westack"
)

func SetupCloud(wstApp *westack.WeStack) {

	registeredModels, err := utils.FindRegisteredModels(wstApp)
	if err != nil {
		log.Fatalf("Failed to find registered models: %v", err)
	}

	// Setup repositories
	externalDiagnoseServiceImpl := services.NewExternalDiagnoseService(services.ExternalDiagnoseServiceOptions{
		Repository: repositories.NewDiagnosesRepository(repositories.DiagnosesRepositoryOptions{
			// Important to use the internal DiagnoseModel, because the External one is not persisted
			PersitedDiagnoseModel: registeredModels.DiagnoseModel,
		}),
	})

	// Setup our Diagnose Cloud
	// TODO: Setup some business logic in the Diagnose Cloud to use it
	cloud := app.NewDiagnoseCloud(app.Options{
		ExternalDiagnoseService: services.NewDiagnoseServiceAdapter(externalDiagnoseServiceImpl),
	})
	fmt.Printf("cloud: %v\n", cloud)

	// Westack comes with multiple built-in operations that can be disabled
	utils.DisableAccountDefaultOperations(registeredModels.AccountModel)
	utils.DisablePatientDefaultOperations(registeredModels.PatientModel)
	utils.DisableDiagnoseDefaultOperations(registeredModels.DiagnoseModel)

	// Setup our custom API endpoints
	model.BindRemoteOperationWithOptions(registeredModels.ExternalDiagnoseModel, externalDiagnoseServiceImpl.GetDiagnoses,
		model.RemoteOptions().
			WithName("Consumer1GetDiagnoses").
			WithPath("/").
			WithVerb("get"))
	model.BindRemoteOperationWithOptions(registeredModels.ExternalDiagnoseModel, externalDiagnoseServiceImpl.CreateDiagnoses,
		model.RemoteOptions().
			WithName("Consumer1CreateDiagnose").
			WithPath("/").
			WithVerb("post"))

}
