#!/bin/bash

# MLM Application - Installation Script
# This script automates the setup process

echo "🚀 MLM Application - Automated Setup"
echo "===================================="
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check prerequisites
echo "📋 Checking prerequisites..."

# Check Go
if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go is not installed${NC}"
    echo "Please install Go 1.21+ from https://golang.org/dl/"
    exit 1
fi
echo -e "${GREEN}✅ Go $(go version | awk '{print $3}')${NC}"

# Check Node.js
if ! command -v node &> /dev/null; then
    echo -e "${RED}❌ Node.js is not installed${NC}"
    echo "Please install Node.js 18+ from https://nodejs.org/"
    exit 1
fi
echo -e "${GREEN}✅ Node.js $(node --version)${NC}"

# Check MySQL
if ! command -v mysql &> /dev/null; then
    echo -e "${YELLOW}⚠️  MySQL command not found${NC}"
    echo "Please ensure MySQL 8.0+ is installed and running"
fi

echo ""
echo "📦 Installing dependencies..."
echo ""

# Backend setup
echo "🔧 Setting up Backend..."
cd backend || exit

if [ ! -f ".env" ]; then
    cp .env.example .env
    echo -e "${GREEN}✅ Created .env file${NC}"
    echo -e "${YELLOW}⚠️  Please edit backend/.env with your database credentials${NC}"
else
    echo -e "${YELLOW}⚠️  .env already exists, skipping${NC}"
fi

echo "📥 Downloading Go dependencies..."
go mod download
go mod tidy
echo -e "${GREEN}✅ Backend dependencies installed${NC}"

cd ..

# E-commerce setup
echo ""
echo "🛒 Setting up E-commerce Frontend..."
cd ecommerce || exit

if [ ! -f ".env.local" ]; then
    cp .env.example .env.local
    echo -e "${GREEN}✅ Created .env.local file${NC}"
fi

echo "📥 Installing npm dependencies..."
npm install
echo -e "${GREEN}✅ E-commerce dependencies installed${NC}"

cd ..

# Admin Dashboard setup
echo ""
echo "👨‍💼 Setting up Admin Dashboard..."
cd admin-dashboard || exit

if [ ! -f ".env" ]; then
    cp .env.example .env
    echo -e "${GREEN}✅ Created .env file${NC}"
fi

echo "📥 Installing npm dependencies..."
npm install
echo -e "${GREEN}✅ Admin Dashboard dependencies installed${NC}"

cd ..

# Distributor Portal setup
echo ""
echo "👥 Setting up Distributor Portal..."
cd distributor-portal || exit

if [ ! -f ".env" ]; then
    cp .env.example .env
    echo -e "${GREEN}✅ Created .env file${NC}"
fi

echo "📥 Installing npm dependencies..."
npm install
echo -e "${GREEN}✅ Distributor Portal dependencies installed${NC}"

cd ..

echo ""
echo "✨ Installation Complete!"
echo "======================="
echo ""
echo "📝 Next Steps:"
echo ""
echo "1. Configure your database:"
echo "   ${YELLOW}Edit backend/.env with your MySQL credentials${NC}"
echo ""
echo "2. Create the database:"
echo "   ${YELLOW}mysql -u root -p${NC}"
echo "   ${YELLOW}CREATE DATABASE mlm_app;${NC}"
echo ""
echo "3. Start the applications:"
echo ""
echo "   Backend:"
echo "   ${GREEN}cd backend && go run cmd/server/main.go${NC}"
echo ""
echo "   E-commerce:"
echo "   ${GREEN}cd ecommerce && npm run dev${NC}"
echo ""
echo "   Admin Dashboard:"
echo "   ${GREEN}cd admin-dashboard && npm run dev${NC}"
echo ""
echo "   Distributor Portal:"
echo "   ${GREEN}cd distributor-portal && npm run dev${NC}"
echo ""
echo "4. Access the applications:"
echo "   Backend API:        ${GREEN}http://localhost:8080${NC}"
echo "   E-commerce:         ${GREEN}http://localhost:3000${NC}"
echo "   Admin Dashboard:    ${GREEN}http://localhost:3001${NC}"
echo "   Distributor Portal: ${GREEN}http://localhost:3002${NC}"
echo ""
echo "📖 For more information, see:"
echo "   - QUICKSTART.md for quick setup"
echo "   - SETUP_GUIDE.md for detailed instructions"
echo "   - ARCHITECTURE.md for technical details"
echo ""
echo "🎉 Happy coding!"
