package things

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Contact is the base structure for Scouts, Leaders, and any other individuals referenced by this application.
type Contact struct {
	ID         string `json:"id"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	BirthYear  int    `json:"birth_year"`
	BirthMonth int    `json:"birth_month"`
	BirthDay   int    `json:"birth_day"`
}

//NewContact returns a pointer to a new Contact
func NewContact(last string, first string, birth time.Time) *Contact {
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
	greeting = greeting + " " + c.ID
	return
}

// GetAllContacts gets all Contact entities
// TODO: make the return slice a slice of pointers to contacts
func GetAllContacts() []Contact {

	return nil
}

// padInt utility function adds leading zeros to string representations of
// months and days <10 characters long.
func padInt(i int) string {
	s := strconv.Itoa(i)
	if len(s) < 2 {
		s = "0" + s
	}
	return s
}
