# 🛒 Fiber E-commerce API

Authentication API with Role-based Access Control built with Go Fiber framework using Clean Architecture principles.

## 🏗️ Architecture

โปรเจ็กต์นี้ใช้หลักการ **Clean Architecture** มีโครงสร้างดังนี้:

```
fiber-ecommerce-api/
├── cmd/
│   ├── api/                   # Application entry point
│   │   └── main.go
│   └── migrate/               # Database migration CLI
│       └── main.go
├── internal/
│   ├── adapters/              # External adapters
│   │   ├── http/              # HTTP layer (handlers, middleware, routes)
│   │   │   ├── handlers/      # HTTP request handlers
│   │   │   ├── middleware/    # HTTP middleware
│   │   │   └── routes/        # Route definitions
│   │   └── persistence/       # Database layer
│   │       ├── models/        # Database models (GORM)
│   │       └── repositories/  # Data access layer
│   ├── config/                # Configuration management
│   │   ├── config.go          # App configuration
│   │   ├── database.go        # Database setup & migration
│   │   └── seeder.go          # Database seeding (Admin user)
│   └── core/                  # Business logic core
│       ├── domain/            # Domain entities and interfaces
│       │   ├── entities/      # Business entities
│       │   └── ports/         # Interfaces (ports)
│       │       ├── repositories/  # Repository interfaces
│       │       └── services/      # Service interfaces
│       └── services/          # Business logic services
├── pkg/utils/                 # Shared utilities
│   ├── jwt.go                 # JWT utilities
│   ├── password.go            # Password hashing
│   └── validator.go           # Validation utilities
├── docs/                      # API documentation (Swagger)
├── scripts/                   # Utility scripts
├── tmp/                       # Temporary build files
├── .env                       # Environment variables (local)
├── .env.example               # Environment variables template
├── .gitignore                 # Git ignore rules
├── .air.toml                  # Hot reload configuration
├── docker-compose.yml         # Docker services
├── Makefile                   # Build commands
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
- **🌱 Database Seeding** (Auto Admin User Creation)
- **🔄 Database Migration Control**
- **🐳 Docker Support**
- **🔥 Hot Reload Development** (Air)

## 📋 Prerequisites

- **Go** 1.21+
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

### 3. Install Development Tools
```bash
# Install Air for hot reloading
go install github.com/cosmtrek/air@latest

# Install Swagger generator (optional)
go install github.com/swaggo/swag/cmd/swag@latest
```

### 4. Environment Configuration
คัดลอกไฟล์ `.env.example` เป็น `.env` และแก้ไขค่าต่างๆ ตามต้องการ:
```bash
cp .env.example .env
```

ตัวอย่างไฟล์ `.env`:
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
JWT_SECRET=your-super-secret-jwt-key-at-least-32-characters-long
JWT_EXPIRES_IN=24h

# 🔄 Database Migration
AUTO_MIGRATE=true

# 👑 Admin User Seeding (Optional)
ADMIN_EMAIL=admin@company.com
ADMIN_PASSWORD=SecurePassword123!
ADMIN_FIRST_NAME=System
ADMIN_LAST_NAME=Administrator
```

### 5. Database Setup

#### Option A: Using Docker (แนะนำ)
```bash
# Start PostgreSQL
docker-compose up -d postgres

# View logs
docker-compose logs -f postgres
```

#### Option B: Manual PostgreSQL Setup
สร้าง PostgreSQL database ด้วยข้อมูลใน `.env` file

### 6. Run the Application

#### Development (with hot reload)
```bash
# Run with hot reload
air

# หรือใช้ make command
make dev
```

#### Production
```bash
# Build และรัน
make build && make run

# หรือรันโดยตรง
go run cmd/api/main.go
```

🚀 **API จะรันที่**: `http://localhost:3000`

## 🗄️ Database Management

