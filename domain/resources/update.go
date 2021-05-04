package resources

import (
	"fmt"
	"net/http"
	"resources-ms/domain/resources/dto"
	"resources-ms/domain/resources/entities"
	"resources-ms/domain/resources/exceptions"
	"resources-ms/domain/resources/gateways"
	"time"

	"github.com/asaskevich/govalidator"
)

type UpdateResourceUseCase struct {
	FindResource   func(string) (gateways.IResource, error)
	SaveResource   func(gateways.IResource) error
	PresentSuccess func(gateways.IResource)
}

func (uc *UpdateResourceUseCase) Run(req *dto.UpdateResourceRequest) *exceptions.ApiError {
	if err := uc.validate(req); err != nil {
		return err
	}

	resource, err := uc.getResource(req)
	if err != nil {
		return err
	}

	resource.Name = req.Name
	resource.Url = req.Url
	resource.Updated = time.Now()

	if err := uc.SaveResource(resource); err != nil {
		return &exceptions.ApiError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	uc.PresentSuccess(resource)

	return nil
}

func (uc *UpdateResourceUseCase) getResource(req *dto.UpdateResourceRequest) (*entities.Resource, *exceptions.ApiError) {
	res, err := uc.FindResource(req.ID)

	if err != nil {
		return nil, exceptions.InternalError(err.Error())
	}

	if res == nil {
		return nil, exceptions.NotFound(fmt.Sprintf("resource %s not found", req.ID))
	}

	return res.(*entities.Resource), nil
}

func (uc *UpdateResourceUseCase) validate(req *dto.UpdateResourceRequest) *exceptions.ApiError {
	valid, err := govalidator.ValidateStruct(req)

	if !valid {
		return exceptions.BadRequest(err.Error())
	}

	return nil
}
