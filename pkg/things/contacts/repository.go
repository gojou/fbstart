package contacts

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// Storage holds the pointer to the Firestore client
type Storage struct {
	DB  *firestore.Client
	Ctx context.Context
}

// NewStorage returns the pointer to the Firstore client
func NewStorage() (*Storage, error) {
	// Use the application default credentials
	s := &Storage{}
	s.Ctx = context.Background()

	client, e := firestore.NewClient(s.Ctx, "fbstart")
	if e != nil {
		log.Fatalln(e)
	}
	// defer client.Close()
	s.DB = client
	return s, e
}

// Add DODODO
func (s *Storage) Add(c *Contact) (e error) {

	_, _, e = s.DB.Collection("contacts").Add(context.Background(), c)
	if e != nil {
		log.Fatalf("Failed adding %v: %v", c.ID, e)
		//return e
	}
	log.Printf("Added Conact: %v", c.ID)
	return e
}

// ListAll returns all documents in a collection
func (s *Storage) ListAll() (cntcts []*Contact, e error) {
	log.Printf("Getting all contacts")
	iter := s.DB.Collection("contacts").Documents(s.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {

			return cntcts, nil

		}
		if err != nil {
			log.Printf("Failed to iterate over contacts: %v", err)
			e = err
			return cntcts, e
		}
		c := new(Contact)
		doc.DataTo(c)
		cntcts = append(cntcts, c)

	}
	// return cntcts, e
}

// Read DODODO
func (s *Storage) Read(cid string) (c *Contact, e error) {
	c.FirstName = "Neil"
	c.LastName = "Armstrong"
	return c, nil
}
