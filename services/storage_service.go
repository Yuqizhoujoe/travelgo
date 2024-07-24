package services

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"travelgo/firebase"

	"firebase.google.com/go/storage"
	"github.com/google/uuid"
)

type StorageService struct {
	bucketName string
	client     *storage.Client
}

func NewStorageService() StorageService {
	return StorageService{
		bucketName: "firebase-bucket",
		client:     firebase.StorageClient,
	}
}

func (ss *StorageService) UploadFile(file *multipart.FileHeader) (string, error) {
	ctx := context.Background()

	bucket, err := ss.client.DefaultBucket()
	if err != nil {
		return "", fmt.Errorf("failed to get default bucket: %v", err)
	}

	fileID := uuid.New().String()
	fileName := fmt.Sprintf("%s/%s", fileID, file.Filename)

	wc := bucket.Object(fileName).NewWriter(ctx)
	wc.ContentType = file.Header.Get("Content-Type")
	wc.Metadata = map[string]string{
		"firebaseStorageDownloadTokens": uuid.New().String(),
	}

	fileHandle, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileHandle.Close()

	if _, err = io.Copy(wc, fileHandle); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	downloadURL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media", ss.bucketName, fileName)
	return downloadURL, nil
}
