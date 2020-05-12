package contacts

import (
	"fmt"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore"
)

// Service returns the Client service
type Service interface {
	Create(Contact) error
	Read(string) (*Contact, error)
}

// Repository interface defines the methods that can be used on the Service
type Repository interface {
	Create(Contact) error
	Read(string) (*Contact, error)
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

func (S service) Create(c Contact) error {
	return S.r.Create(c)
}
func (S service) Read(cID string) (*Contact, error) {
	return S.r.Read(cID)
}

//NewContact returns a pointer to a new Contact
func NewContact(id string, last string,
	first string, birthyear int, birthmonth int, birthday int,
	email string) *Contact {
	return &Contact{
		ID:         id,
		LastName:   last,
		FirstName:  first,
		BirthYear:  birthyear,
		BirthMonth: birthmonth,
		BirthDay:   birthday,
		Email:      email,
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
