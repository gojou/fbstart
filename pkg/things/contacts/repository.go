package contacts

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// Storage holds the pointer to the Firestore client
type Storage struct {
	db *firestore.Client
}

// NewStorage returns the pointer to the Firstore client
func NewStorage() (*Storage, error) {
	// Use the application default credentials
	storage := new(Storage)
	ctx := context.Background()

	client, e := firestore.NewClient(ctx, "fbstart")
	//	defer client.Close()
	if e != nil {
		log.Fatalln(e)
	}
	storage.db = client
	return storage, e
}

// Create DODODO
func (s *Storage) Create(c Contact) error {
	//TODO Implement the Create function
	return nil
}

// Read DODODO
func (s *Storage) Read(cid string) (c Contact, e error) {
	c.FirstName = "Neil"
	c.LastName = "Armstrong"
	return c, nil
}
