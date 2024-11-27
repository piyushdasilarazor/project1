# project1


## **Features**
- Create a new user
- Retrieve all users
- Retrieve a user by ID
- Update an existing user
- Delete a user by ID

---

## **Technologies Used**
- [Go](https://golang.org/) (Programming Language)
- [Gin Web Framework](https://gin-gonic.com/)
- In-memory data store


## **Directory Structure**

```/awesomeProject
├── main.go                # Entry point
├── routes
│   ├── user_routes.go     # Route definitions
├── services
│   ├── user_service.go    # Business logic
├── models
│   ├── user.go
``` 


## API Endpoints

| **Method** | **Endpoint**       | **Description**        | **Payload Example**                |
|------------|--------------------|------------------------|-------------------------------------|
| POST       | `/users`           | Create a new user      | `{ "name": "John", "age": 30 }`    |
| GET        | `/users`           | Get all users          | -                                   |
| GET        | `/users/:id`       | Get a user by ID       | -                                   |
| PUT        | `/users/:id`       | Update a user by ID    | `{ "name": "Jane", "age": 28 }`    |
| DELETE     | `/users/:id`       | Delete a user by ID    | -                                   |


## Request/Response Examples

### **1. Create a New User**

**Request:**
```bash
curl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{
  "name": "Alice",
  "age": 25
}'
```

**Response:**
```bash
{
  "id": 1,
  "name": "Alice",
  "age": 25
}
```

### **2. Get All Users**

**Request:**
```bash
curl -X GET http://localhost:8080/users
```

**Response:**
```bash
[
  {
    "id": 1,
    "name": "Alice",
    "age": 25
  },
  {
    "id": 2,
    "name": "Bob",
    "age": 30
  }
]
```

### **3. Get a User by ID**

**Request:**
```bash
curl -X GET http://localhost:8080/users/1
```

**Response:**
```bash
{
  "id": 1,
  "name": "Alice",
  "age": 25
}
```
If the user is not found:
```bash
{
  "error": "User not found"
}
```
### **4. Update a User by ID**

**Request:**
```bash
curl -X PUT http://localhost:8080/users/1 \
-H "Content-Type: application/json" \
-d '{
  "name": "Alice Updated",
  "age": 26
}'
```

**Response:**
```bash
{
  "id": 1,
  "name": "Alice Updated",
  "age": 26
}
```

If the user is not found:
```bash
{
  "error": "User not found"
}
```

### **5. Delete a User by ID**

**Request:**
```bash
curl -X DELETE http://localhost:8080/users/1
```

**Response:**
```bash
{
  "message": "User deleted"
}
```

If the user is not found:
```bash
{
  "error": "User not found"
}
```



