package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gojou/fbstart/pkg/contacts"
	"github.com/gorilla/mux"

	"cloud.google.com/go/firestore"
)

var scout = entities.New("Poling", "Aden", time.Date(2007, time.May, 23, 0, 0, 0, 0, time.UTC))

func main() {
	log.Printf("Let's light this candle")

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func run() (e error) {

	fmt.Println("Hello world! " + scout.Greeter())
	// omitting explicit return value; e scoped in function call
	// initialize storage, in this case firestore
	e = initdb()
	if e != nil {
		return
	}

	// initialize the web server
	e = initweb()
	if e != nil {
		return
	}
	// e = nil
	return
}

func initweb() (e error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)

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
	defer client.Close()
	if e != nil {
		log.Fatalln(e)
		return e
	}

	e = useDB(ctx, *client)
	if e != nil {
		log.Fatalf("Failed: %v", e)
		return
	}
	return
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
