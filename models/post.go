package models

type Post struct {
	PostID    string `json:"postId"`
	PostTitle string `json:"postTitle"`
	// PostLink      string `json:"postLink"`
	// PostThumbnail string `json:"postThumbnail"`
	EditorJsData EditorData `json:"editorJsData"`
	Timestamp    string     `json:"timestamp"`
}

type PostUploadContent struct {
	PostTitle    string     `json:"postTitle"`
	EditorJsData EditorData `json:"editorJsData"`
}

type EditorUrl struct {
	URL string `json:"url" binding:"required"`
}

type PostResponse struct {
	Success bool   `json:"success"`
	PostID  string `json:"postId"`
}
