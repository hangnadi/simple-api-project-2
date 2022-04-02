package sqlt_test

import (
	"testing"

	. "github.com/kokka-team/nakama-investor-be/lib/storage/database/sqlt"
	"github.com/stretchr/testify/assert"
)

func TestNewFailed(t *testing.T) {

	db := New(Config{
		Driver: "mysql",
		Master: "master",
		Slave:  []string{"slave"},
	})

	assert.Nil(t, db)
}
