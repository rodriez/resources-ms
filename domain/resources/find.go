package resources

import (
	"fmt"
	"resources-ms/domain/resources/gateways"

	"github.com/rodriez/restface"
)

type FindResourceUseCase struct {
	FindResource   func(string) (gateways.IResource, error)
	PresentSuccess func(interface{})
}

func (uc *FindResourceUseCase) Run(id string) *restface.ApiError {
	if err := uc.validate(id); err != nil {
		return err
	}

	resource, err := uc.FindResource(id)

	if err != nil {
		return restface.InternalError(err.Error())
	}

	if resource == nil {
		return restface.NotFound(fmt.Sprintf("resource %s not found", id))
	}

	uc.PresentSuccess(resource)

	return nil
}

func (uc *FindResourceUseCase) validate(id string) *restface.ApiError {
	if id == "" {
		return restface.BadRequest("Invalid id")
	}

	return nil
}
