package handlers

import (
	"encoding/json"
	"net/http"
	"resources-ms/domain/resources"
	"resources-ms/domain/resources/dto"
	"resources-ms/repositories"

	"github.com/rodriez/restface"
)

func CreateResource(res http.ResponseWriter, req *http.Request) {
	repository := repositories.ResourceRepository{}
	presenter := restface.Presenter{Writer: res}

	var request dto.CreateResourceRequest

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		presenter.PresentError(restface.BadRequest("Invalid JSON"))

		return
	}

	useCase := resources.CreateResourceUseCase{
		SaveResource:   repository.SaveResource,
		PresentSuccess: presenter.Present,
	}

	if apiError := useCase.Run(&request); apiError != nil {
		presenter.PresentError(apiError)

		return
	}
}
