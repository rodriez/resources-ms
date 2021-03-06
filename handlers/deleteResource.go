package handlers

import (
	"net/http"
	"resources-ms/domain/resources"
	"resources-ms/repositories"

	"github.com/gorilla/mux"
	"github.com/rodriez/restface"
)

func DeleteResource(res http.ResponseWriter, req *http.Request) {
	repository := repositories.ResourceRepository{}
	presenter := restface.Presenter{Writer: res}

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
