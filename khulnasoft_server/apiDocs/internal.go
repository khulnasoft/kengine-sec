package apiDocs //nolint:stylecheck

import (
	"net/http"

	"github.com/khulnasoft/kengine/khulnasoft_server/model"
)

func (d *OpenAPIDocs) AddInternalAuthOperations() {
	d.AddOperation("getConsoleApiToken", http.MethodGet, "/khulnasoft/internal/console-api-token",
		"Get api-token for console agent", "Get api-token for console agent",
		http.StatusOK, []string{tagInternal}, nil, nil, new(model.APIAuthRequest))
}