### Migration Control
```bash
# Auto-migration (development)
AUTO_MIGRATE=true go run cmd/api/main.go

# Disable auto-migration (production)
AUTO_MIGRATE=false go run cmd/api/main.go

# Manual migration
go run cmd/migrate/main.go

# หรือใช้ make commands
make migrate-up      # Manual migration
make prod-migrate    # Production migration
```

### Admin User Seeding
ระบบจะสร้าง admin user อัตโนมัติเมื่อ:
- มีการตั้งค่า `ADMIN_EMAIL`, `ADMIN_PASSWORD`, `ADMIN_FIRST_NAME`, `ADMIN_LAST_NAME` ใน `.env`
- ยังไม่มี admin user ในระบบ

**Log Messages:**
```bash
# เมื่อสร้าง admin สำเร็จ
✅ Admin user created successfully: admin@company.com
👤 Name: System Administrator
⚠️  Please ensure you're using a secure password!

# เมื่อไม่มี admin credentials
⚠️  ADMIN_EMAIL not set, skipping admin user seeding
💡 To create admin user, set ADMIN_EMAIL, ADMIN_PASSWORD, ADMIN_FIRST_NAME, ADMIN_LAST_NAME in .env
```

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
- `POST /api/admin/register` - สร้างผู้ใช้ใหม่ (admin เท่านั้น)

## 🔐 Authentication Flow

1. **Register**: สร้างบัญชีผู้ใช้ใหม่ด้วย email, password, first name, และ last name (role = "user")
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

#### 👑 Admin Dashboard (Admin only)
```bash
curl -X GET http://localhost:3000/api/admin/dashboard \
  -H "Authorization: Bearer <admin-jwt-token>"
```

#### 👑 Admin Register User (Admin only)
```bash
curl -X POST http://localhost:3000/api/admin/register \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <admin-jwt-token>" \
  -d '{
    "email": "newuser@example.com",
    "password": "password123",
    "first_name": "New",
    "last_name": "User",
    "role": "user"
  }'
```

## 🏛️ Project Structure Details

### 🎯 Core Domain
- `entities.User` - User entity พร้อม roles (admin, user, moderator)
- `entities.LoginRequest` - โครงสร้าง Login request
- `entities.RegisterRequest` - โครงสร้าง Registration request
- `entities.Role` - User role enumeration

### ⚙️ Services
- `services.AuthService` - Authentication business logic
- `services.AuthServiceImpl` - Authentication service implementation

### 🗄️ Repositories
- `repositories.UserRepository` - User data access interface
- `repositories.UserRepositoryImpl` - GORM implementation

### 🔧 Utilities
- `utils.ValidateStruct` - ตรวจสอบความถูกต้องของ struct
- `utils.HashPassword` - เข้ารหัสรหัสผ่าน
- `utils.CheckPassword` - ตรวจสอบรหัสผ่าน
- `utils.GenerateJWT` - สร้าง JWT token
- `utils.ParseJWT` - แปลง JWT token

### ⚙️ Configuration
- `config.LoadConfig` - โหลด environment configuration
- `config.SetupDatabase` - ตั้งค่าการเชื่อมต่อฐานข้อมูล
- `config.SeedAdminUser` - สร้าง admin user อัตโนมัติ

### 🛡️ Middleware
- `middleware.AuthMiddleware` - JWT authentication middleware
- `middleware.RequireRole` - Role-based authorization middleware

## 🔧 Development Tools

