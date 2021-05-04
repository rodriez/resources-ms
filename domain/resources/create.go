package resources

import (
	"math/rand"
	"resources-ms/domain/resources/dto"
	"resources-ms/domain/resources/entities"
	"resources-ms/domain/resources/exceptions"
	"resources-ms/domain/resources/gateways"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateResourceUseCase struct {
	SaveResource   func(gateways.IResource) error
	PresentSuccess func(gateways.IResource)
}

func (uc *CreateResourceUseCase) Run(req *dto.CreateResourceRequest) *exceptions.ApiError {
	if err := uc.validate(req); err != nil {
		return err
	}

	resource := &entities.Resource{
		ID:      strconv.Itoa(rand.Intn(100000000)),
		Name:    req.Name,
		Url:     req.Url,
		Created: time.Now(),
		Updated: time.Now(),
	}

	if err := uc.SaveResource(resource); err != nil {
		return exceptions.InternalError(err.Error())
	}

	uc.PresentSuccess(resource)

	return nil
}

func (uc *CreateResourceUseCase) validate(req *dto.CreateResourceRequest) *exceptions.ApiError {
	if valid, err := govalidator.ValidateStruct(req); !valid {
		return exceptions.BadRequest(err.Error())
	}

	return nil
}
