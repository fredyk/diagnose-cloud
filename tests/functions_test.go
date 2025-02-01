package tests

import (
	"fmt"
	"math/rand/v2"
	"os"
	"time"

	"github.com/fredyk/diagnose-cloud/server/adapters/westack/utils"
	"github.com/fredyk/diagnose-cloud/server/core/domain/entities"
	"github.com/fredyk/diagnose-cloud/tests/expected"
	"github.com/fredyk/westack-go/client/v2/wstfuncs"
	wst "github.com/fredyk/westack-go/v2/common"
	"github.com/fredyk/westack-go/v2/model"
)

func loginWithPasswordCredentials(credentials wst.M) (wst.LoginResult, error) {
	return wstfuncs.InvokeApiTyped[wst.LoginResult]("POST", "/accounts/login", credentials, wst.M{
		"Content-Type": "application/json",
	})
}

func loginAsAdmin() string {
	pwd := os.Getenv("WST_ADMIN_PWD")
	fmt.Printf("Admin password: %s\n", pwd)
	adminLoginResult, err := loginWithPasswordCredentials(wst.M{
		"username": os.Getenv("WST_ADMIN_USERNAME"),
		"password": pwd,
	})
	if err != nil {
		panic(err)
	}
	return adminLoginResult.Id
}

func registerRegularUserAndLogin() string {
	randomSuffix := int(1e9 + float64(rand.IntN(1e9)))
	username := fmt.Sprintf("user_%d", randomSuffix)
	// Register regular user
	_, err := wstfuncs.InvokeApiJsonM("POST", "/accounts", wst.M{
		"username": username,
		"password": "Abcd1234.",
	}, wst.M{
		"Content-Type": "application/json",
	})
	if err != nil {
		panic(err)
	}

	// Login as regular user
	regularUserLoginResult, err := loginWithPasswordCredentials(wst.M{
		"username": username,
		"password": "Abcd1234.",
	})
	if err != nil {
		panic(err)
	}

	return regularUserLoginResult.Id
}

func createInternalUserAndLogin() string {
	randomSuffix := int(1e9 + float64(rand.IntN(1e9)))
	username := fmt.Sprintf("internal_%d", randomSuffix)
	// Register internal user
	internalUserRegisterResult, err := wstfuncs.InvokeApiJsonM("POST", "/accounts", wst.M{
		"username": username,
		"password": "Abcd1234.",
	}, wst.M{
		"Content-Type": "application/json",
	})
	if err != nil {
		panic(err)
	}

	// Assign role to internal user
	_, err = wstfuncs.InvokeApiJsonM("PUT", "/accounts/"+internalUserRegisterResult.GetString("id")+"/roles", wst.M{
		"roles": []string{"internal"},
	}, wst.M{
		"Authorization": "Bearer " + adminToken,
		"Content-Type":  "application/json",
	})
	if err != nil {
		panic(err)
	}

	// Login as internal user
	internalUserLoginResult, err := loginWithPasswordCredentials(wst.M{
		"username": username,
		"password": "Abcd1234.",
	})
	if err != nil {
		panic(err)
	}

	return internalUserLoginResult.Id
}

func createConsumer1AndLogin() string {
	randomSuffix := int(1e9 + float64(rand.IntN(1e9)))
	username := fmt.Sprintf("consumer1_%d", randomSuffix)
	// Register consumer1
	consumer1RegisterResult, err := wstfuncs.InvokeApiJsonM("POST", "/accounts", wst.M{
		"username": username,
		"password": "Abcd1234.",
	}, wst.M{
		"Content-Type": "application/json",
	})
	if err != nil {
		panic(err)
	}

	// Assign role to consumer1
	_, err = wstfuncs.InvokeApiJsonM("PUT", "/accounts/"+consumer1RegisterResult.GetString("id")+"/roles", wst.M{
		"roles": []string{"consumer1"},
	}, wst.M{
		"Authorization": "Bearer " + adminToken,
		"Content-Type":  "application/json",
	})
	if err != nil {
		panic(err)
	}

	// Login as consumer1
	consumer1LoginResult, err := loginWithPasswordCredentials(wst.M{
		"username": username,
		"password": "Abcd1234.",
	})
	if err != nil {
		panic(err)
	}

	return consumer1LoginResult.Id
}

func createDiagnoseWithInternalUser() {

	// First create the patient

	patient, err := wstfuncs.InvokeApiTyped[entities.Patient]("POST", "/patients", wst.M{
		"name":           expected.DiagnosePatientName,
		"documentNumber": "123456789",
		"email":          "john.doe@example.com",
		"dob":            time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
	}, wst.M{
		"Authorization": "Bearer " + internalUserToken,
		"Content-Type":  "application/json",
	})
	if err != nil {
		panic(err)
	}

	patientId = patient.Id

	diagnose, err := wstfuncs.InvokeApiTyped[entities.Diagnose]("POST", "/diagnoses", wst.M{
		"patientId": patientId,
		"diagnose":  expected.DiagnoseExpectedDiagnose,
	}, wst.M{
		"Authorization": "Bearer " + internalUserToken,
		"Content-Type":  "application/json",
	})

	if err != nil {
		panic(err)
	}

	diagnoseId = diagnose.Id

}

func deleteAllInstances(allModels utils.RegisteredModelsResult) {
	systemContext := &model.EventContext{Bearer: &model.BearerToken{Account: &model.BearerAccount{System: true}}}
	var err error
	var deleteResult wst.DeleteResult
	deleteResult, err = allModels.PatientModel.DeleteMany(&wst.Where{"_id": wst.M{"$ne": nil}}, systemContext)
	if err != nil {
		fmt.Printf("Failed to delete patients: %v\n", err)
	}
	fmt.Printf("Deleted %d patients\n", deleteResult.DeletedCount)
	deleteResult, err = allModels.DiagnoseModel.DeleteMany(&wst.Where{"_id": wst.M{"$ne": nil}}, systemContext)
	if err != nil {
		fmt.Printf("Failed to delete diagnoses: %v\n", err)
	}
	fmt.Printf("Deleted %d diagnoses\n", deleteResult.DeletedCount)
}
