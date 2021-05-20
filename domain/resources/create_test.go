package resources_test

import (
	"resources-ms/domain/resources"
	"resources-ms/domain/resources/dto"
	"resources-ms/domain/resources/gateways"
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	req *dto.CreateResourceRequest
}

func Test_Given_Valid_Name_And_Url_Save_Resource_Successfully(t *testing.T) {
	//Given
	mockSave := func(p gateways.IResource) error {
		return nil
	}
	mockPresent := func(p interface{}) {
		return
	}

	uc := &resources.CreateResourceUseCase{
		SaveResource:   mockSave,
		PresentSuccess: mockPresent,
	}

	req := &dto.CreateResourceRequest{
		Name: "products-api",
		Url:  "http://example.com/api/products",
	}

	err := uc.Run(req)

	assert.Empty(t, err)
}

//
func Test_Given_Empty_Name_And_Valid_Url_Return_Invalid_Name_Error(t *testing.T) {
	//Given
	req := &dto.CreateResourceRequest{
		Name: "",
		Url:  "http://example.com/api/products",
	}

	mockSave := func(p gateways.IResource) error {
		return nil
	}
	mockPresent := func(p interface{}) {
		return
	}

	uc := &resources.CreateResourceUseCase{
		SaveResource:   mockSave,
		PresentSuccess: mockPresent,
	}

	//When
	err := uc.Run(req)

	//Then
	assert.EqualError(t, err, "Bad request")
	assert.Equal(t, err.Errors[0], "Invalid name")
}
