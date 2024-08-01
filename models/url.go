package models

type FetchUrlRequest struct {
	URL string `json:"url" binding:"required"`
}

type Metadata struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type FetchUrlResponse struct {
	Success int      `json:"success"`
	Meta    Metadata `json:"meta"`
}
