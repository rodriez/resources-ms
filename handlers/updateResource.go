package handlers

import (
	"encoding/json"
	"net/http"
	"resources-ms/domain/resources"
	"resources-ms/domain/resources/dto"
	"resources-ms/repositories"

	"github.com/gorilla/mux"
	"github.com/rodriez/restface"
)

func UpdateResource(res http.ResponseWriter, req *http.Request) {
	repository := repositories.ResourceRepository{}
	presenter := restface.Presenter{Writer: res}

	var request dto.UpdateResourceRequest

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		presenter.PresentError(restface.BadRequest("Invalid JSON"))

		return
	}

	useCase := resources.UpdateResourceUseCase{
		FindResource:   repository.FindResource,
		SaveResource:   repository.SaveResource,
		PresentSuccess: presenter.Present,
	}

	params := mux.Vars(req)
	request.ID = params["id"]

	if apiError := useCase.Run(&request); apiError != nil {
		presenter.PresentError(apiError)

		return
	}
}
