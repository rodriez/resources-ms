package entities

import (
	"resources-ms/domain/resources/gateways"
	"time"
)

type Resource struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Url     string    `json:"url"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

func (r *Resource) GetID() string {
	return r.ID
}

func (r *Resource) GetName() string {
	return r.Name
}

func (r *Resource) GetUrl() string {
	return r.Url
}

func (r *Resource) GetCreated() time.Time {
	return r.Created
}

func (r *Resource) GetUpdated() time.Time {
	return r.Updated
}

func (r Resource) Parse(res gateways.IResource) *Resource {
	r.ID = res.GetID()
	r.Name = res.GetName()
	r.Url = res.GetUrl()
	r.Created = res.GetCreated()
	r.Updated = res.GetUpdated()

	return &r
}
