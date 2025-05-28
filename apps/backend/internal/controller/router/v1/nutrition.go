package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/pkg/logger"
)

type nutritionRoutes struct {
	uc entity.NutritionUseCase
	l  logger.Interface
}

func NewNutritionRoutes(router fiber.Router, uc entity.NutritionUseCase, l logger.Interface) {
	r := &nutritionRoutes{
		uc: uc,
		l:  l,
	}

	group := router.Group("/nutrition")
	{
		group.Post("/", r.create)
		group.Get("/:id", r.getByID)
		group.Put("/:id", r.update)
		group.Delete("/:id", r.delete)
		group.Get("/", r.list)
	}
}

// @Summary Create a new nutrition entry
// @Description Create a new nutrition entry with the provided information
// @Tags nutrition
// @Accept json
// @Produce json
// @Param nutrition body entity.Nutrition true "Nutrition object"
// @Success 201 {object} entity.Nutrition
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /nutrition [post]
func (r *nutritionRoutes) create(c *fiber.Ctx) error {
	nutrition := new(entity.Nutrition)
	if err := c.BodyParser(nutrition); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := r.uc.CreateNutrition(nutrition); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create nutrition entry",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(nutrition)
}

// @Summary Get nutrition entry by ID
// @Description Get nutrition information by its ID
// @Tags nutrition
// @Accept json
// @Produce json
// @Param id path int true "Nutrition ID"
// @Success 200 {object} entity.Nutrition
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /nutrition/{id} [get]
func (r *nutritionRoutes) getByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid nutrition ID",
		})
	}

	nutrition, err := r.uc.GetNutrition(int64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Nutrition entry not found",
		})
	}

	return c.JSON(nutrition)
}

// @Summary Update nutrition entry
// @Description Update nutrition information by its ID
// @Tags nutrition
// @Accept json
// @Produce json
// @Param id path int true "Nutrition ID"
// @Param nutrition body entity.Nutrition true "Nutrition object"
// @Success 200 {object} entity.Nutrition
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /nutrition/{id} [put]
func (r *nutritionRoutes) update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid nutrition ID",
		})
	}

	nutrition := new(entity.Nutrition)
	if err := c.BodyParser(nutrition); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	nutrition.ID = int64(id)
	if err := r.uc.UpdateNutrition(nutrition); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update nutrition entry",
		})
	}

	return c.JSON(nutrition)
}

// @Summary Delete nutrition entry
// @Description Delete a nutrition entry by its ID
// @Tags nutrition
// @Accept json
// @Produce json
// @Param id path int true "Nutrition ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /nutrition/{id} [delete]
func (r *nutritionRoutes) delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid nutrition ID",
		})
	}

	if err := r.uc.DeleteNutrition(int64(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete nutrition entry",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary List nutrition entries
// @Description Get a list of nutrition entries with pagination
// @Tags nutrition
// @Accept json
// @Produce json
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Success 200 {array} entity.Nutrition
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /nutrition [get]
func (r *nutritionRoutes) list(c *fiber.Ctx) error {
	offset := c.QueryInt("offset", 0)
	limit := c.QueryInt("limit", 10)

	nutritions, err := r.uc.ListNutritions(offset, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to list nutrition entries",
		})
	}

	return c.JSON(nutritions)
}
