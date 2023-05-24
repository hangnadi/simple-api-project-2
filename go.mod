module github.com/hangnadi/simple-api-project-2

go 1.18

require (
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/gops v0.3.5
	github.com/h2non/gock v1.0.14
	github.com/jmoiron/sqlx v1.2.0
	github.com/julienschmidt/httprouter v1.2.0
	github.com/lib/pq v1.0.0
	github.com/nicksnyder/go-i18n v1.10.0
	github.com/nsqio/go-nsq v1.0.7
	github.com/nytimes/gziphandler v1.0.1
	github.com/opentracing/opentracing-go v1.0.2
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/robfig/cron v1.2.0
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.2.2
	gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0
	gopkg.in/gcfg.v1 v1.2.3
	gopkg.in/olivere/elastic.v5 v5.0.76
)

require (
	github.com/BurntSushi/toml v1.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db // indirect
	github.com/h2non/parth v0.0.0-20190131123155-b4df798d6542 // indirect
	github.com/kardianos/osext v0.0.0-20170510131534-ae77be60afb1 // indirect
	github.com/mailru/easyjson v0.0.0-20180823135443-60711f1a8329 // indirect
	github.com/pelletier/go-toml v1.2.0 // indirect
	github.com/pkg/errors v0.8.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.0.0-20181201002055-351d144fa1fc // indirect
	google.golang.org/appengine v1.3.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)

replace github.com/h2non/gock => gopkg.in/h2non/gock.v1 v1.0.14
