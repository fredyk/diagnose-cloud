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

func TestGetDiagnoses(t *testing.T) {

	t.Parallel()

	diagnoses, err := wstfuncs.InvokeApiTyped[[]entities.Diagnose](
		"GET",
		fmt.Sprintf(
			"/diagnoses?filter={\"where\":{\"patient.name\":\"%s\"},\"include\":[{\"relation\":\"patient\"}]}",
			expected.DiagnosePatientName,
		),
		nil,
		wst.M{
			"Authorization": "Bearer " + internalUserToken,
		},
	)
	assert.NoError(t, err)
	assert.NotNil(t, diagnoses)

	assert.Len(t, diagnoses, 1)
	assert.Equal(t, expected.DiagnoseExpectedDiagnose, diagnoses[0].Diagnose)

}

func TestGetDiagnosesWithoutCredentials(t *testing.T) {

	t.Parallel()

	resp, err := wstfuncs.InvokeApiFullResponse("GET", "/diagnoses", nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode)

}

func TestGetDiagnosesWithInvalidToken(t *testing.T) {

	t.Parallel()

	resp, err := wstfuncs.InvokeApiFullResponse("GET", "/diagnoses", nil, wst.M{
		"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50SWQiOiI2NzlkZTI2MDliYjFlZTgyZDBkNjY5M2UiLCJjcmVhdGVkIjoxNzM4NDAwMzUzMDcwLCJyb2xlcyI6WyJVU0VSIiwiY29uc3VtZXIxIl0sInR0bCI6MTIwOTYwMDAwMH2.aP2VLaqcP8pSXs2AozHgHKpcN-cAgxSMSalWq5wqcyk",
	})
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode)

}

func TestPostDiagnosesWithoutCredentials(t *testing.T) {

	t.Parallel()

	resp, err := wstfuncs.InvokeApiFullResponse("POST", "/diagnoses", wst.M{
		"diagnose": expected.DiagnoseExpectedDiagnose,
	}, nil)
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode)

}

func TestPostDiagnosesWithInvalidToken(t *testing.T) {

	t.Parallel()

	resp, err := wstfuncs.InvokeApiFullResponse("POST", "/diagnoses", wst.M{
		"diagnose": expected.DiagnoseExpectedDiagnose,
	}, wst.M{
		"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50SWQiOiI2NzlkZTI2MDliYjFlZTgyZDBkNjY5M2UiLCJjcmVhdGVkIjoxNzM4NDAwMzUzMDcwLCJyb2xlcyI6WyJVU0VSIiwiY29uc3VtZXIxIl0sInR0bCI6MTIwOTYwMDAwMH2.aP2VLaqcP8pSXs2AozHgHKpcN-cAgxSMSalWq5wqcyk",
	})
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode)

}
