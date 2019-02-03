package main

import (
	"context"
	"firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"log"
)

func getFirebaseClient() {
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

	docs :=client.Collection("groupie_user_accounts").Documents(ctx)
	doc,_ := docs.Next()
	fmt.Println(doc.Data())
}
