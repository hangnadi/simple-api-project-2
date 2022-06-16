# Template Project for creating API

## How To Develop

### Onboarding
- Install Golang `1.11.2` 
- Copy `files/etc/template/development` to `development` to main folder
- Run `go mod` to get all dependencies

[ARCHIVE]
- Download NSQ (https://nsq.io/deployment/installing.html), then copy `startNSQ.sh` and `stopNSQ.sh` located inside `docker/script` folder to your NSQ download folder
- To start NSQ, run script `./startNSQ.sh`
- NSQ will running in background until your restart your machine. To kill NSQ without restarting, run script `./stopNSQ.sh`

### Compiling
- For UNIX, build using `make build`
- For Windows, build using `go build ./cmd/api-temp` - `go build ./cmd/cron`

### Running
- Run docker, refer to [this document](../docker/README.md)

or

- Run in local device, `./api-temp`

### Learning
- To learn how this service works, go to `cmd/api/api.main.go` for API and `cmd/cron/cron.main.go` for Cron
- To test db connection `https://stackoverflow.com/questions/26911508/postgres-testing-database-connection-in-bash\`

### Testing
- curl -X GET http://localhost/{version}/testing/getdata/
