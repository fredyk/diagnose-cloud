package tests

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/fredyk/diagnose-cloud/server"
	"github.com/fredyk/diagnose-cloud/server/adapters/westack/utils"
	"github.com/fredyk/westack-go/client/v2/wstfuncs"
)

var adminToken string
var regularUserToken string
var internalUserToken string
var consumer1Token string

var patientId string
var diagnoseId string

func TestMain(m *testing.M) {
	// Start the server
	go server.InitServer()
	time.Sleep(3 * time.Second)

	registeredModels, err := utils.FindRegisteredModels(server.App)
	if err != nil {
		panic(err)
	}

	wstfuncs.SetBaseUrl("http://localhost:8081/api/v1")

	// Login as admin
	adminToken = loginAsAdmin()
	fmt.Printf("Admin token: %s\n", adminToken)

	// Register a regular user
	regularUserToken = registerRegularUserAndLogin()
	fmt.Printf("Regular user token: %s\n", regularUserToken)

	// Login as internal user
	internalUserToken = createInternalUserAndLogin()
	fmt.Printf("Internal user token: %s\n", internalUserToken)

	// Login as consumer1
	consumer1Token = createConsumer1AndLogin()
	fmt.Printf("Consumer1 token: %s\n", consumer1Token)

	// Create a diagnose with internal user
	createDiagnoseWithInternalUser()

	result := m.Run()

	deleteAllInstances(registeredModels)

	os.Exit(result)
}
