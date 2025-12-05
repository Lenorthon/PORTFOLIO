package auth

import "github.com/gofiber/fiber/v2"

func RegisterAuthRoutes(app *fiber.App) {
	api := app.Group("/api/auth")
	api.Post("/register", RegisterHandler)
	api.Post("/login", LoginHandler)
	api.Get("/me", MeHandler) // ajouter middleware JWT dans vrai projet
	api.Post("/invite", InviteHandler)
	api.Post("/invite/accept", AcceptInviteHandler)
	// TODO: /refresh token à ajouter si refresh token implémenté
}
