package dummy_test

import (
	"testing"

	. "github.com/hangnadi/simple-api-project-2/lib/nsq/dummy"
	"github.com/stretchr/testify/assert"
)

func TestPublishSuccess(t *testing.T) {
	module := New(false)
	assert.Nil(t, module.Publish("foo", nil))
}
func TestPublishError(t *testing.T) {
	module := New(true)
	assert.EqualError(t, module.Publish("foo", nil), "Always error NSQ")
}
