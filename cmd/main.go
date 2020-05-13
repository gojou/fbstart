package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gojou/fbstart/pkg/things/contacts"

	"github.com/gorilla/mux"

	"cloud.google.com/go/firestore"
)

// Establishing but NOT initializing the contacts Service
// scope is package wide; will refactor.
var cnt = contacts.NewService()

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
	scout := contacts.Contact{ID: "zzzz"}
	scout.FirstName = "Mark"
	scout.LastName = "Poling"
	scout.BirthYear = 1963
	scout.BirthMonth = 11
	scout.BirthDay = 29
	scout.Email = "mark.poling@gmail.com"

	err := cnt.Create(scout)
	if err != nil {
		e = err
		return
	}

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
		return e
	}
	return e // nil
}

func useDB(ctx context.Context, db firestore.Client) (e error) {

	type User struct {
		First  string `firestore:"first"`
		Middle string `firestore:"middle"`
		Last   string `firestore:"last"`
		Born   int64  `firestore:"born"`
	}
	ada := User{
		First:  "Ada",
		Middle: "Rhiannon",
		Last:   "Lovelace",
		Born:   1815,
	}

	alan := User{
		First:  "Alan",
		Middle: "Mathis",
		Last:   "Turing",
		Born:   1912,
	}
	users := []User{ada, alan}

	for _, u := range users {
		_, _, err := db.Collection("users").Add(ctx, u)
		if err != nil {
			log.Fatalf("Failed adding %v: %v", u.First, err)

		}
		log.Printf("Added %v\n", u.First)

	}

	return e
}
