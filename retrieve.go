package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

func getFirestoreClient() *firestore.Client{
	// Use the application default credentials.
	opt := option.WithCredentialsFile("serviceAccountKey.json")

	// Use context.Background() because the app/client should persist across
	// invocations.
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}
	return client
}

func getDocs(collectionName string, client *firestore.Client) *firestore.DocumentIterator{
	return client.Collection(collectionName).Documents(context.Background())
}
