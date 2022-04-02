package locale

import (
	"os"
	"testing"

	"github.com/kokka-team/nakama-investor-be/lib/constants/env"
)

func TestMain(m *testing.M) {

	Init(env.Development)
	os.Exit(m.Run())
}
