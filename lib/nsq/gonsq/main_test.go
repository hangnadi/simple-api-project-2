package gonsq_test

import (
	"os"
	"testing"

	. "github.com/kokka-team/nakama-investor-be/lib/nsq/gonsq"
)

func TestMain(m *testing.M) {

	TFuncPatch()

	os.Exit(m.Run())
}
