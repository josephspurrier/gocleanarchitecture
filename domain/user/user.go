package user

// Item represents a user.
type Item struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Service represents a service for managing users.
type Service interface {
	User(email string) (*Item, error)
	CreateUser(user *Item) error
	Authenticate(user *Item) error
}
