package contacts

// Contact is the base structure for Scouts, Leaders, and any other individuals referenced by this application.
type Contact struct {
	ID         string `json:"id"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	BirthYear  int    `json:"birth_year"`
	BirthMonth int    `json:"birth_month"`
	BirthDay   int    `json:"birth_day"`
	Email      string `json:"email"`
}
