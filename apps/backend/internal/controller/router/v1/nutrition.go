package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/terrnit/rebound/backend/internal/usecase"
	"github.com/terrnit/rebound/backend/pkg/logger"
)

type NutritionRoutes struct {
	nutritionUC *usecase.NutritionUseCase
	log         logger.Interface
}

func NewNutritionRoutes(handler fiber.Router, uc *usecase.NutritionUseCase, l logger.Interface) {
	r := &NutritionRoutes{
		nutritionUC: uc,
		log:         l,
	}

	h := handler.Group("/nutrition")
	{
		// Nutrition goals routes
		h.Post("/goals", r.createNutritionGoals)
		h.Get("/goals/:id", r.getNutritionGoals)
		h.Put("/goals/:id", r.updateNutritionGoals)
		h.Delete("/goals/:id", r.deleteNutritionGoals)
		h.Get("/goals/user/:userID/active", r.getActiveNutritionGoals)
		h.Get("/goals/user/:userID/history", r.getNutritionGoalsHistory)

		// Biometrics routes
		h.Post("/biometrics", r.createBiometrics)
		h.Get("/biometrics/:id", r.getBiometrics)
		h.Put("/biometrics/:id", r.updateBiometrics)
		h.Delete("/biometrics/:id", r.deleteBiometrics)
		h.Get("/biometrics/user/:userID/history", r.getUserBiometricsHistory)
		h.Get("/biometrics/user/:userID/latest", r.getLatestBiometrics)
	}
}

// @Summary Create nutrition goals
// @Description Create new nutrition goals for a user
// @Tags nutrition
// @Accept json
// @Produce json
// @Param goals body entity.UserNutritionGoal true "Nutrition goals object"
// @Success 201 {object} entity.UserNutritionGoal
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /nutrition/goals [post]
func (r *NutritionRoutes) createNutritionGoals(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

// @Summary Get nutrition goals
// @Description Get nutrition goals by ID
// @Tags nutrition
// @Accept json
// @Produce json
// @Param id path string true "Nutrition goals ID"
// @Success 200 {object} entity.UserNutritionGoal
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /nutrition/goals/{id} [get]
func (r *NutritionRoutes) getNutritionGoals(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

// @Summary Update nutrition goals
// @Description Update existing nutrition goals
// @Tags nutrition
// @Accept json
// @Produce json
// @Param id path string true "Nutrition goals ID"
// @Param goals body entity.UserNutritionGoal true "Nutrition goals object"
// @Success 200 {object} entity.UserNutritionGoal
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /nutrition/goals/{id} [put]
func (r *NutritionRoutes) updateNutritionGoals(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

// @Summary Delete nutrition goals
// @Description Delete nutrition goals by ID
// @Tags nutrition
// @Accept json
// @Produce json
// @Param id path string true "Nutrition goals ID"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /nutrition/goals/{id} [delete]
func (r *NutritionRoutes) deleteNutritionGoals(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

// @Summary Get active nutrition goals
// @Description Get the active nutrition goals for a user
// @Tags nutrition
// @Accept json
// @Produce json
// @Param userID path string true "User ID"
// @Success 200 {object} entity.UserNutritionGoal
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /nutrition/goals/user/{userID}/active [get]
func (r *NutritionRoutes) getActiveNutritionGoals(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

// @Summary Get nutrition goals history
// @Description Get nutrition goals history for a user
// @Tags nutrition
// @Accept json
// @Produce json
// @Param userID path string true "User ID"
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Success 200 {array} entity.UserNutritionGoal
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /nutrition/goals/user/{userID}/history [get]
func (r *NutritionRoutes) getNutritionGoalsHistory(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

// @Summary Create biometrics
// @Description Create new biometrics for a user
// @Tags nutrition
// @Accept json
// @Produce json
// @Param biometrics body entity.UserBiometric true "Biometrics object"
// @Success 201 {object} entity.UserBiometric
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /nutrition/biometrics [post]
func (r *NutritionRoutes) createBiometrics(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

// @Summary Get biometrics
// @Description Get biometrics by ID
// @Tags nutrition
// @Accept json
// @Produce json
// @Param id path string true "Biometrics ID"
// @Success 200 {object} entity.UserBiometric
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /nutrition/biometrics/{id} [get]
func (r *NutritionRoutes) getBiometrics(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

// @Summary Update biometrics
// @Description Update existing biometrics
// @Tags nutrition
// @Accept json
// @Produce json
// @Param id path string true "Biometrics ID"
// @Param biometrics body entity.UserBiometric true "Biometrics object"
// @Success 200 {object} entity.UserBiometric
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /nutrition/biometrics/{id} [put]
func (r *NutritionRoutes) updateBiometrics(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

// @Summary Delete biometrics
// @Description Delete biometrics by ID
// @Tags nutrition
// @Accept json
// @Produce json
// @Param id path string true "Biometrics ID"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /nutrition/biometrics/{id} [delete]
func (r *NutritionRoutes) deleteBiometrics(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

// @Summary Get user biometrics history
// @Description Get biometrics history for a user
// @Tags nutrition
// @Accept json
// @Produce json
// @Param userID path string true "User ID"
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Success 200 {array} entity.UserBiometric
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /nutrition/biometrics/user/{userID}/history [get]
func (r *NutritionRoutes) getUserBiometricsHistory(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

// @Summary Get latest biometrics
// @Description Get the most recent biometrics for a user
// @Tags nutrition
// @Accept json
// @Produce json
// @Param userID path string true "User ID"
// @Success 200 {object} entity.UserBiometric
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /nutrition/biometrics/user/{userID}/latest [get]
func (r *NutritionRoutes) getLatestBiometrics(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}
