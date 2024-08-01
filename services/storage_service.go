package services

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"
	"travelgo/firebase"

	"firebase.google.com/go/storage"

	CloudStorage "cloud.google.com/go/storage"
)

type StorageService struct {
	client *storage.Client
}

func NewStorageService() StorageService {
	return StorageService{
		client: firebase.StorageClient,
	}
}

func (ss *StorageService) UploadFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	defer file.Close()

	// Create a context
	ctx := context.Background()

	// Create a bucket handle
	bucket, err := ss.client.DefaultBucket()
	if err != nil {
		return "", fmt.Errorf("failed to get default bucket: %v", err)
	}

	// Create file handle
	object := bucket.Object(header.Filename)

	// Upload the file
	wc := object.NewWriter(ctx)
	// wc.ContentType = file.Header.Get("Content-Type")
	// wc.Metadata = map[string]string{
	// 	"firebaseStorageDownloadTokens": uuid.New().String(),
	// }

	if _, err = io.Copy(wc, file); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	// Get the public URL of the uploaded file
	opts := &CloudStorage.SignedURLOptions{
		Method:  "GET",
		Expires: time.Now().AddDate(1, 0, 0), // expires in 1 year
	}

	url, err := bucket.SignedURL(header.Filename, opts)
	if err != nil {
		return "", err
	}

	return url, nil
}
