package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gojou/fbstart/pkg/entities"
	"github.com/gojou/fbstart/pkg/handlers"

	"cloud.google.com/go/firestore"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func run() (e error) {
	scout := entities.New("Poling", "Aden")
	fmt.Println("Hello world! " + scout.Greeter())

	// initialize the web server
	w = initweb()
	if e != nil {
		return
	}

	// initialize storage, in this case firestore
	e = initdb()
	if e != nil {
		return
	}
	e = errors.New("oops i did it again")
	return // omitting explicit return value; e scoped in function call
}

func initweb() (e error) {
	http.HandleFunc("/", handlers.Index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	return nil
}

func initdb() (e error) {
	// Use the application default credentials

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, "fbstart")
	if err != nil {
		log.Fatalln(err)
		e = err
		return
	}
	defer client.Close()

	_, _, e = client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if e != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
		return
	}

	_, _, e = client.Collection("users").Add(ctx, map[string]interface{}{
		"first":  "Alan",
		"middle": "Mathison",
		"last":   "Turing",
		"born":   1912,
	})
	if e != nil {
		log.Fatalf("Failed adding aturing: %v", err)
		return
	}
	return
}
