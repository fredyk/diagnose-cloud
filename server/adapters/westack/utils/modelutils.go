package utils

import (
	"fmt"

	wst "github.com/fredyk/westack-go/v2/common"
	"github.com/fredyk/westack-go/v2/model"
	"github.com/fredyk/westack-go/v2/westack"
)

type RegisteredModelsResult struct {
	AccountModel  *model.StatefulModel
	PatientModel  *model.StatefulModel
	DiagnoseModel *model.StatefulModel

	ExternalPatient *model.StatefulModel
}

func FindRegisteredModels(wstApp *westack.WeStack) (result RegisteredModelsResult, err error) {

	result.AccountModel, err = wstApp.FindModel("Account")
	if err != nil {
		return result, fmt.Errorf("failed to finding model: %v", err)
	}

	result.PatientModel, err = wstApp.FindModel("Patient")
	if err != nil {
		return result, fmt.Errorf("failed to finding model: %v", err)
	}

	result.DiagnoseModel, err = wstApp.FindModel("Diagnose")
	if err != nil {
		return result, fmt.Errorf("failed to finding model: %v", err)
	}

	result.ExternalPatient, err = wstApp.FindModel("ExternalPatient")
	if err != nil {
		return result, fmt.Errorf("failed to finding model: %v", err)
	}

	return result, nil
}

func DisableAccountDefaultOperations(AccountModel *model.StatefulModel) {
	AccountModel.DisableRemoteOperation(wst.OperationNameFindMany)
	AccountModel.DisableRemoteOperation(wst.OperationNameCreate)
	AccountModel.DisableRemoteOperation(wst.OperationNameCount)
	AccountModel.DisableRemoteOperation(wst.OperationNameFindSelf)
	AccountModel.DisableRemoteOperation(wst.OperationNameRefreshToken)
	AccountModel.DisableRemoteOperation(wst.OperationName("resetPassword"))
	AccountModel.DisableRemoteOperation(wst.OperationName("sendVerificationEmail"))
	AccountModel.DisableRemoteOperation(wst.OperationName("performEmailVerification"))
	AccountModel.DisableRemoteOperation(wst.OperationNameDeleteById)
	AccountModel.DisableRemoteOperation(wst.OperationNameFindById)
	AccountModel.DisableRemoteOperation(wst.OperationNameUpdateAttributes)
	AccountModel.DisableRemoteOperation(wst.OperationNameEnableMfa)
}

func DisablePatientDefaultOperations(PatientModel *model.StatefulModel) {
	PatientModel.DisableRemoteOperation(wst.OperationNameCount)
}

func DisableDiagnoseDefaultOperations(DiagnoseModel *model.StatefulModel) {
	DiagnoseModel.DisableRemoteOperation(wst.OperationNameCount)
}
