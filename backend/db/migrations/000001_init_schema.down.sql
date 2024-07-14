-- Drop triggers
DROP TRIGGER IF EXISTS trg_update_last_login ON users;
DROP TRIGGER IF EXISTS trg_update_products_timestamp ON products;
DROP TRIGGER IF EXISTS trg_update_product_translations_timestamp ON product_translations;
DROP TRIGGER IF EXISTS trg_update_users_timestamp ON users;
DROP TRIGGER IF EXISTS trg_update_orders_timestamp ON orders;
DROP TRIGGER IF EXISTS trg_update_order_details_timestamp ON order_details;

-- Drop functions
DROP FUNCTION IF EXISTS update_last_login_timestamp();
DROP FUNCTION IF EXISTS update_timestamp();

-- Drop tables
DROP TABLE IF EXISTS order_details;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS product_descriptions;
DROP TABLE IF EXISTS product_translations;
DROP TABLE IF EXISTS products;

-- Drop types
DROP TYPE IF EXISTS gender;
DROP TYPE IF EXISTS order_status;
DROP TYPE IF EXISTS language_code;
DROP TYPE IF EXISTS product_status;

-- Drop extension
DROP EXTENSION IF EXISTS "uuid-ossp";

