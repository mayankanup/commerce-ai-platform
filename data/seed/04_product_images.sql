BEGIN TRANSACTION;

INSERT INTO product_images(product_id, image_url, alt_text, display_order, is_primary)
SELECT
    id,
    'https://picsum.photos/seed/product-' || id || '/800/800',
    name,
    1,
    1
FROM products;

COMMIT;