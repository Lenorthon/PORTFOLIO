package auth

import "time"

type User struct {
	ID           string    `json:"id"`
	OrgID        string    `json:"org_id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}

type Invite struct {
	ID        string    `json:"id"`
	OrgID     string    `json:"org_id"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	Accepted  bool      `json:"accepted"`
	CreatedAt time.Time `json:"created_at"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name,omitempty"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
