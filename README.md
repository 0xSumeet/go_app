# Go Ecommerce Application 

## Project Overview
This backend project involves working with CRUD operation in GO making use of Gin Framework, And GORM for interacting with Postgres database

### Prerequisites
Before you start, make sure to change `.emv.example` to `.env` and provide the relevant enviornment variable fields needed for **postgres** authentication.
```bash
cp .env.example .env
```
#### Setting Up Postgres database
- Set up postgres container using `docker-compose.yml` for database connectivity to application, here the compose file can persists your database data.
- Also, provide the same credentials in the `.env` file.
  ```bash
  docker compose up -d
  ```
- Or simply us `docker run`:
  ```bash
  docker run -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=db_name --name=postgres_container -d -p 5432:5432 postgres:14
  ```
- Access the `postgres` container in linux:
  ```bash
  docker exec -it postgres_container psql -U postgres
  ```

### Getting Started
1. First, Make sure that you have `go` installed into your system.
2. Clone the repo into your local environment.
   ```bash
   git clone https://github.com/0xSumeet/go_app.git
   ```
3. Navigate to the project root directory and add the following command to download all the modules required for the project.
   ```bash
   cd go_app; go mod tidy 
   ```
4. Run the `main.go` file
   ```bash
   go run main.go
   ```
5. This is a backend project, so you can use [Postman](https://www.postman.com/) or **curl command** (linux) to interact with the application.

#### Sending request
**Register User**
  - Make a `POST` request to `http://localhost:8080/register` with the following JSON body
    ```json
    {
      "username": "admin",
      "password": "admin123",
      "role": "Administrator"
    }
    ```
**Login User**
  - Make a `POST` request to `http://localhost:8080/login` with the following JSON body:
    ```json
    {
      "username": "admin",
      "password": "admin123"
    }
    ```
**Logout User**
  - Make a `POST` request to `http://localhost:8080/logout`

**Access Protection Endpoint**
  - Make a `GET` request to `http://localhost:8080/customer-management` with Header as `Authorization: Bearer <your_jwt_token>`

**Note :** The `permission` table data have not been inserted. So, it needs to be added manually for providing certain read/write permission to the departments. You can use tool like `psql`, pgAdmin or DBeaver to manually check and insert data into the database. 
  ```sql
  INSERT INTO permissions (role, module, access) VALUES 
  ('Sales', 'customer-management', 'read/write'),
  ('Sales', 'billing-management', 'read/write'),
  ('Accountant', 'customer-management', 'view only'),
  ('Accountant', 'payroll-management', 'view only'),
  ('HR', 'payroll-management', 'read/write'),
  ('Administrator', 'user-management', 'read/write');
  ``` 
