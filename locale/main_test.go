package locale

import (
	"os"
	"testing"

	"github.com/hangnadi/simple-api-project-2/lib/env"
)

func TestMain(m *testing.M) {

	Init(env.Development)
	os.Exit(m.Run())
}
