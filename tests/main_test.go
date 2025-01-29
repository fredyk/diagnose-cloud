package tests

import (
	"os"
	"testing"
	"time"

	"github.com/fredyk/diagnose-cloud/server"
	"github.com/fredyk/westack-go/client/v2/wstfuncs"
)

func TestMain(m *testing.M) {
	// Start the server
	go server.InitServer()
	time.Sleep(5 * time.Second)

	wstfuncs.SetBaseUrl("http://localhost:8080/api/v1")

	os.Exit(m.Run())
}
