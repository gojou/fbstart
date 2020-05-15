package contacts

// Service returns the Client service
type Service interface {
	Add(*Contact) error
	ListAll() ([]*Contact, error)
	Read(string) (*Contact, error)
	NewContact() *Contact
}

// Repository interface defines the methods that can be used on the Service
type Repository interface {
	Add(*Contact) error
	ListAll() ([]*Contact, error)
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

func (s service) Add(c *Contact) error {
	return s.r.Add(c)
}

func (s service) ListAll() ([]*Contact, error) {
	return s.r.ListAll()

}

func (s service) Read(cID string) (*Contact, error) {
	return s.r.Read(cID)
}

//NewContact returns a pointer to a new Contact
func (s service) NewContact() *Contact {
	contact := new(Contact)
	return contact
}
