# Address Book Backend

Address Book Backend is a CRUD (Create, Read, Update, Delete) application that allows users to manage address information. It provides endpoints to perform operations like adding, retrieving, updating, and deleting address entries.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [API Endpoints](#api-endpoints)
- [License](#license)

# Introduction

The Address Book Backend is designed to manage contact information efficiently. Users can add, retrieve, update, and delete address entries with details such as name, phone, email, address, company, job title, website, and notes.

## Features

- Create a new address entry with details such as name, phone, email, address, etc.
- Retrieve all address entries.
- Retrieve a specific address entry by ID.
- Update a specific address entry by ID.
- Delete a specific address entry by ID.

## Prerequisites

- [Go installed on your machine.](https://golang.org/dl/)
- [Git installed on your machine.](https://git-scm.com/downloads)
- [PostgreSQL installed and running.](https://www.postgresql.org/download/)
- [You can alternatively use ElephantSQL for a quick and free DB setup.](https://www.elephantsql.com/)

# Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/soloappsnmobile/address-book-backend.git
    cd address-book-backend
    ```

2. **Install dependencies:**

    ```bash
    go mod tidy
    ```

3. **Build the project:**

    ```bash
    go build -o address-book .
    ```

## Configuration

1. **Environment Variables:**

    Create a copy of the example environment file:

    ```bash
    cp example.env .env
    ```

    Update the values in `.env` with your configuration.

2. **Database Configuration:**

    Make sure to set up your PostgreSQL database and update the database configuration in your `.env` file.



## API Endpoints

- `POST /address-books`: Create a new address entry.
- `GET /address-books`: Retrieve all address entries.
- `GET /address-book/{id}`: Retrieve a specific address entry by ID.
- `PUT /address-book/{id}`: Update a specific address entry by ID.
- `DELETE /address-book/{id}`: Delete a specific address entry by ID.

### Request and Response Examples

#### `POST /address-books`

**Request:**

```bash
curl -X POST -H "Content-Type: application/json" -d
'{  "name": "John Doe",
    "phone": "+1234567890",
    "email": "john.doe@example.com",
    "address": "123 Main St",
    "company": "",
    "job_title": "",
    "website": "",
    "note": ""
}'
```

**Response:**
```json
{
  "data": {
    "id": 18,
    "name": "John Doe",
    "phone": "+123456789",
    "email": "john.doe@example.com",
    "company": "",
    "job_title": "",
    "website": "",
    "address": "123 Main St",
    "note": " "
  },
  "resp_code": "00",
  "resp_desc": "Address book created successfully"
}
```


#### `GET /address-books`

**Request:**
```bash
curl http://localhost:8080/address-books
```

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "name": "Solomon",
      "phone": "05545355687",
      "email": "Solomon@gmail.com",
      "company": "",
      "job_title": "",
      "website": "",
      "address": "somewher",
      "note": ""
    },
   {
      "id": 18,
      "name": "John Doe",
      "phone": "+123456789",
      "email": "john.doe@example.com",
      "company": "",
      "job_title": "",
      "website": "",
      "address": "123 Main St",
      "note": " "
    }
  ],
  "resp_code": "00",
  "resp_desc": "Address books retrieved successfully"
}
```


#### `GET /address-book/{id}`

**Request:**

```bash
curl http://localhost:8080/address-books/18
```

**Response:**
```json
{
  "data": {
    "id": 18,
    "name": "John Doe",
    "phone": "+123456789",
    "email": "john.doe@example.com",
    "company": "",
    "job_title": "",
    "website": "",
    "address": "123 Main St",
    "note": " "
  },
  "resp_code": "00",
  "resp_desc": "Address book retrieved successfully"
}
```

#### `PUT /address-book/{id}`

**Request:**

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"name": "Updated Name"}' http://localhost:8080/address-book/18

```

**Response:**
```json
{
  "data": {
    "id": 18,
    "name": "Updated Name",
    "phone": "+123456789",
    "email": "john.doe@example.com",
    "company": "",
    "job_title": "",
    "website": "",
    "address": "123 Main St",
    "note": " "
  },
  "resp_code": "00",
  "resp_desc": "Address book updated successfully"
}
```

#### `DELETE /address-book/{id}`

**Request:**

```bash
curl -X DELETE http://localhost:8080/address-book/18
```

**Response:**
```json
{
  "resp_code": "00",
  "resp_desc": "Address book deleted successfully"
}
```


## License
MIT License

Copyright (c) 2024 Solomon Tetteh Sackey






