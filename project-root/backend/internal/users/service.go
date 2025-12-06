package users

import (
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repo *Repo
}

func NewService(r *Repo) *Service {
	return &Service{Repo: r}
}

// CreateUser hashes password and stores user.
func (s *Service) CreateUser(orgID, name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u := &User{
		OrganizationID: orgID,
		Name:           name,
		Email:          email,
		PasswordHash:   string(hash),
		Status:         "active",
	}
	if err := s.Repo.Create(u); err != nil {
		return nil, err
	}
	// Do not return password hash
	u.PasswordHash = ""
	return u, nil
}

// Authenticate verifies credentials and returns user if ok.
func (s *Service) Authenticate(email, password string) (*User, error) {
	u, err := s.Repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return nil, err
	}
	u.PasswordHash = ""
	return u, nil
}
