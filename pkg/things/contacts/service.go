package contacts

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"cloud.google.com/go/firestore"
)

// Service returns the Client service
type Service interface {
	Create(Contact) error
	Read(string) (Contact, error)
}

// Repository interface defines the methods that can be used on the Service
type Repository interface {
	Create(Contact) error
	Read(string) (Contact, error)
}

type service struct {
	r Repository
}

// NewService returns a pointer to the Contact service
func NewService() Service {
	//TODO Deal with the error
	r, _ := NewStorage()
	return service{r}

}

func (s service) Create(c Contact) error {
	return s.r.Create(c)
}
func (s service) Read(cid string) (Contact, error) {
	return s.r.Read(cid)
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

// SaveContact will save the contact created in the main func
// TODO Pull in all the junk to make this work.
func SaveContact(c *Contact, store firestore.Client) (e error) {
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
