@echo off
REM MLM Application - Windows Installation Script
REM This script automates the setup process for Windows

echo ========================================
echo MLM Application - Automated Setup
echo ========================================
echo.

echo Checking prerequisites...
echo.

REM Check Go
where go >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Go is not installed
    echo Please install Go 1.21+ from https://golang.org/dl/
    pause
    exit /b 1
)
for /f "tokens=3" %%i in ('go version') do set GO_VERSION=%%i
echo [OK] Go %GO_VERSION%

REM Check Node.js
where node >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Node.js is not installed
    echo Please install Node.js 18+ from https://nodejs.org/
    pause
    exit /b 1
)
for /f "tokens=*" %%i in ('node --version') do set NODE_VERSION=%%i
echo [OK] Node.js %NODE_VERSION%

REM Check MySQL
where mysql >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo [WARNING] MySQL command not found
    echo Please ensure MySQL 8.0+ is installed and running
)

echo.
echo Installing dependencies...
echo.

REM Backend setup
echo Setting up Backend...
cd backend
if not exist ".env" (
    copy .env.example .env
    echo [OK] Created .env file
    echo [WARNING] Please edit backend\.env with your database credentials
) else (
    echo [WARNING] .env already exists, skipping
)

echo Downloading Go dependencies...
go mod download
go mod tidy
echo [OK] Backend dependencies installed

cd ..

REM E-commerce setup
echo.
echo Setting up E-commerce Frontend...
cd ecommerce
if not exist ".env.local" (
    copy .env.example .env.local
    echo [OK] Created .env.local file
)

echo Installing npm dependencies...
call npm install
echo [OK] E-commerce dependencies installed

cd ..

REM Admin Portal setup
echo.
echo Setting up Admin Portal...
cd admin-portal
if not exist ".env" (
    copy .env.example .env
    echo [OK] Created .env file
)

echo Installing npm dependencies...
call npm install
echo [OK] Admin Portal dependencies installed

cd ..

echo.
echo ========================================
echo Installation Complete!
echo ========================================
echo.
echo Next Steps:
echo.
echo 1. Configure your database:
echo    Edit backend\.env with your MySQL credentials
echo.
echo 2. Create the database:
echo    mysql -u root -p
echo    CREATE DATABASE mlm_app;
echo.
echo 3. Start the applications:
echo.
echo    Backend:
echo    cd backend ^&^& go run cmd\server\main.go
echo.
echo    E-commerce:
echo    cd ecommerce ^&^& npm run dev
echo.
echo    Admin Portal:
echo    cd admin-portal ^&^& npm run dev
echo.
echo 4. Access the applications:
echo    Backend API:     http://localhost:8080
echo    E-commerce:      http://localhost:3000
echo    Admin Portal:    http://localhost:3001
echo.
echo For more information, see:
echo    - QUICKSTART.md for quick setup
echo    - SETUP_GUIDE.md for detailed instructions
echo    - ARCHITECTURE.md for technical details
echo.
echo Happy coding!
echo.
pause
