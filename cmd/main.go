package main

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func run() (e error) {
	fmt.Println("Hello world!")
	return // omitting explicit return value; e scoped in function call
}

func initdb() (e error) {
	// Use the application default credentials

	ctx := context.Background()
	// test the following to see if the format of the command is correct
	conf := &firebase.Config{ProjectID: "fbstart"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer client.Close()
	return
}
