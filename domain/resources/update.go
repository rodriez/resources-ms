package resources

import (
	"fmt"
	"resources-ms/domain/resources/dto"
	"resources-ms/domain/resources/entities"
	"resources-ms/domain/resources/gateways"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/rodriez/restface"
)

type UpdateResourceUseCase struct {
	FindResource   func(string) (gateways.IResource, error)
	SaveResource   func(gateways.IResource) error
	PresentSuccess func(interface{})
}

func (uc *UpdateResourceUseCase) Run(req *dto.UpdateResourceRequest) *restface.ApiError {
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
		return restface.InternalError(err.Error())
	}

	uc.PresentSuccess(resource)

	return nil
}

func (uc *UpdateResourceUseCase) getResource(req *dto.UpdateResourceRequest) (*entities.Resource, *restface.ApiError) {
	res, err := uc.FindResource(req.ID)

	if err != nil {
		return nil, restface.InternalError(err.Error())
	}

	if res == nil {
		return nil, restface.NotFound(fmt.Sprintf("resource %s not found", req.ID))
	}

	return entities.Resource{}.Parse(res), nil
}

func (uc *UpdateResourceUseCase) validate(req *dto.UpdateResourceRequest) *restface.ApiError {
	valid, err := govalidator.ValidateStruct(req)

	if !valid {
		return restface.BadRequest(err.Error())
	}

	return nil
}
