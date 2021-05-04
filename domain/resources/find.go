package resources

import (
	"fmt"
	"resources-ms/domain/resources/exceptions"
	"resources-ms/domain/resources/gateways"
)

type FindResourceUseCase struct {
	FindResource   func(string) (gateways.IResource, error)
	PresentSuccess func(gateways.IResource)
}

func (uc *FindResourceUseCase) Run(id string) *exceptions.ApiError {
	if err := uc.validate(id); err != nil {
		return err
	}

	resource, err := uc.FindResource(id)

	if err != nil {
		return exceptions.InternalError(err.Error())
	}

	if resource == nil {
		return exceptions.NotFound(fmt.Sprintf("resource %s not found", id))
	}

	uc.PresentSuccess(resource)

	return nil
}

func (uc *FindResourceUseCase) validate(id string) *exceptions.ApiError {
	if id == "" {
		return exceptions.BadRequest("Invalid id")
	}

	return nil
}
