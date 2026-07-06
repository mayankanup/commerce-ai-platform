PRAGMA foreign_keys = ON;

--------------------------------------------------------------------------------
-- Schema Version
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS schema_version (
    version         INTEGER PRIMARY KEY,
    applied_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--------------------------------------------------------------------------------
-- Categories
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS categories (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,

    name            TEXT NOT NULL UNIQUE,

    description     TEXT,

    is_active       INTEGER NOT NULL DEFAULT 1
                        CHECK(is_active IN (0,1)),

    created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--------------------------------------------------------------------------------
-- Brands
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS brands (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,

    name            TEXT NOT NULL UNIQUE,

    is_active       INTEGER NOT NULL DEFAULT 1
                        CHECK(is_active IN (0,1)),

    created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--------------------------------------------------------------------------------
-- Colors
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS colors (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,

    name            TEXT NOT NULL UNIQUE,

    hex_code        TEXT,

    is_active       INTEGER NOT NULL DEFAULT 1
                        CHECK(is_active IN (0,1)),

    created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--------------------------------------------------------------------------------
-- Sizes
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS sizes (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,

    code            TEXT NOT NULL UNIQUE,

    display_name    TEXT NOT NULL,

    display_order   INTEGER NOT NULL,

    is_active       INTEGER NOT NULL DEFAULT 1
                        CHECK(is_active IN (0,1)),

    created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--------------------------------------------------------------------------------
-- Products
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS products (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,

    name            TEXT NOT NULL,

    description     TEXT,

    category_id     INTEGER NOT NULL,

    brand_id        INTEGER NOT NULL,

    is_active       INTEGER NOT NULL DEFAULT 1
                        CHECK(is_active IN (0,1)),

    created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(category_id)
        REFERENCES categories(id)
        ON DELETE RESTRICT,

    FOREIGN KEY(brand_id)
        REFERENCES brands(id)
        ON DELETE RESTRICT
);

--------------------------------------------------------------------------------
-- Product Variants
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS product_variants (
    id                  INTEGER PRIMARY KEY AUTOINCREMENT,

    product_id          INTEGER NOT NULL,

    sku                 TEXT NOT NULL UNIQUE,

    color_id            INTEGER NOT NULL,

    size_id             INTEGER NOT NULL,

    barcode             TEXT,

    weight_grams        INTEGER
                            CHECK(weight_grams IS NULL
                            OR weight_grams >= 0),

    is_active           INTEGER NOT NULL DEFAULT 1
                            CHECK(is_active IN (0,1)),

    created_at          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_at          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(product_id)
        REFERENCES products(id)
        ON DELETE CASCADE,

    FOREIGN KEY(color_id)
        REFERENCES colors(id)
        ON DELETE RESTRICT,

    FOREIGN KEY(size_id)
        REFERENCES sizes(id)
        ON DELETE RESTRICT,

    UNIQUE(
        product_id,
        color_id,
        size_id
    )
);

--------------------------------------------------------------------------------
-- Catalog Indexes
--------------------------------------------------------------------------------

CREATE INDEX IF NOT EXISTS idx_products_name
ON products(name);

CREATE INDEX IF NOT EXISTS idx_products_category
ON products(category_id);

CREATE INDEX IF NOT EXISTS idx_products_brand
ON products(brand_id);

CREATE INDEX IF NOT EXISTS idx_variants_product
ON product_variants(product_id);

CREATE INDEX IF NOT EXISTS idx_variants_sku
ON product_variants(sku);

CREATE INDEX IF NOT EXISTS idx_variants_color
ON product_variants(color_id);

CREATE INDEX IF NOT EXISTS idx_variants_size
ON product_variants(size_id);

--------------------------------------------------------------------------------
-- Updated At Triggers
--------------------------------------------------------------------------------

CREATE TRIGGER IF NOT EXISTS trg_categories_updated
AFTER UPDATE ON categories
FOR EACH ROW
BEGIN
    UPDATE categories
       SET updated_at = CURRENT_TIMESTAMP
     WHERE id = OLD.id;
END;

CREATE TRIGGER IF NOT EXISTS trg_brands_updated
AFTER UPDATE ON brands
FOR EACH ROW
BEGIN
    UPDATE brands
       SET updated_at = CURRENT_TIMESTAMP
     WHERE id = OLD.id;
END;

CREATE TRIGGER IF NOT EXISTS trg_colors_updated
AFTER UPDATE ON colors
FOR EACH ROW
BEGIN
    UPDATE colors
       SET updated_at = CURRENT_TIMESTAMP
     WHERE id = OLD.id;
END;

CREATE TRIGGER IF NOT EXISTS trg_sizes_updated
AFTER UPDATE ON sizes
FOR EACH ROW
BEGIN
    UPDATE sizes
       SET updated_at = CURRENT_TIMESTAMP
     WHERE id = OLD.id;
END;

CREATE TRIGGER IF NOT EXISTS trg_products_updated
AFTER UPDATE ON products
FOR EACH ROW
BEGIN
    UPDATE products
       SET updated_at = CURRENT_TIMESTAMP
     WHERE id = OLD.id;
END;

CREATE TRIGGER IF NOT EXISTS trg_product_variants_updated
AFTER UPDATE ON product_variants
FOR EACH ROW
BEGIN
    UPDATE product_variants
       SET updated_at = CURRENT_TIMESTAMP
     WHERE id = OLD.id;
END;--------------------------------------------------------------------------------
-- Product Images
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS product_images (
    id                  INTEGER PRIMARY KEY AUTOINCREMENT,

    product_id          INTEGER NOT NULL,

    image_url           TEXT NOT NULL,

    alt_text            TEXT,

    display_order       INTEGER NOT NULL DEFAULT 1,

    is_primary          INTEGER NOT NULL DEFAULT 0
                            CHECK(is_primary IN (0,1)),

    created_at          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(product_id)
        REFERENCES products(id)
        ON DELETE CASCADE
);

--------------------------------------------------------------------------------
-- Pricing
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS pricing (
    id                      INTEGER PRIMARY KEY AUTOINCREMENT,

    variant_id              INTEGER NOT NULL,

    currency                TEXT NOT NULL
                                CHECK(LENGTH(currency) = 3),

    list_price_cents        INTEGER NOT NULL
                                CHECK(list_price_cents >= 0),

    sale_price_cents        INTEGER
                                CHECK(
                                    sale_price_cents IS NULL OR
                                    sale_price_cents >= 0
                                ),

    effective_from          DATETIME NOT NULL,

    effective_to            DATETIME,

    created_at              DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_at              DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(variant_id)
        REFERENCES product_variants(id)
        ON DELETE CASCADE
);

--------------------------------------------------------------------------------
-- Warehouses
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS warehouses (
    id                  INTEGER PRIMARY KEY AUTOINCREMENT,

    code                TEXT NOT NULL UNIQUE,

    name                TEXT NOT NULL,

    city                TEXT NOT NULL,

    state               TEXT NOT NULL,

    country             TEXT NOT NULL,

    timezone            TEXT NOT NULL,

    is_active           INTEGER NOT NULL DEFAULT 1
                            CHECK(is_active IN (0,1)),

    created_at          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_at          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--------------------------------------------------------------------------------
-- Inventory Levels
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS inventory_levels (
    id                      INTEGER PRIMARY KEY AUTOINCREMENT,

    warehouse_id            INTEGER NOT NULL,

    variant_id              INTEGER NOT NULL,

    quantity_available      INTEGER NOT NULL DEFAULT 0
                                CHECK(quantity_available >= 0),

    quantity_reserved       INTEGER NOT NULL DEFAULT 0
                                CHECK(quantity_reserved >= 0),

    reorder_level           INTEGER NOT NULL DEFAULT 5
                                CHECK(reorder_level >= 0),

    updated_at              DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(warehouse_id)
        REFERENCES warehouses(id)
        ON DELETE CASCADE,

    FOREIGN KEY(variant_id)
        REFERENCES product_variants(id)
        ON DELETE CASCADE,

    UNIQUE(warehouse_id, variant_id)
);

--------------------------------------------------------------------------------
-- Catalog Indexes
--------------------------------------------------------------------------------

CREATE INDEX IF NOT EXISTS idx_product_images_product
ON product_images(product_id);

CREATE INDEX IF NOT EXISTS idx_product_images_primary
ON product_images(product_id, is_primary);

CREATE INDEX IF NOT EXISTS idx_pricing_variant
ON pricing(variant_id);

CREATE INDEX IF NOT EXISTS idx_inventory_variant
ON inventory_levels(variant_id);

CREATE INDEX IF NOT EXISTS idx_inventory_warehouse
ON inventory_levels(warehouse_id);

--------------------------------------------------------------------------------
-- Updated At Triggers
--------------------------------------------------------------------------------

CREATE TRIGGER IF NOT EXISTS trg_pricing_updated
AFTER UPDATE ON pricing
FOR EACH ROW
BEGIN
    UPDATE pricing
       SET updated_at = CURRENT_TIMESTAMP
     WHERE id = OLD.id;
END;

CREATE TRIGGER IF NOT EXISTS trg_warehouses_updated
AFTER UPDATE ON warehouses
FOR EACH ROW
BEGIN
    UPDATE warehouses
       SET updated_at = CURRENT_TIMESTAMP
     WHERE id = OLD.id;
END;

CREATE TRIGGER IF NOT EXISTS trg_inventory_levels_updated
AFTER UPDATE ON inventory_levels
FOR EACH ROW
BEGIN
    UPDATE inventory_levels
       SET updated_at = CURRENT_TIMESTAMP
     WHERE id = OLD.id;
END;

--------------------------------------------------------------------------------
-- Customers
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS customers (
    id                  INTEGER PRIMARY KEY AUTOINCREMENT,

    customer_number     TEXT NOT NULL UNIQUE,

    email               TEXT NOT NULL UNIQUE,

    first_name          TEXT NOT NULL,

    last_name           TEXT NOT NULL,

    phone_number        TEXT,

    is_active           INTEGER NOT NULL DEFAULT 1
                            CHECK(is_active IN (0,1)),

    created_at          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_at          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--------------------------------------------------------------------------------
-- Orders
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS orders (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,

    order_number                TEXT NOT NULL UNIQUE,

    customer_id                 INTEGER NOT NULL,

    warehouse_id                INTEGER NOT NULL,

    shipping_name               TEXT NOT NULL,

    shipping_address_line1      TEXT NOT NULL,

    shipping_address_line2      TEXT,

    shipping_city               TEXT NOT NULL,

    shipping_state              TEXT NOT NULL,

    shipping_postal_code        TEXT NOT NULL,

    shipping_country            TEXT NOT NULL,

    status                      TEXT NOT NULL
                                    CHECK(
                                        status IN (
                                            'PENDING',
                                            'CONFIRMED',
                                            'PROCESSING',
                                            'SHIPPED',
                                            'DELIVERED',
                                            'CANCELLED'
                                        )
                                    ),

    total_amount_cents          INTEGER NOT NULL
                                    CHECK(total_amount_cents >= 0),

    created_at                  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_at                  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(customer_id)
        REFERENCES customers(id)
        ON DELETE RESTRICT,

    FOREIGN KEY(warehouse_id)
        REFERENCES warehouses(id)
        ON DELETE RESTRICT
);

--------------------------------------------------------------------------------
-- Order Items
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS order_items (
    id                      INTEGER PRIMARY KEY AUTOINCREMENT,

    order_id                INTEGER NOT NULL,

    variant_id              INTEGER NOT NULL,

    quantity                INTEGER NOT NULL
                                CHECK(quantity > 0),

    unit_price_cents        INTEGER NOT NULL
                                CHECK(unit_price_cents >= 0),

    created_at              DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(order_id)
        REFERENCES orders(id)
        ON DELETE CASCADE,

    FOREIGN KEY(variant_id)
        REFERENCES product_variants(id)
        ON DELETE RESTRICT
);

--------------------------------------------------------------------------------
-- Commerce Indexes
--------------------------------------------------------------------------------

CREATE INDEX IF NOT EXISTS idx_customers_customer_number
ON customers(customer_number);

CREATE INDEX IF NOT EXISTS idx_customers_email
ON customers(email);

CREATE INDEX IF NOT EXISTS idx_orders_number
ON orders(order_number);

CREATE INDEX IF NOT EXISTS idx_orders_customer
ON orders(customer_id);

CREATE INDEX IF NOT EXISTS idx_orders_status
ON orders(status);

CREATE INDEX IF NOT EXISTS idx_orders_warehouse
ON orders(warehouse_id);

CREATE INDEX IF NOT EXISTS idx_order_items_order
ON order_items(order_id);

CREATE INDEX IF NOT EXISTS idx_order_items_variant
ON order_items(variant_id);

--------------------------------------------------------------------------------
-- Updated At Triggers
--------------------------------------------------------------------------------

CREATE TRIGGER IF NOT EXISTS trg_customers_updated
AFTER UPDATE ON customers
FOR EACH ROW
BEGIN
    UPDATE customers
       SET updated_at = CURRENT_TIMESTAMP
     WHERE id = OLD.id;
END;

CREATE TRIGGER IF NOT EXISTS trg_orders_updated
AFTER UPDATE ON orders
FOR EACH ROW
BEGIN
    UPDATE orders
       SET updated_at = CURRENT_TIMESTAMP
     WHERE id = OLD.id;
END;

--------------------------------------------------------------------------------
-- Returns
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS returns (
    id                      INTEGER PRIMARY KEY AUTOINCREMENT,

    order_id                INTEGER NOT NULL UNIQUE,

    status                  TEXT NOT NULL
                                CHECK (
                                    status IN (
                                        'REQUESTED',
                                        'APPROVED',
                                        'REJECTED',
                                        'COMPLETED'
                                    )
                                ),

    reason                  TEXT NOT NULL,

    requested_at            DATETIME NOT NULL,

    approved_at             DATETIME,

    created_at              DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_at              DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(order_id)
        REFERENCES orders(id)
        ON DELETE CASCADE
);

--------------------------------------------------------------------------------
-- Returns Indexes
--------------------------------------------------------------------------------

CREATE INDEX IF NOT EXISTS idx_returns_order
ON returns(order_id);

CREATE INDEX IF NOT EXISTS idx_returns_status
ON returns(status);

--------------------------------------------------------------------------------
-- Updated At Triggers
--------------------------------------------------------------------------------

CREATE TRIGGER IF NOT EXISTS trg_returns_updated
AFTER UPDATE ON returns
FOR EACH ROW
BEGIN
    UPDATE returns
       SET updated_at = CURRENT_TIMESTAMP
     WHERE id = OLD.id;
END;

--------------------------------------------------------------------------------
-- Initial Schema Version
--------------------------------------------------------------------------------

INSERT OR IGNORE INTO schema_version (
    version
)
VALUES (
    1
);