-- Drop indexes
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_created_at;
DROP INDEX IF EXISTS idx_users_updated_at;
DROP INDEX IF EXISTS idx_products_created_at;
DROP INDEX IF EXISTS idx_products_updated_at;
DROp INDEX IF EXISTS idx_products_brand;

-- Drop tables
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS users; 