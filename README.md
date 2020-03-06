# customer-service
Sample crud operation using Golang and MySql


## API ENDPOINTS

### All customers
- Path : `/customers`
- Method: `GET`
- Response: `200`

### Create Customers
- Path : `/customers`
- Method: `POST`
- Fields: `name`
- Response: `201`

### Details a Customers
- Path : `/customers/{id}`
- Method: `GET`
- Response: `200`

### Update Customers
- Path : `/customers/{id}`
- Method: `PUT`
- Fields: `name`
- Response: `200`

### Delete Customers
- Path : `/customers/{id}`
- Method: `DELETE`
- Response: `204`

## Required Packages
- Dependency management
    * [dep](https://github.com/golang/dep)
- Database
    * [MySql](https://github.com/go-sql-driver/mysql)
- Routing
    * [chi](https://github.com/go-chi/chi)


## DB Migration
    CREATE DATABASE `customer-service`
    sql/customers.sql


## Quick Run Project
    Run project by go main.go
    Alternatively you can update the run.sh file and execute 
    `bash run.sh`

```
cd customer-service

chmod +x build.sh
./build.sh

docker-compose up -d
```

