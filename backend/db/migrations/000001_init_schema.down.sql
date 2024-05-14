-- Drop all triggers
DROP TRIGGER IF EXISTS on_update_updated_at ON order_details;

DROP TRIGGER IF EXISTS on_update_updated_at ON orders;

DROP TRIGGER IF EXISTS on_update_updated_at ON users;

DROP TRIGGER IF EXISTS on_update_updated_at ON product_translations;

DROP TRIGGER IF EXISTS on_update_updated_at ON products;

DROP TRIGGER IF EXISTS on_update_updated_at ON languages;

DROP TRIGGER IF EXISTS on_update_last_login ON users;

-- Drop the function
DROP FUNCTION IF EXISTS update_updated_at;

DROP FUNCTION IF EXISTS update_last_login;

-- Drop all tables
DROP TABLE IF EXISTS order_details;

DROP TABLE IF EXISTS orders;

DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS product_translations;

DROP TABLE IF EXISTS products;

DROP TABLE IF EXISTS languages;

-- Drop all enum types
DROP TYPE IF EXISTS order_status;

DROP TYPE IF EXISTS language_code;

DROP TYPE IF EXISTS product_status;

DROP TYPE IF EXISTS gender;

-- Drop the UUID extension
DROP EXTENSION IF EXISTS "uuid-ossp";
