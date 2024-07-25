package services

import (
	"context"
	"time"

	"travelgo/firebase"
	"travelgo/models"

	"google.golang.org/api/iterator"
)

type PostService struct{}

func NewPostService() PostService {
	return PostService{}
}

func (ps *PostService) CreatePost(content models.PostUploadContent, gcsURL string) (postID string, err error) {
	ctx := context.Background()

	// generate new document reference with an auto-generated ID
	docRef := firebase.FirestoreClient.Collection("posts").NewDoc()

	newPost := models.Post{
		PostID:        docRef.ID,
		PostTitle:     content.PostTitle,
		PostThumbnail: gcsURL,
		PostLink:      "https://platform.com/posts/" + docRef.ID,
		PostContent:   content.Content,
		Timestamp:     time.Now().Format(time.RFC3339),
	}

	_, err = docRef.Set(ctx, newPost)

	return newPost.PostID, err
}

func (ps *PostService) GetPosts() ([]models.Post, error) {
	ctx := context.Background()
	iter := firebase.FirestoreClient.Collection("posts").Documents(ctx)
	var posts []models.Post

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		var postContent models.Post
		doc.DataTo(&postContent)
		posts = append(posts, models.Post{
			PostID:        postContent.PostID,
			PostTitle:     postContent.PostTitle,
			PostLink:      postContent.PostLink,
			PostThumbnail: postContent.PostThumbnail,
			Timestamp:     time.Now().Format(time.RFC3339),
		})
	}

	return posts, nil
}

func (ps *PostService) GetPostContent(postID string) (models.Post, error) {
	ctx := context.Background()
	doc, err := firebase.FirestoreClient.Collection("posts").Doc(postID).Get(ctx)
	if err != nil {
		return models.Post{}, err
	}

	var postContent models.Post
	doc.DataTo(&postContent)
	return postContent, nil
}
