package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gojou/fbstart/pkg/entities"

	firebase "firebase.google.com/go"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func run() (e error) {
	parent := new(entities.Contact)
	parent.LastName = "Poling"
	parent.FirstName = "Mark"
	fmt.Println(parent.Greeter())
	scout := entities.New("Poling", "Aden")
	fmt.Println("Hello world! " + scout.Greeter())
	e = initdb()
	if e != nil {
		return
	}
	e = errors.New("oops i did it again")
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
		e = err
		return
	}

	client, err := app.Firestore(ctx)
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
