package auth

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// In-memory storage pour exemple, à remplacer par DB SQLC
var users = map[string]*User{}
var invites = map[string]*Invite{}

func RegisterUser(email, password, name string, orgID string) (*User, error) {
	for _, u := range users {
		if u.Email == email {
			return nil, errors.New("email already exists")
		}
	}

	hash, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:           uuid.New().String(),
		Email:        email,
		Name:         name,
		OrgID:        orgID,
		PasswordHash: hash,
		Role:         "member",
	}

	users[user.ID] = user
	return user, nil
}

func LoginUser(email, password string) (*User, error) {
	for _, u := range users {
		if u.Email == email && CheckPasswordHash(password, u.PasswordHash) {
			return u, nil
		}
	}
	return nil, errors.New("invalid credentials")
}

func CreateInvite(orgID, email string) (*Invite, error) {
	token := uuid.New().String()
	invite := &Invite{
		ID:        uuid.New().String(),
		OrgID:     orgID,
		Email:     email,
		Token:     token,
		Accepted:  false,
		CreatedAt: time.Now(),
	}
	invites[invite.ID] = invite
	return invite, nil
}

func AcceptInvite(token, name, password string) (*User, error) {
	for _, inv := range invites {
		if inv.Token == token && !inv.Accepted {
			user, err := RegisterUser(inv.Email, password, name, inv.OrgID)
			if err != nil {
				return nil, err
			}
			inv.Accepted = true
			return user, nil
		}
	}
	return nil, errors.New("invalid or expired invite")
}
