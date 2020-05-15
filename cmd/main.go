package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gojou/fbstart/pkg/things/contacts"

	"github.com/gorilla/mux"
)

var cnt = contacts.NewService()

func main() {
	log.Printf("Let's light this candle")

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func run() (e error) {

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

	c := cnt.NewContact()
	c.ID = "bbbb"
	c.FirstName = "Mark"
	c.LastName = "Poling"
	c.BirthYear = 1963
	c.BirthMonth = 11
	c.BirthDay = 29
	c.Email = "mark.poling@hvlst.com"
	log.Printf("Just created %v ",c)

	err := cnt.Add(c)
	if err != nil {
		e = err
		return
	}
	cs,_:=cnt.ListAll()
	cx:=cnt.NewContact()
	cx=cs[0]
	cx.ID="dddd"
	cx.FirstName="rhi"
	cx.Email = "rhi.poling@hvlst.com"
	cnt.Add(cx)
	log.Printf("Just added %v ",cx)

	router := mux.NewRouter()
	// THIS IS THE IMPORTANT LINE
	router.HandleFunc("/", c.Server)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))

	return nil
}
