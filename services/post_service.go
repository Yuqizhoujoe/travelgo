package services

import (
	"context"
	"fmt"
	"time"

	"travelgo/firebase"
	"travelgo/models"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type PostService struct{}

func NewPostService() PostService {
	return PostService{}
}

func (ps *PostService) CreatePost(content models.PostUploadContent) (postID string, err error) {
	ctx := context.Background()

	// generate new document reference with an auto-generated ID
	docRef := firebase.FirestoreClient.Collection("posts").NewDoc()

	newPost := models.Post{
		PostID:    docRef.ID,
		PostTitle: content.PostTitle,
		// PostThumbnail: gcsURL,
		// PostLink:  "https://platform.com/posts/" + docRef.ID,
		EditorJsData: content.EditorJsData,
		Timestamp:    time.Now().Format(time.RFC3339),
	}

	fmt.Println("Post Service")
	fmt.Println(newPost)

	/*
	 Set(): replace the entire document
	*/
	_, err = docRef.Set(ctx, newPost)

	return newPost.PostID, err
}

func (ps *PostService) UpdatePost(id string, content models.PostUploadContent) (postID string, err error) {
	ctx := context.Background()

	post, err := ps.GetPostContent(id)
	if err != nil {
		return "", err
	}

	// Update the post with new content
	// Create a slice of updates
	updates := []firestore.Update{
		{Path: "PostTitle", Value: content.PostTitle},
		{Path: "EditorJsData", Value: content.EditorJsData},
		{Path: "Timestamp", Value: time.Now().Format(time.RFC3339)},
	}

	docRef := firebase.FirestoreClient.Collection("posts").Doc(id)
	_, err = docRef.Update(ctx, updates)

	if err != nil {
		return "", err
	}

	return post.PostID, nil
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
			PostID:    postContent.PostID,
			PostTitle: postContent.PostTitle,
			// PostLink:      postContent.PostLink,
			// PostThumbnail: postContent.PostThumbnail,
			Timestamp: time.Now().Format(time.RFC3339),
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
