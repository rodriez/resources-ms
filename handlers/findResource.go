package handlers

import (
	"net/http"
	"resources-ms/domain/resources"
	"resources-ms/repositories"

	"github.com/gorilla/mux"
)

func FindResource(res http.ResponseWriter, req *http.Request) {
	repository := repositories.ResourceRepository{}
	presenter := Presenter{Writer: res}

	useCase := resources.FindResourceUseCase{
		FindResource:   repository.FindResource,
		PresentSuccess: presenter.Present,
	}

	params := mux.Vars(req)

	if apiError := useCase.Run(params["id"]); apiError != nil {
		presenter.PresentError(apiError)

		return
	}
}
