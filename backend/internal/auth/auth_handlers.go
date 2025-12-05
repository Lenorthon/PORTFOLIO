package auth

import "github.com/gofiber/fiber/v2"

func RegisterHandler(c *fiber.Ctx) error {
	req := new(AuthRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := RegisterUser(req.Email, req.Password, req.Name, "org_default")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	token, _ := GenerateJWT(user.ID, user.Role)
	return c.JSON(AuthResponse{Token: token})
}

func LoginHandler(c *fiber.Ctx) error {
	req := new(AuthRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := LoginUser(req.Email, req.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	token, _ := GenerateJWT(user.ID, user.Role)
	return c.JSON(AuthResponse{Token: token})
}

func MeHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*User)
	return c.JSON(user)
}

func InviteHandler(c *fiber.Ctx) error {
	type Req struct{ Email string }
	req := new(Req)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	invite, _ := CreateInvite("org_default", req.Email)
	return c.JSON(invite)
}

func AcceptInviteHandler(c *fiber.Ctx) error {
	type Req struct {
		Token    string `json:"token"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	req := new(Req)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := AcceptInvite(req.Token, req.Name, req.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	token, _ := GenerateJWT(user.ID, user.Role)
	return c.JSON(AuthResponse{Token: token})
}
