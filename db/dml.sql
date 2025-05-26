
-- ======= CUSTOMER ======= --

-- Insert customer
INSERT INTO customer(customer_id, name, phone, address)
VALUES (1, 'Restu Adi', '08123456789', 'Jl. Mawar');

-- Update customer
UPDATE customer
SET name = 'Restu Wahyujati', phone = '08987654321', address = 'Jl. Melati', updated_at = CURRENT_TIMESTAMP
WHERE customer_id = 1;

-- Delete customer (hanya jika tidak ada relasi dengan order)
DELETE FROM customer
WHERE customer_id = 1;

-- Select all customers
SELECT customer_id, name, phone, address, created_at, updated_at
FROM customer;

-- Select customer by ID
SELECT customer_id, name, phone, address, created_at, updated_at
FROM customer
WHERE customer_id = 1;

-- Check if customer ID exists
SELECT EXISTS(SELECT 1 FROM customer WHERE customer_id = 1);

-- Check if customer ID used in orders
SELECT EXISTS(SELECT 1 FROM "order" WHERE customer_id = 1);




-- =======SERVICE======= --

-- Insert service
INSERT INTO service(service_id, service_name, unit, price)
VALUES (101, 'Cuci Kering', 'kg', 10000);

-- Update service
UPDATE service
SET service_name = 'Cuci Basah', unit = 'kg', price = 12000, updated_at = CURRENT_TIMESTAMP
WHERE service_id = 101;

-- Delete service (hanya jika tidak digunakan dalam order_detail)
DELETE FROM service
WHERE service_id = 101;

-- Select all services
SELECT * FROM service;

-- Select service by ID
SELECT service_id, service_name, unit, price, created_at, updated_at
FROM service
WHERE service_id = 101;

-- Check if service ID exists
SELECT EXISTS(SELECT 1 FROM service WHERE service_id = 101);

-- Check if service ID used in orders
SELECT EXISTS(SELECT 1 FROM order_detail WHERE service_id = 101);





-- ======= ORDER + ORDER_DETAIL ======= --

-- Insert order (transaksi)
BEGIN;

INSERT INTO "order"(order_id, customer_id, order_date, received_by)
VALUES (1001, 1, CURRENT_DATE, 'Restu');

INSERT INTO order_detail(order_id, service_id, qty)
VALUES 
(1001, 101, 2),
(1001, 102, 3);

COMMIT;

-- Update order finish (complete order)
UPDATE "order"
SET completion_date = '2025-05-26'
WHERE order_id = 1001;

-- Get all orders
SELECT order_id, customer_id, order_date, completion_date, received_by
FROM "order";

-- Get order by ID
SELECT order_id, customer_id, order_date, completion_date, received_by
FROM "order"
WHERE order_id = 1001;

-- Get order details
SELECT order_detail_id, order_id, service_id, qty
FROM order_detail
WHERE order_id = 1001;

-- Check if order ID exists
SELECT EXISTS(SELECT 1 FROM "order" WHERE order_id = 1001);

-- Check if customer ID exists
SELECT EXISTS(SELECT 1 FROM customer WHERE customer_id = 1);