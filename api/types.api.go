/*Package api types for api
Please consult with PE Before updating this file

*/
package api

import (
	"context"
	"net/http"
	"time"

	"github.com/hangnadi/simple-api-project-2/lib/router"
	"github.com/julienschmidt/httprouter"
)

type (

	//router for API
	apiRouter struct {
		rt                *httprouter.Router
		config            RouterConfig
		middlewares       []router.Middleware
		inlineMiddlewares []router.Middleware
		panicsHandler     func(context.Context, *http.Request, *router.HandlerParam, func(context.Context, *router.HandlerParam) router.HandlerResult) router.HandlerResult
	}

	// RouterConfig struct for API
	RouterConfig struct {
		Timeout time.Duration // Set timeout duration
		// Connectivity
	}
)
