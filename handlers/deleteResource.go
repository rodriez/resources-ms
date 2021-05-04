package handlers

import (
	"net/http"
	"resources-ms/domain/resources"
	"resources-ms/repositories"

	"github.com/gorilla/mux"
)

func DeleteResource(res http.ResponseWriter, req *http.Request) {
	repository := repositories.ResourceRepository{}
	presenter := Presenter{Writer: res}

	useCase := resources.DeleteResourceUseCase{
		FindResource:   repository.FindResource,
		RemoveResource: repository.RemoveResource,
		PresentSuccess: presenter.Present,
	}

	params := mux.Vars(req)

	if apiError := useCase.Run(params["id"]); apiError != nil {
		presenter.PresentError(apiError)

		return
	}
}
