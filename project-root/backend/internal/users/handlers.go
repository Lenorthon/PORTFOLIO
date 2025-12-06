package users

import (
	"encoding/json"
	"net/http"

	"github.com/Lenorthon/PORTFOLIO/backend/internal/auth"
)

// Handlers expects a Service instance to be provided by caller.
type Handlers struct {
	Svc *Service
}

func NewHandlers(svc *Service) *Handlers {
	return &Handlers{Svc: svc}
}

// createUserRequest represents payload for user creation
type createUserRequest struct {
	OrganizationID string `json:"organization_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
}

// CreateUserHandler creates a new user
func (h *Handlers) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	user, err := h.Svc.CreateUser(req.OrganizationID, req.Name, req.Email, req.Password)
	if err != nil {
		http.Error(w, "could not create user: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// loginRequest for authentication
type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginHandler authenticates and returns JWT
func (h *Handlers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	user, err := h.Svc.Authenticate(req.Email, req.Password)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	// generate token
	token, err := auth.GenerateToken(user.ID, user.OrganizationID, 24*0) // zero uses default in GenerateToken
	if err != nil {
		http.Error(w, "could not generate token", http.StatusInternalServerError)
		return
	}
	resp := map[string]interface{}{
		"token": token,
		"user":  user,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
