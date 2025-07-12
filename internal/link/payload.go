package link

type LinkCreateRequest struct {
	URL string `json:"url" validate:"required,url"`
}

type LinkUpdateRequest struct {
	URL  string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}
