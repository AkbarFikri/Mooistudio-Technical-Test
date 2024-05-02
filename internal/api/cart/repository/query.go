package repository

const CreateCart = `
INSERT INTO carts (
				id,
        		user_id,
        		product_id,
                quantity,
                created_at,
                updated_at
) VALUES (
          :id,
          :user_id,
          :product_id,
          :quantity,
          :created_at,
          :updated_at
)`

const GetAllByUserId = `
SELECT
	carts.id,
	carts.user_id,
	products.name as product_name,
	products.id as product_id,
	products.price as product_price,
	carts.quantity
FROM
    carts
JOIN products ON carts.product_id = products.id
WHERE carts.user_id = :user_id`

const DeleteAllByUserId = `
DELETE FROM carts
WHERE user_id = :user_id`

const DeleteCart = `
DELETE FROM carts
WHERE id = :id`

const GetById = `
SELECT
	*
FROM
    carts
WHERE id = :id
LIMIT 1
`
