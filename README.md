# 🏠 EZKost

EZKost is a simple web-based kost management system for owners and operators built using Golang, Gin Framework, GORM, and Clean Architecture.

## 📋 Features

* ✅ Room Management (CRUD)
* ✅ Tenant Management (CRUD)
* ✅ Payment Management (CRUD)
* ✅ Expense Management (CRUD)
* ✅ Dashboard Summary
* ✅ Authentication & Authorization (JWT)
* ✅ Automatic room status updates
* ✅ Overdue payment tracking

## 🛠️ Tech Stack

- **Backend**: Golang 1.21
- **Framework**: Gin
- **ORM**: GORM
- **Database**: PostgreSQL
- **Auth**: JWT
- **Password**: bcrypt
- **Architecture**: Clean Architecture

## 🏗️ Clean Architecture Layers

```
EZKost/
├── cmd/
│   └── main.go                     # Application entry point
├── internal/
│   ├── config/                     # Configuration
│   │   └── config.go
│   ├── domain/                     # Domain Layer (Business Logic)
│   │   ├── entity/                 # Business Entities
│   │   │   └── entity.go
│   │   └── repository/             # Repository Interfaces
│   │       └── repository.go
│   ├── usecase/                    # Use Case Layer (Application Logic)
│   │   ├── auth_usecase.go
│   │   └── usecase.go
│   ├── repository/                 # Repository Implementation
│   │   ├── model/                  # GORM Models
│   │   │   └── model.go
│   │   └── repository_impl.go
│   └── delivery/                   # Delivery Layer (Presentation)
│       └── http/
│           ├── handler.go          # HTTP Handlers
│           ├── route.go            # Routes
│           └── middleware/
│               └── auth.go
└── pkg/
└── database/                   # Database utilities
└── database.go
```

## 📂 Layer Explanation

### 1. Domain Layer (`internal/domain/`)

This layer contains pure business logic without external dependencies:

- **Entity**: Business object representations (User, Room, Tenant, Payment, Expense)
- **Repository Interface**: Contracts for data access

### 2. Use Case Layer (`internal/usecase/`)

This layer contains **application logic**:

- Business rule implementations
- Flow orchestration between repositories
- Independent of frameworks or databases

### 3. Repository Layer (`internal/repository/`)

This layer contains **data access logic**:

- Repository interface implementations
- GORM models for database mapping
- Conversion between entities and models

### 4. Delivery Layer (`internal/delivery/`)

This layer contains **presentation logic**:

- HTTP handlers
- Request/response mapping
- Middleware
- Routes

### 5. Infrastructure (`pkg/`)

Shared utilities and infrastructure-related code.

## 🎯 Benefits of Clean Architecture

✅ **Testability**: Each layer can be tested independently
✅ **Maintainability**: Code is well-structured and easy to maintain
✅ **Scalability**: Easy to add new features without breaking existing ones
✅ **Flexibility**: Easy to replace frameworks or databases
✅ **Separation of Concerns**: Each layer has a clear responsibility

## 🚀 How to Run
### Docker

```bash
Build and run with docker-compose
docker-compose up -d

Check logs
docker-compose logs -f api
```

The server will run at [http://localhost:8080](http://localhost:8080)

## 📡 API Endpoints

### Authentication
```
POST   /api/v1/auth/login      - Login
POST   /api/v1/auth/register   - Register (first admin)
```

### Dashboard
```
GET    /api/v1/dashboard/summary  - Dashboard summary
```

### Rooms
```
GET    /api/v1/rooms           - List all rooms
GET    /api/v1/rooms/:id       - Room details
POST   /api/v1/rooms           - Create room
PUT    /api/v1/rooms/:id       - Update room
DELETE /api/v1/rooms/:id       - Delete room
```

### Tenants
```
GET    /api/v1/tenants         - List all tenants
GET    /api/v1/tenants/:id     - Tenant details
POST   /api/v1/tenants         - Create tenant
PUT    /api/v1/tenants/:id     - Update tenant
DELETE /api/v1/tenants/:id     - Delete tenant
```

### Payments
```
GET    /api/v1/payments                - List all payments
GET    /api/v1/payments/:id            - Payment details
GET    /api/v1/payments/tenant/:id     - Payments by tenant
GET    /api/v1/payments/overdue        - Overdue payments
POST   /api/v1/payments                - Create payment
PUT    /api/v1/payments/:id            - Update payment
```

### Expenses
```
GET    /api/v1/expenses        - List all expenses
GET    /api/v1/expenses/:id    - Expense details
POST   /api/v1/expenses        - Create expense
PUT    /api/v1/expenses/:id    - Update expense
DELETE /api/v1/expenses/:id    - Delete expense
```

## 🔑 Example Requests

### Register First Admin
```bash
curl -X POST [http://localhost:8080/api/v1/auth/register](http://localhost:8080/api/v1/auth/register)
-H "Content-Type: application/json"
-d '{
"name": "Admin",
"email": "[admin@kos.com](mailto:admin@kos.com)",
"password": "password123",
"role": "owner"
}'
```

### Login
```bash
curl -X POST [http://localhost:8080/api/v1/auth/login](http://localhost:8080/api/v1/auth/login)
-H "Content-Type: application/json"
-d '{
"email": "[admin@kos.com](mailto:admin@kos.com)",
"password": "password123"
}'
```

### Get Dashboard Summary
```bash
curl -X GET [http://localhost:8080/api/v1/dashboard/summary](http://localhost:8080/api/v1/dashboard/summary)
-H "Authorization: Bearer <your-token>"
```

## 🧪 Testing

With Clean Architecture, testing becomes easier:

```go
// Test usecase tanpa database
func TestCreateRoom(t *testing.T) {
    mockRepo := &MockRoomRepository{}
    usecase := NewRoomUsecase(mockRepo)
    
    room := &entity.Room{
        RoomNumber: "A1",
        Price: 1000000,
    }
    
    err := usecase.Create(room)
    assert.NoError(t, err)
}
```

## 🔒 Security

- Passwords are hashed using bcrypt
- Authentication using JWT
- Token expires in 7 days
- Middleware for route protection
- Separation of concerns for the security layer

## 🎯 Development Roadmap

### Phase 1 - MVP ✅
- [x] Clean Architecture implementation
- [x] Basic CRUD operations
- [x] Authentication & Authorization
- [x] Dashboard summary
- [x] Payment tracking

### Phase 2 - Improvement
- [ ] Unit & Integration tests
- [ ] Export to PDF/Excel
- [ ] Advanced filtering & pagination
- [ ] Logging & monitoring

### Phase 3 - Automation
- [ ] WhatsApp reminder integration
- [ ] Payment gateway integration (Midtrans/Xendit)
- [ ] Email notifications

## 📖 Best Practices

1. **Dependency Rule**: Dependencies always point inward
2. **Interface Segregation**: Use interfaces for dependency inversion
3. **Single Responsibility**: Each layer has a clear responsibility
4. **Testability**: Easy to create mocks for testing
5. **Separation**: Business logic is separated from infrastructure

## 📄 License

MIT License