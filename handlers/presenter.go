package handlers

import (
	"encoding/json"
	"net/http"
	"resources-ms/domain/resources/exceptions"
	"resources-ms/domain/resources/gateways"
)

type Presenter struct {
	Writer http.ResponseWriter
}

func (p *Presenter) Present(body gateways.IResource) {
	p.Writer.Header().Set("Content-Type", "application/json")
	p.Writer.WriteHeader(http.StatusOK)

	json.NewEncoder(p.Writer).Encode(body)
}

func (p *Presenter) PresentError(err *exceptions.ApiError) {
	p.Writer.Header().Set("Content-Type", "application/json")
	p.Writer.WriteHeader(err.StatusCode)
	json.NewEncoder(p.Writer).Encode(err)
}
