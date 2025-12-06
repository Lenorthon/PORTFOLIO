package users

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             string    `json:"id"`
	OrganizationID string    `json:"organization_id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	PasswordHash   string    `json:"-"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}

type Repo struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) Create(u *User) error {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	now := time.Now()
	_, err := r.DB.Exec(`
        INSERT INTO users (id, organization_id, name, email, password_hash, status, created_at)
        VALUES ($1,$2,$3,$4,$5,$6,$7)
    `, u.ID, u.OrganizationID, u.Name, u.Email, u.PasswordHash, u.Status, now)
	if err != nil {
		return err
	}
	u.CreatedAt = now
	return nil
}

func (r *Repo) GetByEmail(email string) (*User, error) {
	u := &User{}
	err := r.DB.QueryRow(`SELECT id, organization_id, name, email, password_hash, status, created_at FROM users WHERE email=$1`, email).
		Scan(&u.ID, &u.OrganizationID, &u.Name, &u.Email, &u.PasswordHash, &u.Status, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *Repo) GetByID(id string) (*User, error) {
	u := &User{}
	err := r.DB.QueryRow(`SELECT id, organization_id, name, email, password_hash, status, created_at FROM users WHERE id=$1`, id).
		Scan(&u.ID, &u.OrganizationID, &u.Name, &u.Email, &u.PasswordHash, &u.Status, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}
