package gonsq_test

import (
	"os"
	"testing"

	. "github.com/hangnadi/simple-api-project-2/lib/nsq/gonsq"
)

func TestMain(m *testing.M) {

	TFuncPatch()

	os.Exit(m.Run())
}
