# 🛒 Fiber E-commerce API

Authentication API with Role-based Access Control built with Go Fiber framework using Clean Architecture principles.

## 🏗️ Architecture

โปรเจ็กต์นี้ใช้หลักการ **Clean Architecture** มีโครงสร้างดังนี้:

```
fiber-ecommerce-api/
├── cmd/api/                    # Application entry point
│   └── main.go
├── internal/
│   ├── adapters/              # External adapters
│   │   ├── http/              # HTTP layer (handlers, middleware, routes)
│   │   │   ├── handlers/      # HTTP request handlers
│   │   │   ├── middleware/    # HTTP middleware
│   │   │   └── routes/        # Route definitions
│   │   └── persistence/       # Database layer
│   │       ├── models/        # Database models
│   │       └── repositories/  # Data access layer
│   ├── config/                # Configuration management
│   │   ├── config.go          # App configuration
│   │   └── database.go        # Database setup
│   └── core/                  # Business logic core
│       ├── domain/            # Domain entities and interfaces
│       │   ├── entities/      # Business entities
│       │   └── ports/         # Interfaces (ports)
│       └── services/          # Business logic services
├── pkg/utils/                 # Shared utilities
│   ├── jwt.go                 # JWT utilities
│   ├── password.go            # Password hashing
│   └── validator.go           # Validation utilities
├── docs/                      # API documentation (Swagger)
├── tmp/                       # Temporary build files
├── .env                       # Environment variables (local)
├── .env.example               # Environment variables template
├── .gitignore                 # Git ignore rules
├── .air.toml                  # Hot reload configuration
├── docker-compose.yml         # Docker services
├── go.mod                     # Go modules
└── go.sum                     # Go modules checksum
```

## 🚀 Features

- **🔐 User Authentication** (Register/Login)
- **🎫 JWT Token-based Authorization**
- **👥 Role-based Access Control** (Admin, User, Moderator)
- **🐘 PostgreSQL Database Integration**
- **✅ Input Validation**
- **🔒 Password Hashing** (bcrypt)
- **📚 Swagger API Documentation**
- **🐳 Docker Support**
- **🔥 Hot Reload Development** (Air)

## 📋 Prerequisites

- **Go** 1.24.2+
- **PostgreSQL** 15+
- **Docker & Docker Compose** (optional)

## 🛠️ Installation & Setup

### 1. Clone the repository
```bash
git clone <repository-url>
cd fiber-ecommerce-api
```

### 2. Install dependencies
```bash
go mod download
```

### 3. Install Swagger Generator
```bash
# Install swag for generating Swagger documentation
go install github.com/swaggo/swag/cmd/swag@latest

# Generate swagger docs
swag init -g cmd/api/main.go -o docs
```

### 4. Environment Configuration
คัดลอกไฟล์ `.env.example` เป็น `.env` และแก้ไขค่าต่างๆ ตามต้องการ:
```bash
cp .env.example .env
```

หรือสร้างไฟล์ `.env` ใหม่ด้วยเนื้อหาดังนี้:
```env
# 🌐 Environment
APP_ENV=development
APP_PORT=3000
APP_URL=http://localhost:3000

# 📦 Database (PostgreSQL)
DB_HOST=localhost
DB_PORT=5432
DB_NAME=fiberecomapidb
DB_USER=postgres
DB_PASS=123456
DB_SSL=disable

# 🔐 JWT Config
JWT_SECRET=fibernextcommerce_jwt_secret_key_2024
JWT_EXPIRES_IN=24h
```

### 5. Database Setup

#### Option A: Using Docker (แนะนำ)
```bash
docker-compose up -d
```

#### Option B: Manual PostgreSQL Setup
สร้าง PostgreSQL database ด้วยข้อมูลใน `.env` file

### 6. Run the Application

#### Development (with hot reload)
```bash
# Install Air for hot reloading
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

#### Production
```bash
go run cmd/api/main.go
```

🚀 **API จะรันที่**: `http://localhost:3000`

## 📚 API Documentation

เข้าถึง Swagger UI documentation ได้ที่: **`http://localhost:3000/swagger/`**

### 🛣️ Available Endpoints

#### 🔐 Authentication
- `POST /api/auth/register` - สมัครสมาชิกใหม่
- `POST /api/auth/login` - เข้าสู่ระบบ

