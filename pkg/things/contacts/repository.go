package contacts

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// Storage holds the pointer to the Firestore client
type Storage struct {
	DB  *firestore.Client
	Ctx context.Context
}

// NewStorage returns the pointer to the Firstore client
func NewStorage() (*Storage, error) {
	// Use the application default credentials
	s := new(Storage)
	s.Ctx = context.Background()

	client, e := firestore.NewClient(s.Ctx, "fbstart")
	//	defer client.Close()
	if e != nil {
		log.Fatalln(e)
	}
	s.DB = client
	return s, e
}

// Create DODODO
func (s *Storage) Add(c Contact) (e error) {

	_, _, e = s.DB.Collection("contacts").Add(context.Background(), c)
	if e != nil {
		log.Fatalf("Failed adding %v: %v", c.ID, e)
		//return e
	}
	log.Printf("Added Conact: %v", c.ID)
	return e
}

// Read DODODO
func (s *Storage) Read(cid string) (c Contact, e error) {
	c.FirstName = "Neil"
	c.LastName = "Armstrong"
	return c, nil
}
