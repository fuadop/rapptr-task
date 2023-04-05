# Todo List API

## Table of contents

1. [Documentation](#documentation)
2. [Run Application](#run-application)
3. [Common API request body validation](#common-field-validations)

## Documentation

The application documentation is located at the `/docs/index.html` and `/swagger/index.html` routes on the server.

## Run Application

To start the application, change directory to the root directory and run `docker compose up` or `docker-compose up` command depending on your docker version.

```bash
docker compose up
```

To stop the application, change directory to the root directory and run `docker compose down` or `docker-comopse down` command depending on your docker version.

```bash
docker compose down
```

## Common Field Validations

### /auth/login [POST]

- password
	- minimum of 8 characters

### /auth/register [POST]

- password
	- minimum of 8 characters