#### 👤 User (Protected Routes)
- `GET /api/user/profile` - ดูข้อมูลโปรไฟล์ผู้ใช้

#### 👑 Admin (Admin only)
- `GET /api/admin/dashboard` - หน้า Admin dashboard

## 🔐 Authentication Flow

1. **Register**: สร้างบัญชีผู้ใช้ใหม่ด้วย email, password, first name, และ last name
2. **Login**: เข้าสู่ระบบด้วย email และ password เพื่อรับ JWT token
3. **Protected Routes**: ใส่ JWT token ใน `Authorization` header เป็น `Bearer <token>`

### 💡 Example Requests

#### 📝 Register
```bash
curl -X POST http://localhost:3000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123",
    "first_name": "John",
    "last_name": "Doe"
  }'
```

#### 🔑 Login
```bash
curl -X POST http://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

#### 👤 Get Profile (Protected)
```bash
curl -X GET http://localhost:3000/api/user/profile \
  -H "Authorization: Bearer <your-jwt-token>"
```

## 🏛️ Project Structure Details

### 🎯 Core Domain
- `entities.User` - User entity พร้อม roles
- `entities.LoginRequest` - โครงสร้าง Login request
- `entities.RegisterRequest` - โครงสร้าง Registration request

### ⚙️ Services
- `services.AuthService` - Authentication business logic

### 🔧 Utilities
- `utils.ValidateStruct` - ตรวจสอบความถูกต้องของ struct
- `utils.HashPassword` - เข้ารหัสรหัสผ่าน
- `utils.GenerateJWT` - สร้าง JWT token

### ⚙️ Configuration
- `config.LoadConfig` - โหลด environment configuration
- `config.SetupDatabase` - ตั้งค่าการเชื่อมต่อฐานข้อมูล

## 🔧 Development Tools

### 🔥 Hot Reload
โปรเจ็กต์ใช้ [Air](https://github.com/cosmtrek/air) สำหรับ hot reloading ระหว่างการพัฒนา ดูการตั้งค่าได้ที่ `.air.toml`

### 🗄️ Database Migration
Auto-migration จัดการโดย GORM และจะรันโดยอัตโนมัติเมื่อเริ่ม application

## 🛠️ Tech Stack

### Backend Framework
- **Fiber v2** - Fast HTTP web framework
- **GORM** - ORM สำหรับ Go
- **PostgreSQL** - ฐานข้อมูลเชิงสัมพันธ์

### Authentication & Security
- **JWT** - JSON Web Tokens
- **bcrypt** - Password hashing
- **Validator** - Input validation

### Development Tools
- **Air** - Hot reload สำหรับ Go
- **Swagger** - API documentation
- **Docker** - Containerization

## 🐳 Docker Support

ใช้ `docker-compose.yml` ที่มีให้เพื่อรัน PostgreSQL:

```bash
# Start PostgreSQL only
docker-compose up -d postgres

# View logs
docker-compose logs -f postgres

# Stop services
docker-compose down
```

## 🧪 Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

## 📝 API Response Format

### ✅ Success Response
```json
{
  "success": true,
  "message": "Operation successful",
  "data": {
    // response data
  }
}
```

### ❌ Error Response
```json
{
  "success": false,
  "error": "Error message",
  "details": {
    // error details
  }
}
```

## 🚀 Deployment

### Production Build
```bash
# Build for production
go build -o bin/api cmd/api/main.go

# Run production build
./bin/api
```

### Environment Variables for Production
อย่าลืมตั้งค่า environment variables สำหรับ production:
- เปลี่ยน `JWT_SECRET` เป็นค่าที่ปลอดภัยกว่า
- ตั้งค่า database credentials ที่ถูกต้อง
- เปลี่ยน `APP_ENV` เป็น `production`

## 🤝 Contributing

1. ทำตามหลักการ Clean Architecture
2. รักษาโครงสร้างโปรเจ็กต์ที่มีอยู่
3. เพิ่ม error handling และ validation ที่เหมาะสม
4. อัปเดต documentation สำหรับ features ใหม่
5. เขียน tests สำหรับ functionality ใหม่

## 📄 License

This project is licensed under the Apache 2.0 License.

## 👥 Author

- **Samit Koyom** - [iamsamitdev](https://github.com/iamsamitdev)

---

🎉 **Happy Coding!** หากมีคำถามหรือปัญหาใดๆ สามารถสร้าง issue ได้เลยครับ
