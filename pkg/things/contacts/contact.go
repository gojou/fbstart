package contacts

// Contact is the base structure for Scouts, Leaders, and any other individuals referenced by this application.
type Contact struct {
	ID         string `firestore:"id"`
	LastName   string `firestore:"last_name"`
	FirstName  string `firestore:"first_name"`
	BirthYear  int    `firestore:"birth_year"`
	BirthMonth int    `firestore:"birth_month"`
	BirthDay   int    `firestore:"birth_day"`
	Email      string `firestore:"email"`
}
