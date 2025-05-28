package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/pkg/logger"
)

type authRoutes struct {
	uc entity.AuthUseCase
	l  logger.Interface
}

func NewAuthRoutes(router fiber.Router, uc entity.AuthUseCase, l logger.Interface) {
	r := &authRoutes{
		uc: uc,
		l:  l,
	}

	group := router.Group("/auth")
	{
		group.Post("/token", r.createToken)
		group.Delete("/token/:token", r.deleteToken)
		group.Get("/validate/:token", r.validateToken)
	}
}

// @Summary Create a new token
// @Description Create a new authentication token for a user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body entity.User true "User object"
// @Success 201 {object} entity.Token
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/token [post]
func (r *authRoutes) createToken(c *fiber.Ctx) error {
	user := new(entity.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	token, err := r.uc.GenerateToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(token)
}

// @Summary Validate token
// @Description Validate a token and get the associated user
// @Tags auth
// @Accept json
// @Produce json
// @Param token path string true "Token string"
// @Success 200 {object} entity.User
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/validate/{token} [get]
func (r *authRoutes) validateToken(c *fiber.Ctx) error {
	token := c.Params("token")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Token is required",
		})
	}

	user, err := r.uc.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	return c.JSON(user)
}

// @Summary Revoke token
// @Description Revoke a specific token
// @Tags auth
// @Accept json
// @Produce json
// @Param token path string true "Token string"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/token/{token} [delete]
func (r *authRoutes) deleteToken(c *fiber.Ctx) error {
	token := c.Params("token")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Token is required",
		})
	}

	if err := r.uc.RevokeToken(token); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to revoke token",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
