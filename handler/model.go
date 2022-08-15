package handler

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}
