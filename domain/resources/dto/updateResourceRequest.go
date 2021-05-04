package dto

type UpdateResourceRequest struct {
	ID   string `json:"id" valid:"numeric,required~Invalid id"`
	Name string `json:"name,omitempty" valid:"required~Invalid name"`
	Url  string `json:"url,omitempty" valid:"url,required~Invalid url"`
}
