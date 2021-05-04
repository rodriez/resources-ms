package resources

import (
	"fmt"
	"resources-ms/domain/resources/entities"
	"resources-ms/domain/resources/exceptions"
	"resources-ms/domain/resources/gateways"
)

type DeleteResourceUseCase struct {
	FindResource   func(string) (gateways.IResource, error)
	RemoveResource func(string) error
	PresentSuccess func(gateways.IResource)
}

func (uc *DeleteResourceUseCase) Run(id string) *exceptions.ApiError {
	if err := uc.validate(id); err != nil {
		return err
	}

	resource, err := uc.getResource(id)
	if err != nil {
		return err
	}

	if err := uc.RemoveResource(id); err != nil {
		return exceptions.InternalError(err.Error())
	}

	uc.PresentSuccess(resource)

	return nil
}

func (uc *DeleteResourceUseCase) getResource(id string) (*entities.Resource, *exceptions.ApiError) {
	res, err := uc.FindResource(id)

	if err != nil {
		return nil, exceptions.InternalError(err.Error())
	}

	if res == nil {
		return nil, exceptions.NotFound(fmt.Sprintf("resource %s not found", id))
	}

	return res.(*entities.Resource), nil
}

func (uc *DeleteResourceUseCase) validate(id string) *exceptions.ApiError {
	if id == "" {
		return exceptions.BadRequest("Invalid id")
	}

	return nil
}
