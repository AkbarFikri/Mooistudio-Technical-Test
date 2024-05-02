package repository

const CreateOrder = `
INSERT INTO orders (
            id,
            status,
            user_id,
            total
) VALUES (
        	:id,
          	:status,
          	:user_id,
          	:total
)`

const CreateOrderItem = `
INSERT INTO order_items (
            id,
            order_id,
            product_id,
            quantity
) VALUES (
            :id,
            :order_id,
            :product_id,
            :quantity
)`

const GetAllByUserId = `
SELECT 
    id, 
    status, 
    user_id, 
    total, 
    created_at 
FROM 
    orders
WHERE
    user_id = :user_id`

const GetOneById = `
SELECT 
    id, 
    status, 
    user_id, 
    total, 
    created_at 
FROM 
    orders
WHERE
    id = :id
LIMIT 1`

const GetOrderItems = `
SELECT
	order_items.id,
	products.name as product_name,
	products.id as product_id,
	products.price as product_price,
	order_items.quantity
FROM
    order_items
JOIN products ON order_items.product_id = products.id
WHERE order_items.order_id = :order_id`
