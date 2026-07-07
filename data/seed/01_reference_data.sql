BEGIN TRANSACTION;

--------------------------------------------------------------------------------
-- Categories
--------------------------------------------------------------------------------

INSERT INTO categories (name, description) VALUES
('T-Shirts', 'Casual short sleeve and long sleeve t-shirts'),
('Shirts', 'Formal and casual shirts'),
('Hoodies', 'Pullover and zip hoodies'),
('Jackets', 'Lightweight and winter jackets'),
('Jeans', 'Denim jeans'),
('Shorts', 'Casual and athletic shorts'),
('Shoes', 'Running, casual and formal shoes'),
('Accessories', 'Belts, caps, wallets and more'),
('Activewear', 'Workout and performance clothing'),
('Sweaters', 'Knitted sweaters and pullovers');

--------------------------------------------------------------------------------
-- Brands
--------------------------------------------------------------------------------

INSERT INTO brands (name) VALUES
('NorthPeak'),
('UrbanThread'),
('BlueRiver'),
('SummitWear'),
('EverTrail'),
('MetroStyle'),
('NovaFit'),
('CedarLane'),
('Pioneer Apparel'),
('Horizon Clothing');

--------------------------------------------------------------------------------
-- Colors
--------------------------------------------------------------------------------

INSERT INTO colors (name, hex_code) VALUES
('Black', '#000000'),
('White', '#FFFFFF'),
('Navy', '#001F54'),
('Gray', '#808080'),
('Red', '#D32F2F'),
('Blue', '#1976D2'),
('Green', '#388E3C'),
('Brown', '#795548');

--------------------------------------------------------------------------------
-- Sizes
--------------------------------------------------------------------------------

INSERT INTO sizes (code, display_name, display_order) VALUES
('XS', 'Extra Small', 1),
('S',  'Small',       2),
('M',  'Medium',      3),
('L',  'Large',       4),
('XL', 'Extra Large', 5),
('XXL','2X Large',    6);

COMMIT;