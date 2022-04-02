package internal

import (
	"github.com/kokka-team/nakama-investor-be/internal/usecase"
	nsqio "github.com/nsqio/go-nsq"
)

// Config big struct
type Config struct {

	// Main
	Server      ServerConfig
	API         APIConfig
	Environment EnvironmentConfig

	// DB
	Database map[string]*DatabaseConfig
	Elastic  map[string]*ElasticConfig
	Redis    map[string]*RedisConfig

	// Vendor
	NSQd        NSQdConfig
	NSQLookupds NSQLookupdsConfig
}

type ConsumerParam struct {
	Count       int
	Config      *nsqio.Config
	Topic       string
	Channel     string
	LogPrefix   string
	NSQLookupds []string
	Handler     nsqio.HandlerFunc
	Name        string
}

type Usecase struct {
	System usecase.System
}
