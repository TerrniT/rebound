package v1

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/internal/usecase"
	"github.com/terrnit/rebound/backend/pkg/logger"
)

type foodItemHandler struct {
	usecase *usecase.FoodItemUseCase
	logger  logger.Interface
}

// @Summary Create a new food item
// @Description Create a new food item with the provided details
// @Tags food-items
// @Accept json
// @Produce json
// @Param foodItem body entity.FoodItem true "Food item details"
// @Success 201 {object} entity.FoodItem
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /food-items [post]
func (h *foodItemHandler) create(c *fiber.Ctx) error {
	var foodItem entity.FoodItem
	if err := c.BodyParser(&foodItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Error: "Invalid request body"})
	}

	created, err := h.usecase.CreateFoodItem(c.Context(), &foodItem)
	if err != nil {
		h.logger.Error("Failed to create food item", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Error: "Failed to create food item"})
	}

	return c.Status(fiber.StatusCreated).JSON(created)
}

// @Summary Get a food item by ID
// @Description Get a food item by its ID
// @Tags food-items
// @Produce json
// @Param id path string true "Food item ID"
// @Success 200 {object} entity.FoodItem
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /food-items/{id} [get]
func (h *foodItemHandler) getByID(c *fiber.Ctx) error {
	id := c.Params("id")
	foodItem, err := h.usecase.GetFoodItem(c.Context(), id)
	if err != nil {
		h.logger.Error("Failed to get food item", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Error: "Failed to get food item"})
	}
	if foodItem == nil {
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{Error: "Food item not found"})
	}

	return c.JSON(foodItem)
}

// @Summary List food items
// @Description Get a paginated list of food items
// @Tags food-items
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param page_size query int false "Items per page (default: 10)"
// @Success 200 {object} PaginatedResponse{data=[]entity.FoodItem}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /food-items [get]
func (h *foodItemHandler) list(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	items, total, err := h.usecase.ListFoodItems(c.Context(), nil, page, pageSize)
	if err != nil {
		h.logger.Error("Failed to list food items", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Error: "Failed to list food items"})
	}

	return c.JSON(PaginatedResponse{
		Data:  items,
		Total: total,
		Page:  page,
		Size:  pageSize,
	})
}

// @Summary Search food items
// @Description Search food items by name or brand
// @Tags food-items
// @Produce json
// @Param query query string true "Search query"
// @Param page query int false "Page number (default: 1)"
// @Param page_size query int false "Items per page (default: 10)"
// @Success 200 {object} PaginatedResponse{data=[]entity.FoodItem}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /food-items/search [get]
func (h *foodItemHandler) search(c *fiber.Ctx) error {
	query := c.Query("query")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Error: "Search query is required"})
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	items, total, err := h.usecase.SearchFoodItems(c.Context(), query, page, pageSize)
	if err != nil {
		h.logger.Error("Failed to search food items", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Error: "Failed to search food items"})
	}

	return c.JSON(PaginatedResponse{
		Data:  items,
		Total: total,
		Page:  page,
		Size:  pageSize,
	})
}

// @Summary Update a food item
// @Description Update an existing food item
// @Tags food-items
// @Accept json
// @Produce json
// @Param id path string true "Food item ID"
// @Param foodItem body entity.FoodItem true "Updated food item details"
// @Success 200 {object} entity.FoodItem
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /food-items/{id} [put]
func (h *foodItemHandler) update(c *fiber.Ctx) error {
	id := c.Params("id")
	var foodItem entity.FoodItem
	if err := c.BodyParser(&foodItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Error: "Invalid request body"})
	}

	foodItem.ID = id
	if err := h.usecase.UpdateFoodItem(c.Context(), &foodItem); err != nil {
		h.logger.Error("Failed to update food item", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Error: "Failed to update food item"})
	}

	return c.JSON(foodItem)
}

// @Summary Delete a food item
// @Description Delete a food item by its ID
// @Tags food-items
// @Param id path string true "Food item ID"
// @Success 204 "No Content"
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /food-items/{id} [delete]
func (h *foodItemHandler) delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.usecase.DeleteFoodItem(c.Context(), id); err != nil {
		h.logger.Error("Failed to delete food item", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Error: "Failed to delete food item"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// NewFoodItemRoutes creates routes for food item operations
func NewFoodItemRoutes(router fiber.Router, uc *usecase.FoodItemUseCase, l logger.Interface) {
	handler := &foodItemHandler{
		usecase: uc,
		logger:  l,
	}

	foodItems := router.Group("/food-items")
	{
		foodItems.Post("/", handler.create)
		foodItems.Get("/:id", handler.getByID)
		foodItems.Get("/", handler.list)
		foodItems.Get("/search", handler.search)
		foodItems.Put("/:id", handler.update)
		foodItems.Delete("/:id", handler.delete)
	}
}
