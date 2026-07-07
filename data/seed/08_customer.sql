BEGIN TRANSACTION;

INSERT INTO customers (
    customer_number,
    email,
    first_name,
    last_name,
    phone_number
)
VALUES

('CUST0001','john.miller@example.com','John','Miller','+1-555-1001'),

('CUST0002','emma.johnson@example.com','Emma','Johnson','+1-555-1002'),

('CUST0003','olivia.wilson@example.com','Olivia','Wilson','+1-555-1003'),

('CUST0004','liam.brown@example.com','Liam','Brown','+1-555-1004'),

('CUST0005','noah.davis@example.com','Noah','Davis','+1-555-1005'),

('CUST0006','ava.moore@example.com','Ava','Moore','+1-555-1006'),

('CUST0007','ethan.taylor@example.com','Ethan','Taylor','+1-555-1007'),

('CUST0008','mia.anderson@example.com','Mia','Anderson','+1-555-1008'),

('CUST0009','lucas.thomas@example.com','Lucas','Thomas','+1-555-1009'),

('CUST0010','sophia.jackson@example.com','Sophia','Jackson','+1-555-1010');

COMMIT;