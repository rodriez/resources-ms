package repositories

import "resources-ms/domain/resources/gateways"

type ResourceRepository struct{}

var storage map[string]gateways.IResource

func init() {
	storage = make(map[string]gateways.IResource)
}

func (repo *ResourceRepository) SaveResource(res gateways.IResource) error {
	storage[res.GetID()] = res

	return nil
}

func (repo *ResourceRepository) FindResource(id string) (gateways.IResource, error) {

	return storage[id], nil
}

func (repo *ResourceRepository) RemoveResource(id string) error {
	delete(storage, id)

	return nil
}
