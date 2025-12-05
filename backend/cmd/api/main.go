package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/Lenorthon/PORTFOLIO/backend/internal/auth"
	// import autres modules ici (workflows, billing, etc.)
)

func main() {
	// Crée le serveur Fiber
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
	})

	// Middleware logger
	app.Use(logger.New())

	// Routes Auth publiques
	auth.RegisterAuthRoutes(app)

	// Groupe de routes protégées
	api := app.Group("/api")
	api.Get("/users/me", auth.JWTMiddleware(), auth.MeHandler)

	// Routes test ou modules supplémentaires
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "API is running!"})
	})

	// Récupération du port depuis .env ou fallback
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
