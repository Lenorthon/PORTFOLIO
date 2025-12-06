package auth

import (
	"context"
	"net/http"
	"strings"
)

type ctxKey string

const (
	ContextUserID ctxKey = "user_id"
	ContextOrgID  ctxKey = "org_id"
)

// RequireAuth is an HTTP middleware that validates Bearer JWT and injects user/org into context.
func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authz := r.Header.Get("Authorization")
		if authz == "" {
			http.Error(w, "missing authorization", http.StatusUnauthorized)
			return
		}
		parts := strings.SplitN(authz, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "invalid authorization header", http.StatusUnauthorized)
			return
		}
		token := parts[1]
		claims, err := ParseToken(token)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		sub, _ := claims["sub"].(string)
		org, _ := claims["org"].(string)
		ctx := context.WithValue(r.Context(), ContextUserID, sub)
		ctx = context.WithValue(ctx, ContextOrgID, org)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
