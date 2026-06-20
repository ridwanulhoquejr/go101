package user

import "example.com/go101/interfaces/notification"

type User struct {
	Name string
	ID   string
}

type Repository struct {
	DB string
}

func NewRepository(db string) *Repository {
	return &Repository{DB: db}
}

// interface compilance
// compile time check to know Repository implements notification.UserFetcher
var _ notification.UserFetcher = (*Repository)(nil)

func (r *Repository) GetEmail(id string) (string, error) {
	return "abc@gmail.com", nil
}
