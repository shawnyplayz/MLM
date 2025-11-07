# MLM Application - Quick Start Guide

Get your MLM application up and running in minutes!

## Prerequisites Check

Before starting, ensure you have:
- ✅ Go 1.21+ installed
- ✅ Node.js 18+ installed
- ✅ MySQL 8.0+ installed and running
- ✅ Git installed

## Quick Setup (5 Minutes)

### Step 1: Database Setup (1 minute)

Open MySQL and run:

```sql
CREATE DATABASE mlm_app CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### Step 2: Backend Setup (2 minutes)

```bash
# Navigate to backend
cd backend

# Install dependencies
go mod download

# Create .env file
cp .env.example .env

# Edit .env with your MySQL credentials
# DB_USER=root
# DB_PASSWORD=your_password
# DB_NAME=mlm_app

# Run the server
go run cmd/server/main.go
```

✅ Backend running at http://localhost:8080

### Step 3: E-commerce Frontend (1 minute)

Open a new terminal:

```bash
# Navigate to ecommerce
cd ecommerce

# Install dependencies
npm install

# Create .env file
cp .env.example .env.local

# Run development server
npm run dev
```

✅ E-commerce running at http://localhost:3000

### Step 4: Admin Dashboard (1 minute)

Open a new terminal:

```bash
# Navigate to admin dashboard
cd admin-dashboard

# Install dependencies
npm install

# Create .env file
cp .env.example .env

# Run development server
npm run dev
```

✅ Admin Dashboard running at http://localhost:3001

### Step 5: Distributor Portal (Optional)

Open a new terminal:

```bash
# Navigate to distributor portal
cd distributor-portal

# Install dependencies
npm install

# Create .env file
cp .env.example .env

# Run development server
npm run dev
```

✅ Distributor Portal running at http://localhost:3002

## Test the Application

### 1. Register a Distributor

Open http://localhost:3000 and click "Join Now"

Or use curl:

```bash
curl -X POST http://localhost:8080/api/v1/distributors/register \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john@example.com",
    "password": "password123",
    "tree_type": "binary"
  }'
```

### 2. Login

```bash
curl -X POST http://localhost:8080/api/v1/distributors/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

Save the returned token!

### 3. Get Profile

```bash
curl -X GET http://localhost:8080/api/v1/distributors/profile \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

## Access Points

| Application | URL | Purpose |
|------------|-----|---------|
| Backend API | http://localhost:8080 | REST API |
| E-commerce | http://localhost:3000 | Public shopping site |
| Admin Dashboard | http://localhost:3001 | Admin management |
| Distributor Portal | http://localhost:3002 | Distributor dashboard |

## Default Credentials

After seeding, the following data is available:

### Ranks
- Bronze (Entry)
- Silver
- Gold
- Platinum
- Diamond

### Packages
- Starter ($99.99)
- Professional ($299.99)
- Elite ($999.99)

### Categories
- Health & Wellness
- Beauty & Personal Care
- Home & Living
- Nutrition

## Common Commands

### Backend

```bash
# Run server
go run cmd/server/main.go

# Build binary
go build -o mlm-api cmd/server/main.go

# Run tests
go test ./...

# Format code
go fmt ./...
```

### Frontend (All)

```bash
# Install dependencies
npm install

# Run development server
npm run dev

# Build for production
npm run build

# Run production build
npm start  # (Next.js)
npm run preview  # (Vite)
```

## Troubleshooting

### Backend won't start

**Error**: "Failed to connect to database"

**Solution**:
1. Check MySQL is running: `mysql -u root -p`
2. Verify database exists: `SHOW DATABASES;`
3. Check credentials in `.env`

### Frontend won't start

**Error**: "Cannot find module"

**Solution**:
```bash
rm -rf node_modules package-lock.json
npm install
```

### Port already in use

**Error**: "Port 8080 is already in use"

**Solution**:
- Change port in backend `.env`: `SERVER_PORT=8081`
- Or kill the process using the port

### CORS errors

**Solution**: Ensure backend `.env` has:
```env
CORS_ORIGINS=http://localhost:3000,http://localhost:3001,http://localhost:3002
```

## Next Steps

1. ✅ **Explore the UI**: Visit http://localhost:3000
2. ✅ **Create test accounts**: Register multiple distributors
3. ✅ **Test tree structures**: Add members with different tree types
4. ✅ **Check commissions**: Create orders and see commission calculations
5. ✅ **Review documentation**: Read ARCHITECTURE.md and FEATURES.md

## Development Workflow

### Making Changes

1. **Backend Changes**:
   - Edit Go files
   - Server auto-restarts (if using air/nodemon)
   - Or manually restart: `Ctrl+C` then `go run cmd/server/main.go`

2. **Frontend Changes**:
   - Edit React/Next.js files
   - Hot reload automatically updates browser
   - No restart needed

### Adding New Features

1. **Backend**:
   - Add model in `internal/domain/`
   - Add repository in `internal/repository/`
   - Add service in `internal/service/`
   - Add controller in `internal/controller/`
   - Register routes in `cmd/server/main.go`

2. **Frontend**:
   - Add component in `src/components/`
   - Add Redux slice in `src/features/`
   - Add API service in `src/services/`
   - Add route/page

## Production Checklist

Before deploying to production:

- [ ] Change JWT_SECRET to a strong random string
- [ ] Use strong database passwords
- [ ] Enable HTTPS
- [ ] Set GIN_MODE=release
- [ ] Configure proper CORS origins
- [ ] Set up database backups
- [ ] Configure logging
- [ ] Set up monitoring
- [ ] Add rate limiting
- [ ] Review security settings

## Getting Help

- 📖 **Full Documentation**: See README.md
- 🏗️ **Architecture**: See ARCHITECTURE.md
- ✨ **Features**: See FEATURES.md
- 🔧 **Detailed Setup**: See SETUP_GUIDE.md

## Summary

You now have a fully functional MLM application with:
- ✅ Backend API running
- ✅ E-commerce frontend
- ✅ Admin dashboard
- ✅ Distributor portal
- ✅ Database configured
- ✅ Sample data seeded

**Total setup time**: ~5 minutes

Start building your MLM business today! 🚀
