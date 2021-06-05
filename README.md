# A Simple API rest using Go lang

## This api is able to CRUD a entity of Book

### Routes

`
GET http://localhost:5000/api/v1/books
List all Books

GET http://localhost:5000/api/v1/books/:id
List a single Book

POST http://localhost:5000/api/v1/books/
Create a Book

PUT http://localhost:5000/api/v1/books/
Update a Book

DELETE http://localhost:5000/api/v1/books/:id
Delete a Book (soft delete)
`

### Payloads

```js

// creating a book...

// Body

{
  "name":"Book 1",
  "description":"A nice book",
  "medium_price":199.99,
  "author":"the book author",
  "img_url":"url.host"
}




// updating a book

// Body

{
  "id":1,
  "name":"Book 1",
  "description":"A nice book",
  "medium_price":199.99,
  "author":"the book author",
  "img_url":"url.host"
}

```

### how to run the project ?

- Start the Database Client

```shell

$ docker-compose up

```

- Create a database

```sql

# you can run this command in your db client

CREATE DATABASE books

```

- Start the application

```shell

$ go run main.go

```
