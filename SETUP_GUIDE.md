# MLM Application - Complete Setup Guide

This guide will walk you through setting up the entire MLM application stack.

## Prerequisites

Ensure you have the following installed:
- **Go** 1.21 or higher
- **Node.js** 18 or higher
- **MySQL** 8.0 or higher
- **Git**

## Project Structure

```
MLM/
├── backend/              # Gin API Server (Go)
├── ecommerce/           # Next.js E-commerce Frontend
├── admin-dashboard/     # React Admin Dashboard
├── distributor-portal/  # React Distributor Portal
└── README.md
```

## Step 1: Database Setup

### 1.1 Create MySQL Database

```sql
CREATE DATABASE mlm_app CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'mlm_user'@'localhost' IDENTIFIED BY 'your_secure_password';
GRANT ALL PRIVILEGES ON mlm_app.* TO 'mlm_user'@'localhost';
FLUSH PRIVILEGES;
```

## Step 2: Backend Setup (Gin + GORM)

### 2.1 Navigate to backend directory

```bash
cd backend
```

### 2.2 Install Go dependencies

```bash
go mod download
go mod tidy
```

### 2.3 Configure environment variables

```bash
cp .env.example .env
```

Edit `.env` file with your database credentials:

```env
SERVER_PORT=8080
GIN_MODE=debug

DB_HOST=localhost
DB_PORT=3306
DB_USER=mlm_user
DB_PASSWORD=your_secure_password
DB_NAME=mlm_app

JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
JWT_EXPIRY=24h

CORS_ORIGINS=http://localhost:3000,http://localhost:3001,http://localhost:3002
```

### 2.4 Run the backend server

```bash
go run cmd/server/main.go
```

The API will be available at `http://localhost:8080`

### 2.5 Verify backend is running

```bash
curl http://localhost:8080/health
```

Expected response: `{"status":"ok"}`

## Step 3: E-commerce Frontend Setup (Next.js)

### 3.1 Navigate to ecommerce directory

```bash
cd ../ecommerce
```

### 3.2 Install dependencies

```bash
npm install
```

### 3.3 Configure environment

```bash
cp .env.example .env.local
```

Edit `.env.local`:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
```

### 3.4 Run development server

```bash
npm run dev
```

Access at `http://localhost:3000`

## Step 4: Admin Dashboard Setup (React + Redux)

### 4.1 Navigate to admin-dashboard directory

```bash
cd ../admin-dashboard
```

### 4.2 Install dependencies

```bash
npm install
```

### 4.3 Configure environment

```bash
cp .env.example .env
```

Edit `.env`:

```env
VITE_API_URL=http://localhost:8080/api/v1
```

### 4.4 Run development server

```bash
npm run dev
```

Access at `http://localhost:3001`

## Step 5: Distributor Portal Setup (React + Redux)

### 5.1 Navigate to distributor-portal directory

```bash
cd ../distributor-portal
```

### 5.2 Install dependencies

```bash
npm install
```

### 5.3 Configure environment

```bash
cp .env.example .env
```

Edit `.env`:

```env
VITE_API_URL=http://localhost:8080/api/v1
```

### 5.4 Run development server

```bash
npm run dev
```

Access at `http://localhost:3002`

## Step 6: Testing the Application

### 6.1 Create a test distributor account

Use the e-commerce frontend or make a direct API call:

```bash
curl -X POST http://localhost:8080/api/v1/distributors/register \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john@example.com",
    "password": "password123",
    "phone": "1234567890",
    "tree_type": "binary"
  }'
```

### 6.2 Login

```bash
curl -X POST http://localhost:8080/api/v1/distributors/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

Save the returned `token` for authenticated requests.

### 6.3 Get distributor profile

```bash
curl -X GET http://localhost:8080/api/v1/distributors/profile \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

## Running All Services

You can run all services simultaneously using separate terminal windows:

### Terminal 1 - Backend
```bash
cd backend
go run cmd/server/main.go
```

### Terminal 2 - E-commerce
```bash
cd ecommerce
npm run dev
```

### Terminal 3 - Admin Dashboard
```bash
cd admin-dashboard
npm run dev
```

### Terminal 4 - Distributor Portal
```bash
cd distributor-portal
npm run dev
```

## Application URLs

- **Backend API**: http://localhost:8080
- **E-commerce**: http://localhost:3000
- **Admin Dashboard**: http://localhost:3001
- **Distributor Portal**: http://localhost:3002

## Key Features

### MLM Tree Types Supported

1. **Binary Tree** - 2 legs per distributor (left/right)
2. **Matrix Tree** - Fixed width and depth (e.g., 3x9)
3. **Unilevel Tree** - Unlimited width, single level
4. **Breakaway Tree** - Groups break away at certain ranks
5. **Hybrid Tree** - Combination of multiple tree types

### Commission System

- **Direct Referral Commission** - Earn from direct recruits
- **Level Commissions** - Earn from multiple levels deep
- **Rank Bonuses** - Monthly bonuses based on rank
- **Team Sales Commissions** - Percentage of team sales

### Rank System

- Bronze (Entry Level)
- Silver (Intermediate)
- Gold (Advanced)
- Platinum (Elite)
- Diamond (Top Tier)

Each rank has specific requirements:
- Minimum personal sales
- Minimum team sales
- Minimum downlines
- Minimum active downlines

## Troubleshooting

### Backend won't start

1. Check MySQL is running:
   ```bash
   mysql -u root -p
   ```

2. Verify database exists:
   ```sql
   SHOW DATABASES;
   ```

3. Check Go version:
   ```bash
   go version
   ```

### Frontend won't start

1. Clear node_modules and reinstall:
   ```bash
   rm -rf node_modules package-lock.json
   npm install
   ```

2. Check Node version:
   ```bash
   node --version
   ```

### CORS errors

Ensure backend `.env` has correct CORS_ORIGINS:
```env
CORS_ORIGINS=http://localhost:3000,http://localhost:3001,http://localhost:3002
```

### Database connection errors

1. Verify MySQL credentials in backend `.env`
2. Check MySQL is running on port 3306
3. Ensure user has proper permissions

## Production Deployment

### Backend

1. Build the binary:
   ```bash
   cd backend
   go build -o mlm-api cmd/server/main.go
   ```

2. Set production environment variables
3. Run with process manager (PM2, systemd, etc.)

### Frontend Applications

1. Build for production:
   ```bash
   npm run build
   ```

2. Deploy to hosting service (Vercel, Netlify, etc.)
3. Update environment variables with production API URL

## Database Migrations

The application automatically runs migrations on startup. To manually run migrations:

```bash
cd backend
go run cmd/server/main.go migrate
```

## Seeding Data

Initial data (ranks, packages, categories) is automatically seeded on first run.

## API Documentation

Once the backend is running, API documentation is available at:
- Swagger UI: `http://localhost:8080/swagger` (if configured)

## Support

For issues or questions:
1. Check the troubleshooting section
2. Review logs in each application
3. Open an issue on GitHub

## Security Notes

- Change JWT_SECRET in production
- Use strong database passwords
- Enable HTTPS in production
- Implement rate limiting
- Regular security audits

## Next Steps

1. Customize the UI/UX to match your brand
2. Add payment gateway integration
3. Implement email notifications
4. Add SMS notifications
5. Create mobile apps
6. Add analytics and reporting
7. Implement backup strategies

## License

MIT License - See LICENSE file for details
