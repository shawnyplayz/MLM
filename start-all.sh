#!/bin/bash

# MLM Application - Start All Services (Linux/Mac)
# This script starts all applications in separate terminal tabs/windows

echo "========================================"
echo "Starting MLM Application Services"
echo "========================================"
echo ""

# Detect OS
OS="$(uname -s)"

case "${OS}" in
    Linux*)
        # Linux - use gnome-terminal or xterm
        if command -v gnome-terminal &> /dev/null; then
            gnome-terminal --tab --title="Backend API" -- bash -c "cd backend && go run cmd/server/main.go; exec bash"
            gnome-terminal --tab --title="E-commerce" -- bash -c "cd ecommerce && npm run dev; exec bash"
            gnome-terminal --tab --title="Admin Dashboard" -- bash -c "cd admin-dashboard && npm run dev; exec bash"
            gnome-terminal --tab --title="Distributor Portal" -- bash -c "cd distributor-portal && npm run dev; exec bash"
        elif command -v xterm &> /dev/null; then
            xterm -T "Backend API" -e "cd backend && go run cmd/server/main.go" &
            xterm -T "E-commerce" -e "cd ecommerce && npm run dev" &
            xterm -T "Admin Dashboard" -e "cd admin-dashboard && npm run dev" &
            xterm -T "Distributor Portal" -e "cd distributor-portal && npm run dev" &
        else
            echo "No suitable terminal emulator found"
            echo "Please run each service manually in separate terminals"
            exit 1
        fi
        ;;
    Darwin*)
        # macOS - use Terminal.app
        osascript -e 'tell application "Terminal" to do script "cd '"$(pwd)"'/backend && go run cmd/server/main.go"'
        osascript -e 'tell application "Terminal" to do script "cd '"$(pwd)"'/ecommerce && npm run dev"'
        osascript -e 'tell application "Terminal" to do script "cd '"$(pwd)"'/admin-dashboard && npm run dev"'
        osascript -e 'tell application "Terminal" to do script "cd '"$(pwd)"'/distributor-portal && npm run dev"'
        ;;
    *)
        echo "Unsupported OS: ${OS}"
        echo "Please run each service manually"
        exit 1
        ;;
esac

echo ""
echo "========================================"
echo "All services are starting!"
echo "========================================"
echo ""
echo "Please wait a few moments for all services to initialize..."
echo ""
echo "Access URLs:"
echo "  Backend API:        http://localhost:8080"
echo "  E-commerce:         http://localhost:3000"
echo "  Admin Dashboard:    http://localhost:3001"
echo "  Distributor Portal: http://localhost:3002"
echo ""
echo "Press Ctrl+C in each terminal window to stop the services."
echo ""
