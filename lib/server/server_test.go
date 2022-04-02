package server_test

import (
	"testing"

	. "github.com/kokka-team/nakama-investor-be/lib/server"
	"github.com/stretchr/testify/assert"
)

func TestGetIPAddress(t *testing.T) {
	assert.NotNil(t, GetIPAddress())
}
