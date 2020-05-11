package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gojou/fbstart/pkg/things/contacts"

	"github.com/gorilla/mux"

	"cloud.google.com/go/firestore"
)

var scout = contacts.NewContact("Poling", "Mark",
	time.Date(1963, time.November, 29, 0, 0, 0, 0, time.UTC))

func main() {
	log.Printf("Let's light this candle")

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func run() (e error) {

	// omitting explicit return value; e scoped in function return
	// initialize storage, in this case firestore

	e = initdb()
	if e != nil {
		return e
	}

	// initialize the web server
	e = initweb()
	if e != nil {
		return e
	}
	// e = nil
	return e
}

func initweb() (e error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	scout := contacts.NewContact("Poling", "Mark", time.Date(1963, time.November, 29, 0, 0, 0, 0, time.UTC))

	router := mux.NewRouter()
	// THIS IS THE IMPORTANT LINE
	router.HandleFunc("/", scout.Server)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))

	return nil
}

func initdb() (e error) {
	// Use the application default credentials

	ctx := context.Background()

	client, e := firestore.NewClient(ctx, "fbstart")
	//	defer client.Close()
	if e != nil {
		log.Fatalln(e)
		return e
	}

	e = useDB(ctx, *client)
	if e != nil {
		log.Fatalf("Failed: %v", e)
		return e
	}
	return nil
}

func useDB(ctx context.Context, db firestore.Client) (err error) {

	_, _, err = db.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
		return
	}
	log.Println("added Ada")

	_, _, err = db.Collection("users").Add(ctx, map[string]interface{}{
		"first":  "Alan",
		"middle": "Mathison",
		"last":   "Turing",
		"born":   1912,
	})
	log.Println("added Alan")
	if err != nil {
		log.Fatalf("Failed adding aturing: %v", err)
		return
	}
	return

}
