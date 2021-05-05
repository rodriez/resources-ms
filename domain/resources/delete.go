package resources

import (
	"fmt"
	"resources-ms/domain/resources/entities"
	"resources-ms/domain/resources/gateways"

	"github.com/rodriez/restface"
)

type DeleteResourceUseCase struct {
	FindResource   func(string) (gateways.IResource, error)
	RemoveResource func(string) error
	PresentSuccess func(interface{})
}

func (uc *DeleteResourceUseCase) Run(id string) *restface.ApiError {
	if err := uc.validate(id); err != nil {
		return err
	}

	resource, err := uc.getResource(id)
	if err != nil {
		return err
	}

	if err := uc.RemoveResource(id); err != nil {
		return restface.InternalError(err.Error())
	}

	uc.PresentSuccess(resource)

	return nil
}

func (uc *DeleteResourceUseCase) getResource(id string) (*entities.Resource, *restface.ApiError) {
	res, err := uc.FindResource(id)

	if err != nil {
		return nil, restface.InternalError(err.Error())
	}

	if res == nil {
		return nil, restface.NotFound(fmt.Sprintf("resource %s not found", id))
	}

	return entities.Resource{}.Parse(res), nil
}

func (uc *DeleteResourceUseCase) validate(id string) *restface.ApiError {
	if id == "" {
		return restface.BadRequest("Invalid id")
	}

	return nil
}
