@echo off
REM MLM Application - Start All Services (Windows)
REM This script starts all applications in separate windows

echo ========================================
echo Starting MLM Application Services
echo ========================================
echo.

REM Start Backend
echo Starting Backend API Server...
start "MLM Backend API" cmd /k "cd backend && go run cmd/server/main.go"
timeout /t 2 /nobreak >nul

REM Start E-commerce
echo Starting E-commerce Frontend...
start "MLM E-commerce" cmd /k "cd ecommerce && npm run dev"
timeout /t 2 /nobreak >nul

REM Start Admin Portal
echo Starting Admin Portal...
start "MLM Admin Portal" cmd /k "cd admin-portal && npm run dev"
timeout /t 2 /nobreak >nul

echo.
echo ========================================
echo All services are starting!
echo ========================================
echo.
echo Please wait a few moments for all services to initialize...
echo.
echo Access URLs:
echo   Backend API:     http://localhost:8080
echo   E-commerce:      http://localhost:3000
echo   Admin Portal:    http://localhost:3001
echo.
echo Press any key to open the E-commerce site in your browser...
pause >nul

start http://localhost:3000

echo.
echo All services are running in separate windows.
echo Close those windows to stop the services.
echo.
pause
