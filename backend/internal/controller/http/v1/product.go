package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/pkg/logger"
)

type productRoutes struct {
	uc entity.ProductUseCase
	l  logger.Interface
}

func NewProductRoutes(router fiber.Router, uc entity.ProductUseCase, l logger.Interface) {
	r := &productRoutes{
		uc: uc,
		l:  l,
	}

	group := router.Group("/products")
	{
		group.Post("/", r.create)
		group.Get("/:id", r.getByID)
		group.Put("/:id", r.update)
		group.Delete("/:id", r.delete)
		group.Get("/", r.list)
		group.Get("/search", r.search)
		group.Get("/barcode/:barcode", r.getByBarcode)
	}
}

func (r *productRoutes) create(c *fiber.Ctx) error {
	product := new(entity.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := r.uc.CreateProduct(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

func (r *productRoutes) getByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	product, err := r.uc.GetProduct(int64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	return c.JSON(product)
}

func (r *productRoutes) update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	product := new(entity.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	product.ID = int64(id)
	if err := r.uc.UpdateProduct(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update product",
		})
	}

	return c.JSON(product)
}

func (r *productRoutes) delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	if err := r.uc.DeleteProduct(int64(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete product",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (r *productRoutes) list(c *fiber.Ctx) error {
	offset := c.QueryInt("offset", 0)
	limit := c.QueryInt("limit", 10)

	products, err := r.uc.ListProducts(offset, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to list products",
		})
	}

	return c.JSON(products)
}

func (r *productRoutes) search(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Search query is required",
		})
	}

	products, err := r.uc.SearchProducts(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to search products",
		})
	}

	return c.JSON(products)
}

func (r *productRoutes) getByBarcode(c *fiber.Ctx) error {
	barcode := c.Params("barcode")
	if barcode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Barcode is required",
		})
	}

	product, err := r.uc.GetProductByBarcode(barcode)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	return c.JSON(product)
}
