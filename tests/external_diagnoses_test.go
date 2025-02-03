package tests

import (
	"fmt"
	"testing"

	"github.com/fredyk/diagnose-cloud/server/adapters/westack/entities"
	"github.com/fredyk/diagnose-cloud/tests/expected"
	"github.com/fredyk/westack-go/client/v2/wstfuncs"
	wst "github.com/fredyk/westack-go/v2/common"
	"github.com/stretchr/testify/assert"
)

func TestGetDiagnosesExternal(t *testing.T) {

	t.Parallel()

	diagnoses, err := wstfuncs.InvokeApiTyped[[]entities.ExternalDiagnose](
		"GET",
		fmt.Sprintf("/external-diagnoses?patient_name=%s", expected.DiagnosePatientName),
		nil, // No request body
		wst.M{
			"Authorization": "Bearer " + consumer1Token,
		},
	)
	assert.NoError(t, err)
	assert.NotNil(t, diagnoses)
	assert.Len(t, diagnoses, 1)
	assert.Equal(t, expected.DiagnosePatientName, diagnoses[0].PatientName)

}

func TestGetDiagnosesExternalWithoutCredentials(t *testing.T) {

	t.Parallel()

	resp, err := wstfuncs.InvokeApiFullResponse("GET", "/external-diagnoses", nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode)

}

func TestGetDiagnosesExternalWithInvalidToken(t *testing.T) {

	t.Parallel()

	resp, err := wstfuncs.InvokeApiFullResponse("GET", "/external-diagnoses", nil, wst.M{
		"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50SWQiOiI2NzlkZTI2MDliYjFlZTgyZDBkNjY5M2UiLCJjcmVhdGVkIjoxNzM4NDAwMzUzMDcwLCJyb2xlcyI6WyJVU0VSIiwiY29uc3VtZXIxIl0sInR0bCI6MTIwOTYwMDAwMH2.aP2VLaqcP8pSXs2AozHgHKpcN-cAgxSMSalWq5wqcyk",
	})
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode)

}

func TestPostDiagnosesExternal(t *testing.T) {

	t.Parallel()

	createdDiagnose, err := wstfuncs.InvokeApiTyped[entities.ExternalDiagnose](
		"POST",
		"/external-diagnoses",
		wst.M{
			"patient_name": expected.CreatedDiagnosePatientName,
			"diagnose":     expected.CreatedDiagnoseExpectedDiagnose,
		},
		wst.M{
			"Authorization": "Bearer " + consumer1Token,
			"Content-Type":  "application/json",
		},
	)
	assert.NoError(t, err)
	assert.NotNil(t, createdDiagnose)
	assert.Equal(t, expected.CreatedDiagnoseExpectedDiagnose, createdDiagnose.Diagnose)

}

func TestPostDiagnosesExternalWithoutCredentials(t *testing.T) {

	t.Parallel()

	resp, err := wstfuncs.InvokeApiFullResponse("POST", "/external-diagnoses", wst.M{
		"patient_name": expected.CreatedDiagnosePatientName,
		"diagnose":     expected.CreatedDiagnoseExpectedDiagnose,
	}, wst.M{"Content-Type": "application/json"})
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode)

}

func TestPostDiagnosesExternalWithInvalidToken(t *testing.T) {

	t.Parallel()

	resp, err := wstfuncs.InvokeApiFullResponse("POST", "/external-diagnoses", wst.M{
		"patient_name": expected.CreatedDiagnosePatientName,
		"diagnose":     expected.CreatedDiagnoseExpectedDiagnose,
	}, wst.M{
		"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50SWQiOiI2NzlkZTI2MDliYjFlZTgyZDBkNjY5M2UiLCJjcmVhdGVkIjoxNzM4NDAwMzUzMDcwLCJyb2xlcyI6WyJVU0VSIiwiY29uc3VtZXIxIl0sInR0bCI6MTIwOTYwMDAwMH2.aP2VLaqcP8pSXs2AozHgHKpcN-cAgxSMSalWq5wqcyk",
		"Content-Type":  "application/json",
	})
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode)

}
