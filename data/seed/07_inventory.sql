BEGIN TRANSACTION;

--------------------------------------------------------------------------------
-- Los Angeles (Warehouse 1)
-- Primary fulfillment center
--------------------------------------------------------------------------------

INSERT INTO inventory_levels (
    warehouse_id,
    variant_id,
    quantity_available,
    quantity_reserved,
    reorder_level
)
SELECT
    1,
    id,
    100,
    5,
    20
FROM product_variants;

--------------------------------------------------------------------------------
-- Chicago (Warehouse 2)
-- Medium inventory
--------------------------------------------------------------------------------

INSERT INTO inventory_levels (
    warehouse_id,
    variant_id,
    quantity_available,
    quantity_reserved,
    reorder_level
)
SELECT
    2,
    id,

    CASE
        WHEN id BETWEEN 16 AND 20 THEN 3
        ELSE 45
    END,

    2,
    10

FROM product_variants;

--------------------------------------------------------------------------------
-- New York (Warehouse 3)
-- Smaller warehouse
--------------------------------------------------------------------------------

INSERT INTO inventory_levels (
    warehouse_id,
    variant_id,
    quantity_available,
    quantity_reserved,
    reorder_level
)
SELECT
    3,
    id,

    CASE
        WHEN id BETWEEN 31 AND 35 THEN 0
        ELSE 18
    END,

    1,
    8

FROM product_variants;

COMMIT;