### 🔥 Hot Reload
โปรเจ็กต์ใช้ [Air](https://github.com/cosmtrek/air) สำหรับ hot reloading ระหว่างการพัฒนา ดูการตั้งค่าได้ที่ `.air.toml`

### 🛠️ Make Commands
```bash
make help          # แสดงคำสั่งที่ใช้ได้ทั้งหมด
make build         # Build แอปพลิเคชัน
make run           # รันแอปพลิเคชัน
make dev           # รันแบบ development mode (hot reload)
make test          # รัน tests
make clean         # ลบไฟล์ build
make deps          # ดาวน์โหลด dependencies
make migrate-up    # รัน database migration
make prod-migrate  # รัน production migration
```

### 🗄️ Database Migration
- **Development**: Auto-migration เปิดใช้งานโดยค่าเริ่มต้น
- **Production**: ต้องตั้งค่า `AUTO_MIGRATE=true` หรือใช้ manual migration
- **Manual Control**: ใช้ `cmd/migrate/main.go` หรือ make commands

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
- **Make** - Build automation

## 🐳 Docker Support

ใช้ `docker-compose.yml` ที่มีให้เพื่อรัน PostgreSQL:

```bash
# Start PostgreSQL only
docker-compose up -d postgres

# View logs
docker-compose logs -f postgres

# Stop services
docker-compose down

# Remove volumes (reset database)
docker-compose down -v
```

## 🧪 Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# หรือใช้ make command
make test
```

## 📝 API Response Format

### ✅ Success Response
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "email": "user@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "role": "user",
    "is_active": true,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### ❌ Error Response
```json
{
  "error": "Invalid email or password"
}
```

## 🚀 Deployment

### Production Build
```bash
# Build for production
make build

# Run production build
./bin/api
```

### Environment Variables for Production
สำคัญสำหรับ production:

```env
# ต้องตั้งค่าให้ปลอดภัย
APP_ENV=production
JWT_SECRET=your-super-secure-jwt-secret-at-least-32-characters-long
DB_PASS=your-secure-database-password

# Database SSL
DB_SSL=require

# Migration control
AUTO_MIGRATE=false

# Admin credentials (ต้องตั้งค่า)
ADMIN_EMAIL=admin@yourcompany.com
ADMIN_PASSWORD=YourSecureAdminPassword123!
ADMIN_FIRST_NAME=System
ADMIN_LAST_NAME=Administrator
```

### Production Checklist
- [ ] เปลี่ยน `JWT_SECRET` เป็นค่าที่ปลอดภัย (อย่างน้อย 32 ตัวอักษร)
- [ ] ตั้งค่า database credentials ที่ถูกต้อง
- [ ] เปลี่ยน `APP_ENV` เป็น `production`
- [ ] เปิดใช้งาน SSL สำหรับ database (`DB_SSL=require`)
- [ ] ตั้งค่า admin credentials ที่ปลอดภัย
- [ ] ปิด auto-migration (`AUTO_MIGRATE=false`)
- [ ] รัน manual migration ก่อน deploy

## 🔒 Security Features

### Password Security
- **bcrypt hashing** - รหัสผ่านถูกเข้ารหัสด้วย bcrypt
- **Minimum length** - รหัสผ่านต้องมีอย่างน้อย 6 ตัวอักษร
- **Admin password** - รหัสผ่าน admin ต้องมีอย่างน้อย 8 ตัวอักษร

### JWT Security
- **Secure secret** - JWT secret ต้องมีอย่างน้อย 32 ตัวอักษรใน production
- **Token expiration** - Token หมดอายุใน 24 ชั่วโมง (ปรับได้)
- **Role-based access** - ตรวจสอบ role ในทุก protected route

### Database Security
- **SSL connection** - แนะนำให้ใช้ SSL ใน production
- **Environment variables** - ข้อมูลสำคัญเก็บใน environment variables
- **No default passwords** - ไม่มีรหัสผ่าน default ใน production

## 🤝 Contributing

1. ทำตามหลักการ Clean Architecture
2. รักษาโครงสร้างโปรเจ็กต์ที่มีอยู่
3. เพิ่ม error handling และ validation ที่เหมาะสม
4. อัปเดต documentation สำหรับ features ใหม่
5. เขียน tests สำหรับ functionality ใหม่
6. ใช้ make commands สำหรับ development

## 📄 License

This project is licensed under the Apache 2.0 License.

## 👥 Author

- **Samit Koyom** - [iamsamitdev](https://github.com/iamsamitdev)

---

🎉 **Happy Coding!** หากมีคำถามหรือปัญหาใดๆ สามารถสร้าง issue ได้เลยครับ
