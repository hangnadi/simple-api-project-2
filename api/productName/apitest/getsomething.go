package apitest

import (
	"context"
	"net/http"

	"github.com/hangnadi/simple-api-project-2/lib/language"
	"github.com/hangnadi/simple-api-project-2/lib/router"
	"github.com/hangnadi/simple-api-project-2/locale"
)

// Get all holidays
func (m *Module) getTestingData(ctx context.Context, param *router.HandlerParam) (apiResult router.HandlerResult) {

	// Validate language
	langExists := false
	availableLanguages := []string{"id-ID", "en-US"}
	for l := range availableLanguages {
		if availableLanguages[l] == param.Lang {
			langExists = true
			break
		}
	}

	// Language didn't exists
	if !langExists {
		apiResult.SetError(locale.Translate(language.ID, "ErrorServer"), "45", http.StatusBadRequest)
		return apiResult
	}

	// Process result
	resultData := router.ArrObject{}

	type testing struct {
		Label string
		Value string
	}

	testRes := testing{
		Label: "testing",
		Value: "testing",
	}
	test := router.Object{
		ID:         "testing",
		Attributes: testRes,
	}

	resultData = append(resultData, test)

	apiResult.JSON = &router.Data{
		Data: resultData,
	}

	return apiResult
}
