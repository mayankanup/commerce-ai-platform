BEGIN TRANSACTION;

--------------------------------------------------------------------------------
-- Products
--------------------------------------------------------------------------------

INSERT INTO products (name, description, category_id, brand_id)
VALUES

('Essential Cotton T-Shirt',
 '100% premium cotton everyday t-shirt.',
 1, 1),

('Classic Polo Shirt',
 'Soft cotton polo with modern fit.',
 2, 2),

('Premium Pullover Hoodie',
 'Warm fleece hoodie for everyday comfort.',
 3, 3),

('Performance Running Hoodie',
 'Lightweight hoodie for outdoor running.',
 3, 7),

('Denim Trucker Jacket',
 'Classic denim jacket.',
 4, 4),

('Water Resistant Windbreaker',
 'Lightweight weather resistant jacket.',
 4, 5),

('Slim Fit Jeans',
 'Stretch denim slim fit jeans.',
 5, 6),

('Relaxed Fit Jeans',
 'Comfort fit denim.',
 5, 8),

('Performance Running Shorts',
 'Moisture wicking athletic shorts.',
 6, 7),

('Cargo Shorts',
 'Multi-pocket outdoor shorts.',
 6, 5),

('Everyday Sneakers',
 'Comfortable casual sneakers.',
 7, 9),

('Trail Running Shoes',
 'Designed for off-road running.',
 7, 5),

('Leather Belt',
 'Genuine leather belt.',
 8, 10),

('Canvas Backpack',
 'Lightweight everyday backpack.',
 8, 2),

('Training Joggers',
 'Performance joggers for workouts.',
 9, 7),

('Compression Tights',
 'Athletic compression tights.',
 9, 7),

('Crew Neck Sweater',
 'Soft knitted sweater.',
 10, 8),

('Half Zip Sweater',
 'Warm half-zip sweater.',
 10, 1),

('Oxford Button Down Shirt',
 'Business casual oxford shirt.',
 2, 10),

('Graphic T-Shirt',
 'Modern printed cotton t-shirt.',
 1, 6);

COMMIT;