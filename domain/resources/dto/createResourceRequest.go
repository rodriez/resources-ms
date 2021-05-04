package dto

type CreateResourceRequest struct {
	Name string `json:"name,omitempty" valid:"required~Invalid name"`
	Url  string `json:"url,omitempty" valid:"url,required~Invalid url"`
}
