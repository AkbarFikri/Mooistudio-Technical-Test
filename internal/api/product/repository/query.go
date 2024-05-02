package repository

const CreateProduct = `
INSERT INTO
	products (
	      	id,
	       	name,
	        category_id,
	       	description,
	        price,
			created_at,
	    	updated_at
) VALUES (
          	:id,
          	:name,
          	:category_id,
          	:description,
          	:price,
        	:created_at,
	    	:updated_at
)`

const GetProductById = `
SELECT
	*
FROM
    products
WHERE
    id = :id
`

const GetAll = `
SELECT
		    products.id,
	       	products.name,
	       	products.category_id,
	       	categories.name as category_name,
	       	products.description,
	       	products.price,
			products.created_at,
	    	products.updated_at
FROM
    products
JOIN categories ON categories.id = products.category_id
`

const CreateProductCategory = `
INSERT INTO
	categories (
	        id,
	        name
) VALUES (
        	:id,
          	:name
)`

const GetAllCategories = `
SELECT
		    id,
	       	name
FROM
    categories
`
