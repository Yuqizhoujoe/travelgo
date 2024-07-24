package firebase

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/storage"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App
var AuthClient *auth.Client
var FirestoreClient *firestore.Client
var StorageClient *storage.Client

func InitFirebase() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("../firebase-adminsdk.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	FirebaseApp = app

	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	AuthClient = authClient

	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error getting Firestore client: %v\n", err)
	}
	FirestoreClient = firestoreClient

	storageClient, err := app.Storage(ctx)
	if err != nil {
		log.Fatalf("error getting Store client: %v\n", err)
	}
	StorageClient = storageClient
}
