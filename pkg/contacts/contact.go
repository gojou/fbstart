package entities

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Contact is the base structure for Scouts, Leaders, and any other individual
// referenced by this application.
type Contact struct {
	ID         string `json:"id"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	BirthYear  int    `json:"birth_year"`
	BirthMonth int    `json:"birth_month"`
	BirthDay   int    `json:"birth_day"`
}

//New returns a pointer to a new Contact
func New(last string, first string, birth time.Time) *Contact {
	birthYear := birth.Year()
	birthMonth := int(birth.Month())
	birthDay := birth.Day()
	id := first + last + padInt(birthDay) + padInt(birthMonth)
	return &Contact{
		ID:         id,
		LastName:   last,
		FirstName:  first,
		BirthYear:  birthYear,
		BirthMonth: birthMonth,
		BirthDay:   birthDay,
	}
}

// Server serves
func (c Contact) Server(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>"+c.Greeter()+"</h1>")
}

// Greeter is a sample method
func (c Contact) Greeter() (greeting string) {
	greeting = "My name is " + c.FirstName + " " + c.LastName
	return
}

// GetAllContacts gets all Contact entities
// TODO: make the return slice a slice of pointers to contacts
func GetAllContacts() []Contact {

	return nil
}

func padInt(i int) string {
	s := strconv.Itoa(i)
	if len(s) < 2 {
		s = "0" + s
	}
	return s
}
