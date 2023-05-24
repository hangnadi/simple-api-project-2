package gonsq

import (
	"github.com/hangnadi/simple-api-project-2/lib/nsq"
	nsqlib "github.com/nsqio/go-nsq"
)

type mockConsumer struct {
}

// NewMockConsumer is mock for consumer
func NewMockConsumer() nsq.Consumer {
	return &mockConsumer{}
}

func (m *mockConsumer) GetConsumer(topic, channel string, config *nsqlib.Config) (err error) {
	return nil
}

func (m *mockConsumer) ConnectNSQLookupD() error {
	return nil
}
func (m *mockConsumer) ConnectNSQD() error {
	return nil
}

func (m *mockConsumer) AddHandler(handler nsqlib.Handler, concurrency int) {
}
