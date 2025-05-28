package postgres

import (
	"context"
	"time"

	"github.com/terrnit/rebound/backend/internal/entity"
	"github.com/terrnit/rebound/backend/pkg/postgres"
)

type ProductRepository struct {
	db *postgres.Postgres
}

func NewProductRepository(db *postgres.Postgres) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *entity.Product) error {
	query := `
		INSERT INTO products (name, brand, nutrition_id, barcode, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`

	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now

	return r.db.Pool.QueryRow(
		context.Background(),
		query,
		product.Name,
		product.Brand,
		product.Nutrition.ID,
		product.Barcode,
		product.Description,
		product.CreatedAt,
		product.UpdatedAt,
	).Scan(&product.ID)
}

func (r *ProductRepository) GetByID(id int64) (*entity.Product, error) {
	query := `
		SELECT p.id, p.name, p.brand, p.barcode, p.description, p.created_at, p.updated_at,
			   n.id, n.calories, n.proteins, n.carbs, n.fats, n.serving_size, n.unit, n.created_at, n.updated_at
		FROM products p
		JOIN nutrition n ON p.nutrition_id = n.id
		WHERE p.id = $1`

	product := &entity.Product{}
	err := r.db.Pool.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Brand,
		&product.Barcode,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.Nutrition.ID,
		&product.Nutrition.Calories,
		&product.Nutrition.Proteins,
		&product.Nutrition.Carbs,
		&product.Nutrition.Fats,
		&product.Nutrition.ServingSize,
		&product.Nutrition.Unit,
		&product.Nutrition.CreatedAt,
		&product.Nutrition.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *ProductRepository) Update(product *entity.Product) error {
	query := `
		UPDATE products
		SET name = $1, brand = $2, nutrition_id = $3, barcode = $4, description = $5, updated_at = $6
		WHERE id = $7`

	product.UpdatedAt = time.Now()
	_, err := r.db.Pool.Exec(
		context.Background(),
		query,
		product.Name,
		product.Brand,
		product.Nutrition.ID,
		product.Barcode,
		product.Description,
		product.UpdatedAt,
		product.ID,
	)
	return err
}

func (r *ProductRepository) Delete(id int64) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.Pool.Exec(context.Background(), query, id)
	return err
}

func (r *ProductRepository) List(offset, limit int) ([]*entity.Product, error) {
	query := `
		SELECT p.id, p.name, p.brand, p.barcode, p.description, p.created_at, p.updated_at,
			   n.id, n.calories, n.proteins, n.carbs, n.fats, n.serving_size, n.unit, n.created_at, n.updated_at
		FROM products p
		JOIN nutrition n ON p.nutrition_id = n.id
		ORDER BY p.id
		LIMIT $1 OFFSET $2`

	rows, err := r.db.Pool.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		product := &entity.Product{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Brand,
			&product.Barcode,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.Nutrition.ID,
			&product.Nutrition.Calories,
			&product.Nutrition.Proteins,
			&product.Nutrition.Carbs,
			&product.Nutrition.Fats,
			&product.Nutrition.ServingSize,
			&product.Nutrition.Unit,
			&product.Nutrition.CreatedAt,
			&product.Nutrition.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) Search(query string) ([]*entity.Product, error) {
	sqlQuery := `
		SELECT p.id, p.name, p.brand, p.barcode, p.description, p.created_at, p.updated_at,
			   n.id, n.calories, n.proteins, n.carbs, n.fats, n.serving_size, n.unit, n.created_at, n.updated_at
		FROM products p
		JOIN nutrition n ON p.nutrition_id = n.id
		WHERE p.name ILIKE $1 OR p.brand ILIKE $1 OR p.description ILIKE $1
		ORDER BY p.name`

	searchPattern := "%" + query + "%"
	rows, err := r.db.Pool.Query(context.Background(), sqlQuery, searchPattern)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		product := &entity.Product{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Brand,
			&product.Barcode,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.Nutrition.ID,
			&product.Nutrition.Calories,
			&product.Nutrition.Proteins,
			&product.Nutrition.Carbs,
			&product.Nutrition.Fats,
			&product.Nutrition.ServingSize,
			&product.Nutrition.Unit,
			&product.Nutrition.CreatedAt,
			&product.Nutrition.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) GetByBarcode(barcode string) (*entity.Product, error) {
	query := `
		SELECT p.id, p.name, p.brand, p.barcode, p.description, p.created_at, p.updated_at,
			   n.id, n.calories, n.proteins, n.carbs, n.fats, n.serving_size, n.unit, n.created_at, n.updated_at
		FROM products p
		JOIN nutrition n ON p.nutrition_id = n.id
		WHERE p.barcode = $1`

	product := &entity.Product{}
	err := r.db.Pool.QueryRow(
		context.Background(),
		query,
		barcode,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Brand,
		&product.Barcode,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.Nutrition.ID,
		&product.Nutrition.Calories,
		&product.Nutrition.Proteins,
		&product.Nutrition.Carbs,
		&product.Nutrition.Fats,
		&product.Nutrition.ServingSize,
		&product.Nutrition.Unit,
		&product.Nutrition.CreatedAt,
		&product.Nutrition.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}
