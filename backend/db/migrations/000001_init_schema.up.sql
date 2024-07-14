CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Enum Types
CREATE TYPE product_status AS ENUM (
    'in-stock',
    'out-of-stock',
    'discontinued'
);

CREATE TYPE language_code AS ENUM (
    'chn',
    'jp'
);

CREATE TYPE order_status AS ENUM (
    'pending',
    'processing',
    'shipped',
    'delivered',
    'cancelled'
);

CREATE TYPE gender AS ENUM (
    'male',
    'female',
    'not-specified',
    'not-disclosed'
);

-- Function to update updated_at
CREATE FUNCTION update_timestamp() RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Function to update last_login
CREATE FUNCTION update_last_login_timestamp() RETURNS TRIGGER AS $$
BEGIN
    NEW.last_login = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Products table
CREATE TABLE "products" (
    "pk" bigserial PRIMARY KEY,
    "id" uuid UNIQUE DEFAULT uuid_generate_v4(),
    "price" int NOT NULL DEFAULT 0,
    "discount" int NOT NULL DEFAULT 0,
    "is_hot" boolean NOT NULL DEFAULT FALSE,
    "imageURLs" text[] NOT NULL DEFAULT ARRAY[]::text[],
    "status" product_status NOT NULL DEFAULT 'in-stock',
    "quantity" int NOT NULL DEFAULT 0,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_products_status ON "products" ("status");
CREATE INDEX idx_products_is_hot ON "products" ("is_hot");

-- Product translations table
CREATE TABLE "product_translations" (
    "pk" bigserial PRIMARY KEY,
    "product_pk" bigint NOT NULL REFERENCES products (pk) ON DELETE CASCADE,
    "name" text NOT NULL DEFAULT '',
    "language" language_code NOT NULL DEFAULT 'chn',
    "category" text NOT NULL DEFAULT '',
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_product_translations_product_pk ON "product_translations" ("product_pk");
CREATE INDEX idx_product_translations_language ON "product_translations" ("language");

-- Product descriptions table
CREATE TABLE "product_descriptions" (
    "pk" bigserial PRIMARY KEY,
    "product_translation_pk" bigint NOT NULL REFERENCES product_translations (pk) ON DELETE CASCADE,
    "introduction" text NOT NULL DEFAULT '',
    "prize" text NOT NULL DEFAULT '',
    "item_info" text NOT NULL DEFAULT '',
    "recommendation" text NOT NULL DEFAULT ''
);

CREATE INDEX idx_product_descriptions_translation_pk ON "product_descriptions" ("product_translation_pk");

-- Users table
CREATE TABLE "users" (
    "pk" bigserial PRIMARY KEY,
    "id" uuid UNIQUE DEFAULT uuid_generate_v4(),
    "line_id" text UNIQUE,
    "birth_year" int,
    "gender" gender NOT NULL DEFAULT 'not-disclosed',
    "phone" text UNIQUE,
    "email" text UNIQUE NOT NULL DEFAULT '',
    "password" text NOT NULL DEFAULT '',
    "first_name" text NOT NULL DEFAULT '',
    "last_name" text NOT NULL DEFAULT '',
    "language" language_code NOT NULL DEFAULT 'chn',
    "address" text NOT NULL DEFAULT '',
    "longitude" float,
    "latitude" float,
    "last_login" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_line_id ON "users" ("line_id");
CREATE INDEX idx_users_email ON "users" ("email");
CREATE INDEX idx_users_phone ON "users" ("phone");

-- Orders table
CREATE TABLE "orders" (
    "pk" bigserial PRIMARY KEY,
    "id" uuid UNIQUE DEFAULT uuid_generate_v4(),
    "user_pk" bigint NOT NULL REFERENCES users (pk) ON DELETE CASCADE,
    "status" order_status NOT NULL DEFAULT 'pending',
    "is_paid" boolean NOT NULL DEFAULT FALSE,
    "total_price" int NOT NULL DEFAULT 0,
    "shipping_address" text NOT NULL,
    "email" text NOT NULL,
    "phone" text NOT NULL,
    "shipping_date" timestamptz,
    "delivered_date" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_orders_user_pk ON "orders" ("user_pk");

-- Order details table
CREATE TABLE "order_details" (
    "pk" bigserial PRIMARY KEY,
    "order_pk" bigint NOT NULL REFERENCES orders (pk) ON DELETE CASCADE,
    "product_pk" bigint NOT NULL REFERENCES products (pk) ON DELETE CASCADE,
    "quantity" int NOT NULL DEFAULT 0,
    "price" int NOT NULL DEFAULT 0,
    "discount" int NOT NULL DEFAULT 0,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_order_details_order_pk ON "order_details" ("order_pk");
CREATE INDEX idx_order_details_product_pk ON "order_details" ("product_pk");

-- Triggers to update timestamps
CREATE TRIGGER trg_update_last_login
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_last_login_timestamp();

CREATE TRIGGER trg_update_products_timestamp
    BEFORE UPDATE ON products
    FOR EACH ROW
    EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER trg_update_product_translations_timestamp
    BEFORE UPDATE ON product_translations
    FOR EACH ROW
    EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER trg_update_users_timestamp
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER trg_update_orders_timestamp
    BEFORE UPDATE ON orders
    FOR EACH ROW
    EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER trg_update_order_details_timestamp
    BEFORE UPDATE ON order_details
    FOR EACH ROW
    EXECUTE FUNCTION update_timestamp();

