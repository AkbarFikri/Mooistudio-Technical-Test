package repository

const CreateUser = `
INSERT INTO
	users (
	       id,
	       email,
	       password,
	       full_name
) VALUES (
          :id,
          :email,
          :password,
          :full_name
)`

const GetUserByEmail = `
SELECT
	id,
	full_name,
	password,
	email
FROM
	users
WHERE
    email = :email`

const CountEmail = `
SELECT 
	COUNT(email) 
FROM 
	users 
WHERE 
	email = :email
`
