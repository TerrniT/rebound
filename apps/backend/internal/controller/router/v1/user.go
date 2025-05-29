package v1

import (
	"strconv"

	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/internal/usecase"
	"github.com/terrnit/rebound/backend/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body entity.User true "User information"
// @Success 201 {object} entity.User
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/users [post]
func (h *userHandler) create(c *fiber.Ctx) error {
	var body struct {
		User     entity.User `json:"user"`
		Password string      `json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Error: "Invalid request body"})
	}

	createdUser, err := h.userUC.CreateUser(c.Context(), &body.User, body.Password)
	if err != nil {
		switch err {
		case usecase.ErrUsernameTaken:
			return c.Status(fiber.StatusConflict).JSON(ErrorResponse{Error: "Username already taken"})
		case usecase.ErrEmailTaken:
			return c.Status(fiber.StatusConflict).JSON(ErrorResponse{Error: "Email already taken"})
		case usecase.ErrInvalidInput:
			return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Error: "Invalid input"})
		default:
			h.logger.Error("Failed to create user", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Error: "Failed to create user"})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(createdUser)
}

// @Summary Get a user by ID
// @Description Get a user by their ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} entity.User
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/users/{id} [get]
func (h *userHandler) getByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.userUC.GetUser(c.Context(), id)
	if err != nil {
		switch err {
		case usecase.ErrUserNotFound:
			return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{Error: "User not found"})
		default:
			h.logger.Error("Failed to get user", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Error: "Failed to get user"})
		}
	}

	return c.JSON(user)
}

// @Summary List users
// @Description Get a paginated list of users
// @Tags users
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param size query int false "Page size (default: 10)"
// @Success 200 {object} PaginatedResponse{data=[]entity.User}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/users [get]
func (h *userHandler) list(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	size, _ := strconv.Atoi(c.Query("size", "10"))

	users, total, err := h.userUC.ListUsers(c.Context(), nil, page, size)
	if err != nil {
		h.logger.Error("Failed to list users", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Error: "Failed to list users"})
	}

	return c.JSON(PaginatedResponse{
		Data:  users,
		Total: total,
		Page:  page,
		Size:  size,
	})
}

// @Summary Update a user
// @Description Update a user's information
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body entity.User true "Updated user information"
// @Success 200 {object} entity.User
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/users/{id} [put]
func (h *userHandler) update(c *fiber.Ctx) error {
	id := c.Params("id")

	var user entity.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Error: "Invalid request body"})
	}

	user.ID = id
	err := h.userUC.UpdateUser(c.Context(), &user)
	if err != nil {
		switch err {
		case usecase.ErrUserNotFound:
			return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{Error: "User not found"})
		case usecase.ErrUsernameTaken:
			return c.Status(fiber.StatusConflict).JSON(ErrorResponse{Error: "Username already taken"})
		case usecase.ErrEmailTaken:
			return c.Status(fiber.StatusConflict).JSON(ErrorResponse{Error: "Email already taken"})
		case usecase.ErrInvalidInput:
			return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Error: "Invalid input"})
		default:
			h.logger.Error("Failed to update user", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Error: "Failed to update user"})
		}
	}

	return c.JSON(user)
}

// @Summary Delete a user
// @Description Delete a user by their ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/users/{id} [delete]
func (h *userHandler) delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.userUC.DeleteUser(c.Context(), id)
	if err != nil {
		switch err {
		case usecase.ErrUserNotFound:
			return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{Error: "User not found"})
		default:
			h.logger.Error("Failed to delete user", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Error: "Failed to delete user"})
		}
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary Update user password
// @Description Update a user's password
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param password body map[string]string true "New password"
// @Success 200 {object} entity.User
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/users/{id}/password [put]
func (h *userHandler) updatePassword(c *fiber.Ctx) error {
	id := c.Params("id")

	var body struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Error: "Invalid request body"})
	}

	err := h.userUC.UpdatePassword(c.Context(), id, body.CurrentPassword, body.NewPassword)
	if err != nil {
		switch err {
		case usecase.ErrUserNotFound:
			return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{Error: "User not found"})
		case usecase.ErrInvalidPassword:
			return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Error: "Invalid password"})
		default:
			h.logger.Error("Failed to update password", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Error: "Failed to update password"})
		}
	}

	return c.SendStatus(fiber.StatusOK)
}

// @Summary Update email verification status
// @Description Update a user's email verification status
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param verified body map[string]bool true "Verification status"
// @Success 200 {object} entity.User
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/users/{id}/verify-email [put]
func (h *userHandler) updateEmailVerification(c *fiber.Ctx) error {
	id := c.Params("id")

	var body struct {
		Verified bool `json:"verified"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Error: "Invalid request body"})
	}

	err := h.userUC.UpdateEmailVerification(c.Context(), id, body.Verified)
	if err != nil {
		switch err {
		case usecase.ErrUserNotFound:
			return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{Error: "User not found"})
		default:
			h.logger.Error("Failed to update email verification", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Error: "Failed to update email verification"})
		}
	}

	return c.SendStatus(fiber.StatusOK)
}

type userHandler struct {
	userUC *usecase.UserUseCase
	logger logger.Interface
}

// NewUserRoutes creates a new user routes handler
func NewUserRoutes(router fiber.Router, userUC *usecase.UserUseCase, l logger.Interface) {
	handler := &userHandler{
		userUC: userUC,
		logger: l,
	}

	users := router.Group("/users")
	users.Post("/", handler.create)
	users.Get("/:id", handler.getByID)
	users.Get("/", handler.list)
	users.Put("/:id", handler.update)
	users.Delete("/:id", handler.delete)
	users.Put("/:id/password", handler.updatePassword)
	users.Put("/:id/verify-email", handler.updateEmailVerification)
}
