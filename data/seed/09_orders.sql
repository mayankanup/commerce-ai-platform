BEGIN TRANSACTION;

--------------------------------------------------------------------------------
-- Orders
--------------------------------------------------------------------------------

INSERT INTO orders (
    order_number,
    customer_id,
    warehouse_id,
    shipping_name,
    shipping_address_line1,
    shipping_address_line2,
    shipping_city,
    shipping_state,
    shipping_postal_code,
    shipping_country,
    status,
    total_amount_cents
)
VALUES

-- Delivered -------------------------------------------------------------

('ORD-2026-000001',1,1,'John Miller','101 Main St','','Los Angeles','California','90001','USA','DELIVERED',1799),
('ORD-2026-000002',2,2,'Emma Johnson','45 Lake Shore Dr','','Chicago','Illinois','60601','USA','DELIVERED',4499),
('ORD-2026-000003',3,3,'Olivia Wilson','12 Madison Ave','','New York','New York','10001','USA','DELIVERED',6999),
('ORD-2026-000004',4,1,'Liam Brown','200 Sunset Blvd','','Los Angeles','California','90011','USA','DELIVERED',3499),
('ORD-2026-000005',5,2,'Noah Davis','300 State St','','Chicago','Illinois','60605','USA','DELIVERED',2499),
('ORD-2026-000006',6,3,'Ava Moore','55 Park Ave','','New York','New York','10011','USA','DELIVERED',7999),
('ORD-2026-000007',7,1,'Ethan Taylor','77 Ocean Ave','','Los Angeles','California','90022','USA','DELIVERED',3999),
('ORD-2026-000008',8,2,'Mia Anderson','88 Michigan Ave','','Chicago','Illinois','60611','USA','DELIVERED',5499),
('ORD-2026-000009',9,3,'Lucas Thomas','99 Broadway','','New York','New York','10022','USA','DELIVERED',1799),
('ORD-2026-000010',10,1,'Sophia Jackson','11 Hollywood Blvd','','Los Angeles','California','90028','USA','DELIVERED',8999),
('ORD-2026-000011',1,2,'John Miller','101 Main St','','Los Angeles','California','90001','USA','DELIVERED',2999),
('ORD-2026-000012',2,3,'Emma Johnson','45 Lake Shore Dr','','Chicago','Illinois','60601','USA','DELIVERED',4999),
('ORD-2026-000013',3,1,'Olivia Wilson','12 Madison Ave','','New York','New York','10001','USA','DELIVERED',3499),
('ORD-2026-000014',4,2,'Liam Brown','200 Sunset Blvd','','Los Angeles','California','90011','USA','DELIVERED',6999),
('ORD-2026-000015',5,3,'Noah Davis','300 State St','','Chicago','Illinois','60605','USA','DELIVERED',2499),

-- Shipped ---------------------------------------------------------------

('ORD-2026-000016',6,1,'Ava Moore','55 Park Ave','','New York','New York','10011','USA','SHIPPED',1799),
('ORD-2026-000017',7,2,'Ethan Taylor','77 Ocean Ave','','Los Angeles','California','90022','USA','SHIPPED',3999),
('ORD-2026-000018',8,3,'Mia Anderson','88 Michigan Ave','','Chicago','Illinois','60611','USA','SHIPPED',7999),
('ORD-2026-000019',9,1,'Lucas Thomas','99 Broadway','','New York','New York','10022','USA','SHIPPED',3499),
('ORD-2026-000020',10,2,'Sophia Jackson','11 Hollywood Blvd','','Los Angeles','California','90028','USA','SHIPPED',5499),

-- Processing ------------------------------------------------------------

('ORD-2026-000021',1,3,'John Miller','101 Main St','','Los Angeles','California','90001','USA','PROCESSING',2999),
('ORD-2026-000022',2,1,'Emma Johnson','45 Lake Shore Dr','','Chicago','Illinois','60601','USA','PROCESSING',8999),
('ORD-2026-000023',3,2,'Olivia Wilson','12 Madison Ave','','New York','New York','10001','USA','PROCESSING',4499),

-- Confirmed -------------------------------------------------------------

('ORD-2026-000024',4,3,'Liam Brown','200 Sunset Blvd','','Los Angeles','California','90011','USA','CONFIRMED',3499),
('ORD-2026-000025',5,1,'Noah Davis','300 State St','','Chicago','Illinois','60605','USA','CONFIRMED',2499),

-- Pending ---------------------------------------------------------------

('ORD-2026-000026',6,2,'Ava Moore','55 Park Ave','','New York','New York','10011','USA','PENDING',6999),
('ORD-2026-000027',7,3,'Ethan Taylor','77 Ocean Ave','','Los Angeles','California','90022','USA','PENDING',1799),
('ORD-2026-000028',8,1,'Mia Anderson','88 Michigan Ave','','Chicago','Illinois','60611','USA','PENDING',3999),
('ORD-2026-000029',9,2,'Lucas Thomas','99 Broadway','','New York','New York','10022','USA','PENDING',5499),
('ORD-2026-000030',10,3,'Sophia Jackson','11 Hollywood Blvd','','Los Angeles','California','90028','USA','PENDING',8999);

--------------------------------------------------------------------------------
-- Order Items
--------------------------------------------------------------------------------

INSERT INTO order_items (
    order_id,
    variant_id,
    quantity,
    unit_price_cents
)
SELECT
    id,
    ((id - 1) % 60) + 1,
    CASE
        WHEN id % 5 = 0 THEN 2
        ELSE 1
    END,
    total_amount_cents
FROM orders;

COMMIT;