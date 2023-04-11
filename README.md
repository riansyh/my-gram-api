# MyGram API

This project is a final project submission for Fresh Graduate Academy Scalable Web Service with Golang pelatihan.

## Tech Stack

Golang, Gin Gonic, GORM, PostgreSQL, go-jwt, Railway(for deployment), Swagger (API Documentation)

## Features

-   There are 4 tables with relations between them.
    ![Database Structure](https://i.ibb.co/c8Mxnbz/image-2.png)
-   There are a total of 17 endpoints in the API.
    ![Endpoints List](https://i.ibb.co/GC7kkxF/Group-6.png)
-   Validation is implemented for every column in each table.
-   Authentication is required for every endpoint.
-   Authorization is required for modifying ownership data in every endpoint.

## Run Locally / Installation

Clone the project

```bash
  git clone https://github.com/riansyh/quiz-dot-challenge.git
```

Go to the project directory

```bash
  cd quiz-dot-challenge
```

Install dependencies

```bash
  go mod install
```

Configure the environment variable, consists of:

```bash
  PGHOST=<fill with PostgreSQL host>
  PGUSER=<fill with PostgreSQL user>
  PGPASSWORD=<fill with PostgreSQL database password>
  PGPORT=<fill with PostgreSQL PORT>
  PGDATABASE=<fill with database name>
  PORT=<fill with PORT to run the server>
```

Start the server

```bash
  go run .
```

## Documentation and Deployment

Deploy using Railway in:

https://my-gram-api-production.up.railway.app

For the API documentation you can access this URL:

[Documentation](https://my-gram-api-production.up.railway.app/swagger/index.html)
