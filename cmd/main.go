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

	scout := contacts.NewContact("PolingM", "Poling", "Mark", 1963, 11, 29, "mark@poling.com")

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
		// _, e = db.Collection("users").Doc(ada.First).Set(ctx, ada)
		if err != nil {
			log.Fatalf("Failed adding %v: %v", u.First, err)

		}

	}

	// _, err := db.Collection("users").Doc("one").Set(ctx, ada)
	// if err != nil {
	// 	log.Fatalf("Failed adding alovelace: %v", err)
	// 	return
	// }
	// log.Println("added Ada")

	// _, _, e = db.Collection("users").Add(ctx, map[string]interface{}{
	// 	"first": "Ada",
	// 	"last":  "Lovelace",
	// 	"born":  1815,
	// })
	// if e != nil {
	// 	log.Fatalf("Failed adding alovelace: %v", e)
	// 	return
	// }
	// log.Println("added Ada")

	//	_, e = db.Collection("users").Doc(alan.First).Set(ctx, alan)
	// _, _, e := db.Collection("users").Add(ctx, ada)
	//
	// log.Println("added Alan")
	// if e != nil {
	// 	log.Fatalf("Failed adding aturing: %v", e)
	// 	return
	// }
	// return
	return e
}
