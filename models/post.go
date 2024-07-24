package models

type Post struct {
	PostID        string `json:"postId"`
	PostTitle     string `json:"postTitle"`
	PostLink      string `json:"postLink"`
	PostThumbnail string `json:"postThumbnail"`
	PostContent   string `json:"postContent"`
	Timestamp     string `json:"timestamp"`
}

type PostUploadContent struct {
	PostTitle     string `json:"postTitle"`
	PostThumbnail string `json:"postThumbnail"`
	Content       string `json:"content"`
}
