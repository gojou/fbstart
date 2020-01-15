package entities

// Contact is the base structure for Scouts, Leaders, and any other individual
// referenced by this application.
type Contact struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}

//New returns a pointer to a new Contact
func New(Last string, First string) *Contact {
	return &Contact{
		LastName:  Last,
		FirstName: First,
	}
}

// Greeter is a sample method
func (c Contact) Greeter() (greeting string) {
	greeting = "My name is " + c.FirstName + " " + c.LastName
	return
}
