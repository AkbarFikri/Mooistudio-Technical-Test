package repository

const CreateProduct = `
INSERT INTO
	products (
	      	 id,
	       	name,
	       	description,
			created_at,
	    	updated_at
) VALUES (
          	:id,
          	:name,
          	:description,
        	:created_at,
	    	:updated_at
)`

const GetAll = `
SELECT
		    id,
	       	name,
	       	description,
			created_at,
	    	updated_at
FROM
    products
`
