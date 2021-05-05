package resources

import (
	"math/rand"
	"resources-ms/domain/resources/dto"
	"resources-ms/domain/resources/entities"
	"resources-ms/domain/resources/gateways"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/rodriez/restface"
)

type CreateResourceUseCase struct {
	SaveResource   func(gateways.IResource) error
	PresentSuccess func(interface{})
}

func (uc *CreateResourceUseCase) Run(req *dto.CreateResourceRequest) *restface.ApiError {
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
		return restface.InternalError(err.Error())
	}

	uc.PresentSuccess(resource)

	return nil
}

func (uc *CreateResourceUseCase) validate(req *dto.CreateResourceRequest) *restface.ApiError {
	if valid, err := govalidator.ValidateStruct(req); !valid {
		return restface.BadRequest(err.Error())
	}

	return nil
}
