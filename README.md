# MLM Application - Full Stack

A comprehensive Multi-Level Marketing (MLM) platform built with modern technologies and best coding practices.

## Architecture

### Backend
- **Framework**: Gin (Go)
- **ORM**: GORM
- **Database**: MySQL
- **Architecture**: Clean Architecture (Repository → Service → Controller pattern)

### Frontend
- **E-commerce**: Next.js 14+ (App Router)
- **Admin Dashboard**: React 18+ with Redux Toolkit
- **Distributor Portal**: React 18+ with Redux Toolkit
- **Styling**: TailwindCSS
- **UI Components**: shadcn/ui
- **Icons**: Lucide React

## Features

### Distributor Management
- Complete distributor profiles with personal details
- Sponsor tracking
- Immediate and referral downline management
- Sales tracking and analytics
- Commission and bonus payout system
- Rank achievement tracking
- Package management
- Add new members directly from tree view

### MLM Tree Types
- **Binary Tree**: 2 legs per distributor
- **Matrix Tree**: Fixed width and depth (e.g., 3x9)
- **Unilevel Tree**: Unlimited width, single level
- **Breakaway Tree**: Groups break away at certain ranks
- **Hybrid Tree**: Combination of multiple tree types

## Project Structure

```
mlm-app/
├── backend/                 # Gin API Server
│   ├── cmd/
│   │   └── server/
│   ├── internal/
│   │   ├── config/
│   │   ├── domain/         # Domain models
│   │   ├── repository/     # Data access layer
│   │   ├── service/        # Business logic
│   │   ├── controller/     # HTTP handlers
│   │   ├── middleware/
│   │   └── utils/
│   ├── pkg/
│   └── migrations/
├── ecommerce/              # Next.js E-commerce
│   ├── src/
│   │   ├── app/
│   │   ├── components/
│   │   ├── lib/
│   │   └── styles/
│   └── public/
├── admin-dashboard/        # React Admin Dashboard
│   ├── src/
│   │   ├── components/
│   │   ├── features/       # Redux slices
│   │   ├── pages/
│   │   ├── services/       # API calls
│   │   ├── store/          # Redux store
│   │   └── utils/
│   └── public/
└── distributor-portal/     # React Distributor Portal
    ├── src/
    │   ├── components/
    │   ├── features/       # Redux slices
    │   ├── pages/
    │   ├── services/       # API calls
    │   ├── store/          # Redux store
    │   └── utils/
    └── public/
```

## Setup Instructions

### Prerequisites
- Go 1.21+
- Node.js 18+
- MySQL 8.0+
- npm or yarn

### Backend Setup

1. Navigate to backend directory:
```bash
cd backend
```

2. Install dependencies:
```bash
go mod download
```

3. Configure environment variables:
```bash
cp .env.example .env
# Edit .env with your database credentials
```

4. Run migrations:
```bash
go run cmd/server/main.go migrate
```

5. Start the server:
```bash
go run cmd/server/main.go
```

The API will be available at `http://localhost:8080`

### E-commerce Frontend Setup

1. Navigate to ecommerce directory:
```bash
cd ecommerce
```

2. Install dependencies:
```bash
npm install
```

3. Configure environment:
```bash
cp .env.example .env.local
```

4. Run development server:
```bash
npm run dev
```

Access at `http://localhost:3000`

### Admin Dashboard Setup

1. Navigate to admin-dashboard directory:
```bash
cd admin-dashboard
```

2. Install dependencies:
```bash
npm install
```

3. Configure environment:
```bash
cp .env.example .env
```

4. Run development server:
```bash
npm run dev
```

Access at `http://localhost:3001`

### Distributor Portal Setup

1. Navigate to distributor-portal directory:
```bash
cd distributor-portal
```

2. Install dependencies:
```bash
npm install
```

3. Configure environment:
```bash
cp .env.example .env
```

4. Run development server:
```bash
npm run dev
```

Access at `http://localhost:3002`

## API Documentation

API documentation is available at `http://localhost:8080/swagger` when the backend is running.

## Key Features Implementation

### Tree Management
- Dynamic tree visualization with D3.js/React Flow
- Drag-and-drop member addition
- Real-time tree updates
- Multi-tree type support

### Commission Engine
- Configurable commission rules
- Multiple bonus types
- Automated payout calculations
- Historical tracking

### Rank System
- Configurable rank criteria
- Automatic rank upgrades
- Rank achievement notifications
- Rank-based benefits

## Testing

### Backend Tests
```bash
cd backend
go test ./...
```

### Frontend Tests
```bash
cd [frontend-directory]
npm test
```

## Deployment

Refer to individual deployment guides:
- [Backend Deployment](./backend/DEPLOYMENT.md)
- [Frontend Deployment](./ecommerce/DEPLOYMENT.md)

## License

MIT License

## Support

For issues and questions, please open an issue on GitHub.
