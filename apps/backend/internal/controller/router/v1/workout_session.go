package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/terrnit/rebound/backend/internal/usecase"
	"github.com/terrnit/rebound/backend/pkg/logger"
)

type WorkoutSessionRoutes struct {
	workoutSessionUC *usecase.WorkoutSessionUseCase
	log              logger.Interface
}

func NewWorkoutSessionRoutes(handler fiber.Router, uc *usecase.WorkoutSessionUseCase, l logger.Interface) {
	r := &WorkoutSessionRoutes{
		workoutSessionUC: uc,
		log:              l,
	}

	h := handler.Group("/workout-sessions")
	{
		h.Post("/", r.createWorkoutSession)
		h.Get("/:id", r.getWorkoutSession)
		h.Put("/:id", r.updateWorkoutSession)
		h.Delete("/:id", r.deleteWorkoutSession)
		h.Get("/user/:userID", r.getUserWorkoutSessions)
		h.Post("/:id/exercises", r.addExercise)
		h.Get("/:id/exercises", r.getExercises)
		h.Put("/exercises/:id", r.updateExercise)
		h.Delete("/exercises/:id", r.deleteExercise)
	}
}

// @Summary Create a new workout session
// @Description Create a new workout session for a user
// @Tags workout-sessions
// @Accept json
// @Produce json
// @Param session body entity.UserWorkoutSession true "Workout session object"
// @Success 201 {object} entity.UserWorkoutSession
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /workout-sessions [post]
func (r *WorkoutSessionRoutes) createWorkoutSession(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

// @Summary Get a workout session by ID
// @Description Get a workout session by its ID
// @Tags workout-sessions
// @Accept json
// @Produce json
// @Param id path string true "Workout session ID"
// @Success 200 {object} entity.UserWorkoutSession
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /workout-sessions/{id} [get]
func (r *WorkoutSessionRoutes) getWorkoutSession(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

// @Summary Update a workout session
// @Description Update an existing workout session
// @Tags workout-sessions
// @Accept json
// @Produce json
// @Param id path string true "Workout session ID"
// @Param session body entity.UserWorkoutSession true "Workout session object"
// @Success 200 {object} entity.UserWorkoutSession
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /workout-sessions/{id} [put]
func (r *WorkoutSessionRoutes) updateWorkoutSession(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

// @Summary Delete a workout session
// @Description Delete a workout session by its ID
// @Tags workout-sessions
// @Accept json
// @Produce json
// @Param id path string true "Workout session ID"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /workout-sessions/{id} [delete]
func (r *WorkoutSessionRoutes) deleteWorkoutSession(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

// @Summary Get user workout sessions
// @Description Get all workout sessions for a specific user
// @Tags workout-sessions
// @Accept json
// @Produce json
// @Param userID path string true "User ID"
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Success 200 {array} entity.UserWorkoutSession
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /workout-sessions/user/{userID} [get]
func (r *WorkoutSessionRoutes) getUserWorkoutSessions(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

// @Summary Add exercise to workout session
// @Description Add a new exercise to a workout session
// @Tags workout-sessions
// @Accept json
// @Produce json
// @Param id path string true "Workout session ID"
// @Param exercise body entity.UserWorkoutSession true "Exercise object"
// @Success 201 {object} entity.UserWorkoutSession
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /workout-sessions/{id}/exercises [post]
func (r *WorkoutSessionRoutes) addExercise(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

// @Summary Get workout session exercises
// @Description Get all exercises for a workout session
// @Tags workout-sessions
// @Accept json
// @Produce json
// @Param id path string true "Workout session ID"
// @Success 200 {array} entity.UserWorkoutSession
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /workout-sessions/{id}/exercises [get]
func (r *WorkoutSessionRoutes) getExercises(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

// @Summary Update workout exercise
// @Description Update an existing exercise in a workout session
// @Tags workout-sessions
// @Accept json
// @Produce json
// @Param id path string true "Exercise ID"
// @Param exercise body entity.UserWorkoutSession true "Exercise object"
// @Success 200 {object} entity.UserWorkoutSession
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /workout-sessions/exercises/{id} [put]
func (r *WorkoutSessionRoutes) updateExercise(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}

// @Summary Delete workout exercise
// @Description Delete an exercise from a workout session
// @Tags workout-sessions
// @Accept json
// @Produce json
// @Param id path string true "Exercise ID"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /workout-sessions/exercises/{id} [delete]
func (r *WorkoutSessionRoutes) deleteExercise(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(fiber.StatusNotImplemented)
}
