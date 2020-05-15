package contacts

import (
	"cloud.google.com/go/firestore"
)

// Service returns the Client service
type Service interface {
	NewContact() *Contact
	Add(*Contact) error
	Read(string) (Contact, error)
}

// Repository interface defines the methods that can be used on the Service
type Repository interface {
	Add(*Contact) error
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

func (s service) Add(c *Contact) error {
	return s.r.Add(c)
}
func (s service) Read(cID string) (Contact, error) {
	return s.r.Read(cID)
}

//NewContact returns new Contact
func (s service) NewContact() *Contact {
	contact := new(Contact)
	return contact
}

// GetAllContacts gets all Contact entities
// TODO: make the return slice a slice of pointers to contacts
func GetAllContacts() []Contact {

	return nil
}

// SaveContact will save the contact created in the main func
// TODO Pull in all the junk to make this work.
func SaveContact(c Contact, store firestore.Client) (e error) {
	return nil
}
