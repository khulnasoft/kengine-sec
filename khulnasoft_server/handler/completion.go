package handler

import (
	"net/http"

	"github.com/khulnasoft/kengine/khulnasoft_server/model"
	"github.com/khulnasoft/kengine/khulnasoft_server/reporters"
	"github.com/khulnasoft/kengine/khulnasoft_server/reporters/completion"
	"github.com/khulnasoft/kengine/khulnasoft_utils/log"
	httpext "github.com/go-playground/pkg/v5/net/http"
)

func (h *Handler) CompleteProcessInfo(w http.ResponseWriter, r *http.Request) {
	genericCompleteInfoHandler[model.Process](w, r, h)
}

func (h *Handler) CompleteVulnerabilityInfo(w http.ResponseWriter, r *http.Request) {
	genericCompleteInfoHandler[model.VulnerabilityRule](w, r, h)
}

func (h *Handler) CompleteHostInfo(w http.ResponseWriter, r *http.Request) {
	genericCompleteInfoHandler[model.Host](w, r, h)
}

func (h *Handler) CompleteComplianceInfo(w http.ResponseWriter, r *http.Request) {
	genericCompleteInfoHandler[model.ComplianceRule](w, r, h)
}

func (h *Handler) CompleteCloudComplianceInfo(w http.ResponseWriter, r *http.Request) {
	genericCompleteInfoHandler[model.CloudCompliance](w, r, h)
}

func (h *Handler) CompletePodInfo(w http.ResponseWriter, r *http.Request) {
	genericCompleteInfoHandler[model.Pod](w, r, h)
}

func (h *Handler) CompleteContainerInfo(w http.ResponseWriter, r *http.Request) {
	genericCompleteInfoHandler[model.Container](w, r, h)
}

func genericCompleteInfoHandler[T reporters.Cypherable](w http.ResponseWriter, r *http.Request, h *Handler) {
	defer r.Body.Close()
	var req completion.CompletionNodeFieldReq
	err := httpext.DecodeJSON(r, httpext.NoQueryParams, MaxPostRequestSize, &req)
	if err != nil {
		h.respondError(&BadDecoding{err}, w)
		return
	}

	entries, err := completion.FieldValueCompletion[T](r.Context(), req)
	if err != nil {
		log.Error().Msg(err.Error())
		h.respondError(err, w)
		return
	}

	err = httpext.JSON(w, http.StatusOK, completion.CompletionNodeFieldRes{
		PossibleValues: entries,
	})
	if err != nil {
		log.Error().Msg(err.Error())
	}
}
