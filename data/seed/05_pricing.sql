BEGIN TRANSACTION;

--------------------------------------------------------------------------------
-- Product Pricing
--------------------------------------------------------------------------------

INSERT INTO pricing (
    variant_id,
    currency,
    list_price_cents,
    sale_price_cents,
    effective_from
)
SELECT
    pv.id,
    'USD',

    CASE p.category_id

        -- T-Shirts
        WHEN 1 THEN 1999

        -- Shirts
        WHEN 2 THEN 3499

        -- Hoodies
        WHEN 3 THEN 4999

        -- Jackets
        WHEN 4 THEN 7999

        -- Jeans
        WHEN 5 THEN 5999

        -- Shorts
        WHEN 6 THEN 2999

        -- Shoes
        WHEN 7 THEN 8999

        -- Accessories
        WHEN 8 THEN 2499

        -- Activewear
        WHEN 9 THEN 3999

        -- Sweaters
        WHEN 10 THEN 5499

    END,

    CASE p.category_id

        WHEN 1 THEN 1799
        WHEN 2 THEN NULL
        WHEN 3 THEN 4499
        WHEN 4 THEN 6999
        WHEN 5 THEN NULL
        WHEN 6 THEN 2499
        WHEN 7 THEN 7999
        WHEN 8 THEN NULL
        WHEN 9 THEN 3499
        WHEN 10 THEN 4999

    END,

    CURRENT_TIMESTAMP

FROM product_variants pv
JOIN products p
    ON p.id = pv.product_id;

COMMIT;