DROP TABLE IF EXISTS schema_migrations;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP

);

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    nutrition JSONB NOT NULL,
    barcode VARCHAR(255) UNIQUE,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Trigger function for auto-updating updated_at
-- CREATE OR REPLACE FUNCTION update_products_updated_at_column()
-- RETURNS TRIGGER AS $$
-- BEGIN
--     NEW.updated_at = NOW();
--     RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;

-- Apply trigger to products table
-- CREATE TRIGGER update_products_updated_at
-- BEFORE UPDATE ON products
-- FOR EACH ROW
-- EXECUTE FUNCTION update_products_updated_at_column();


-- -- Trigger function for auto-updating updated_at
-- CREATE OR REPLACE FUNCTION update_updated_at_column()
-- RETURNS TRIGGER AS $$
-- BEGIN
--     NEW.updated_at = NOW();
--     RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;

-- -- Apply trigger to users table
-- CREATE TRIGGER update_users_updated_at
-- BEFORE UPDATE ON users
-- FOR EACH ROW
-- EXECUTE FUNCTION update_updated_at_column();

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_created_at ON users (created_at);
CREATE INDEX IF NOT EXISTS idx_users_updated_at ON users (updated_at);
CREATE INDEX IF NOT EXISTS idx_products_created_at ON products (created_at);
CREATE INDEX IF NOT EXISTS idx_products_updated_at ON products (updated_at);
CREATE INDEX IF NOT EXISTS idx_products_brand ON products (brand);