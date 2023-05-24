package main

import (
	"github.com/hangnadi/simple-api-project-2/cmd/internal"
	"github.com/hangnadi/simple-api-project-2/lib/panics"
	c "github.com/robfig/cron"
)

var (
	cronObj *c.Cron
)

func initCron(panicWrapper panics.Panics, ucase *internal.Usecase) {

	// Initialize cron object
	cronObj = c.New()

	//---------------------
	// START - Cron list
	//---------------------

	// -- Open files logging datadog
	// -- Every 10 seconds
	cronObj.AddFunc("@every 10s", panicWrapper.Cron(func() {
		ucase.System.LogOpenFile()
	}))

	// END - Cron list
	//---------------------

	cronObj.Start()
}
