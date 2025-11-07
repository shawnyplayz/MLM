# MLM Application - Architecture Documentation

## Overview

This is a comprehensive Multi-Level Marketing (MLM) platform built with modern, scalable technologies following best coding practices and clean architecture principles.

## Technology Stack

### Backend
- **Language**: Go 1.21+
- **Framework**: Gin (HTTP web framework)
- **ORM**: GORM
- **Database**: MySQL 8.0+
- **Authentication**: JWT (JSON Web Tokens)
- **Architecture Pattern**: Clean Architecture (Repository → Service → Controller)

### Frontend
- **E-commerce**: Next.js 14+ (React framework with App Router)
- **Admin Dashboard**: React 18+ with Vite
- **Distributor Portal**: React 18+ with Vite
- **State Management**: Redux Toolkit
- **Styling**: TailwindCSS
- **UI Components**: Radix UI + shadcn/ui
- **Icons**: Lucide React
- **HTTP Client**: Axios

## Architecture Principles

### 1. Clean Architecture (Backend)

```
┌─────────────────────────────────────────┐
│           Controllers                    │
│  (HTTP Handlers, Request/Response)      │
└──────────────┬──────────────────────────┘
               │
┌──────────────▼──────────────────────────┐
│            Services                      │
│     (Business Logic Layer)              │
└──────────────┬──────────────────────────┘
               │
┌──────────────▼──────────────────────────┐
│          Repositories                    │
│      (Data Access Layer)                │
└──────────────┬──────────────────────────┘
               │
┌──────────────▼──────────────────────────┐
│           Database                       │
│         (MySQL/GORM)                    │
└─────────────────────────────────────────┘
```

**Benefits:**
- Separation of concerns
- Testability
- Maintainability
- Scalability
- Independence from frameworks

### 2. Modular Design

Each component is highly modular and independent:

#### Backend Modules
- **Config**: Configuration management
- **Domain**: Business entities and models
- **Repository**: Data access interfaces and implementations
- **Service**: Business logic and orchestration
- **Controller**: HTTP request handling
- **Middleware**: Authentication, CORS, logging
- **Utils**: Helper functions and utilities

#### Frontend Modules
- **Features**: Redux slices (state management)
- **Services**: API communication layer
- **Components**: Reusable UI components
- **Pages**: Route-specific views
- **Store**: Redux store configuration
- **Utils**: Helper functions

### 3. Dependency Injection

The backend uses constructor-based dependency injection:

```go
// Repository layer
distributorRepo := repository.NewDistributorRepository(db)

// Service layer (depends on repository)
treeService := service.NewTreeService(distributorRepo)
distributorService := service.NewDistributorService(
    distributorRepo, 
    rankRepo, 
    treeService,
)

// Controller layer (depends on service)
distributorController := controller.NewDistributorController(
    distributorService, 
    cfg,
)
```

## Core Features

### 1. Distributor Management

**Models:**
- Personal information (name, email, phone, address)
- Authentication (password hashing with bcrypt)
- MLM structure (sponsor, tree type, position, level)
- Business metrics (sales, commissions, bonuses)
- Status and rank tracking

**Operations:**
- Registration with sponsor assignment
- Login with JWT token generation
- Profile management
- Tree position calculation
- Rank eligibility checking

### 2. MLM Tree Structures

#### Binary Tree
- 2 positions per distributor (left/right)
- Spillover mechanism
- Balanced tree growth

```go
func (s *treeService) findBinaryPosition(sponsorID uint) (string, error) {
    // Check left position
    // Check right position
    // Recursive spillover if both filled
}
```

#### Matrix Tree
- Fixed width (e.g., 3 wide)
- Fixed depth (e.g., 9 levels)
- Position format: "pos_1", "pos_2", "pos_3"

#### Unilevel Tree
- Unlimited width
- Single level depth
- Direct downlines only

#### Breakaway Tree
- Similar to unilevel
- Groups "break away" at certain ranks
- Independent team management

#### Hybrid Tree
- Combination of multiple tree types
- Flexible structure
- Configurable rules

### 3. Commission System

**Commission Types:**

1. **Direct Referral Commission**
   - Earned from direct recruits
   - Configurable percentage
   - Based on commissionable value

2. **Level Commissions**
   - Multi-level earning (up to configured depth)
   - Decreasing percentage per level
   - Upline chain calculation

3. **Rank Bonuses**
   - Monthly bonuses based on rank
   - Achievement rewards
   - Percentage bonuses on commissions

**Commission Calculation Flow:**

```
Order Placed
    ↓
Calculate Commissionable Value
    ↓
Direct Referral Commission (Level 1)
    ↓
Level Commissions (Level 2-N)
    ↓
Apply Rank Bonuses
    ↓
Create Commission Records
    ↓
Update Distributor Totals
```

### 4. Rank System

**Rank Progression:**

```
Bronze → Silver → Gold → Platinum → Diamond
```

**Requirements per Rank:**
- Minimum personal sales
- Minimum team sales
- Minimum total downlines
- Minimum active downlines

**Benefits:**
- Commission bonus percentage
- Monthly bonus amount
- Exclusive features
- Recognition

### 5. Authentication & Authorization

**JWT-based Authentication:**

```
User Login
    ↓
Validate Credentials
    ↓
Generate JWT Token (24h expiry)
    ↓
Return Token to Client
    ↓
Client stores token (localStorage)
    ↓
Include in Authorization header
    ↓
Backend validates on each request
```

**Token Structure:**
```json
{
  "distributor_id": 123,
  "email": "user@example.com",
  "exp": 1234567890,
  "iat": 1234567890
}
```

## Database Schema

### Core Tables

1. **distributors**
   - Personal information
   - MLM structure (sponsor_id, tree_type, position, level)
   - Business metrics
   - Status and rank

2. **ranks**
   - Rank definitions
   - Requirements
   - Benefits

3. **packages**
   - Distributor packages/plans
   - Pricing
   - Commission rates
   - Features

4. **orders**
   - Order information
   - Distributor association
   - Payment details
   - Shipping information

5. **order_items**
   - Product details
   - Quantity and pricing
   - Commissionable value

6. **products**
   - Product catalog
   - Pricing
   - Stock management
   - Categories

7. **commissions**
   - Commission records
   - Type and level
   - Amount and percentage
   - Status tracking

8. **rank_achievements**
   - Rank progression history
   - Achievement dates
   - Notes

9. **payouts**
   - Payout records
   - Payment methods
   - Status tracking
   - Transaction IDs

### Relationships

```
distributors (1) ←→ (N) distributors (self-referencing sponsor)
distributors (1) ←→ (N) orders
distributors (1) ←→ (N) commissions
distributors (N) ←→ (1) ranks
distributors (N) ←→ (1) packages
orders (1) ←→ (N) order_items
products (N) ←→ (1) categories
```

## API Design

### RESTful Endpoints

**Authentication:**
- `POST /api/v1/distributors/register` - Register new distributor
- `POST /api/v1/distributors/login` - Login

**Distributor Management:**
- `GET /api/v1/distributors/profile` - Get current user profile
- `PUT /api/v1/distributors/profile` - Update profile
- `GET /api/v1/distributors/:id` - Get distributor by ID
- `GET /api/v1/distributors` - List distributors (paginated)
- `GET /api/v1/distributors/:id/downlines` - Get downlines
- `GET /api/v1/distributors/:id/tree` - Get tree structure
- `POST /api/v1/distributors/add-member` - Add member to tree

**Response Format:**

Success:
```json
{
  "data": {...},
  "message": "Success"
}
```

Error:
```json
{
  "error": "Error message"
}
```

Paginated:
```json
{
  "data": [...],
  "total": 100,
  "page": 1,
  "limit": 10
}
```

## Frontend Architecture

### State Management (Redux Toolkit)

**Store Structure:**
```
store/
├── store.ts          # Store configuration
├── hooks.ts          # Typed hooks
└── features/
    ├── auth/
    │   └── authSlice.ts
    ├── distributors/
    │   └── distributorSlice.ts
    └── dashboard/
        └── dashboardSlice.ts
```

**Slice Example:**
```typescript
const distributorSlice = createSlice({
  name: 'distributors',
  initialState,
  reducers: {
    setDistributors,
    setCurrentDistributor,
    updateDistributor,
  },
  extraReducers: (builder) => {
    // Async thunks
  },
});
```

### Component Structure

**Atomic Design Pattern:**
- **Atoms**: Basic UI elements (Button, Input, Label)
- **Molecules**: Simple component groups (FormField, Card)
- **Organisms**: Complex components (Header, TreeView, Dashboard)
- **Templates**: Page layouts
- **Pages**: Complete views

### Routing

**E-commerce (Next.js App Router):**
```
app/
├── page.tsx              # Home
├── products/
│   └── page.tsx          # Products list
├── login/
│   └── page.tsx          # Login
└── register/
    └── page.tsx          # Registration
```

**Admin Dashboard & Portal (React Router):**
```typescript
<Routes>
  <Route path="/" element={<Dashboard />} />
  <Route path="/distributors" element={<Distributors />} />
  <Route path="/tree" element={<TreeView />} />
  <Route path="/commissions" element={<Commissions />} />
  <Route path="/reports" element={<Reports />} />
</Routes>
```

## Security Considerations

### Backend Security

1. **Password Hashing**: bcrypt with salt
2. **JWT Tokens**: Secure, expiring tokens
3. **CORS**: Configured allowed origins
4. **Input Validation**: Request validation
5. **SQL Injection Prevention**: GORM parameterized queries
6. **Rate Limiting**: (To be implemented)

### Frontend Security

1. **Token Storage**: localStorage (consider httpOnly cookies for production)
2. **XSS Prevention**: React's built-in escaping
3. **CSRF Protection**: Token-based authentication
4. **HTTPS**: Required in production
5. **Environment Variables**: Sensitive data in .env files

## Performance Optimization

### Backend

1. **Database Indexing**: Key fields indexed
2. **Query Optimization**: Efficient GORM queries
3. **Caching**: (To be implemented - Redis)
4. **Connection Pooling**: GORM connection pool
5. **Pagination**: All list endpoints paginated

### Frontend

1. **Code Splitting**: Next.js automatic splitting
2. **Lazy Loading**: React.lazy for components
3. **Image Optimization**: Next.js Image component
4. **Bundle Optimization**: Tree shaking
5. **Memoization**: React.memo, useMemo, useCallback

## Testing Strategy

### Backend Testing

```go
// Unit tests
func TestDistributorService_Register(t *testing.T) {
    // Test implementation
}

// Integration tests
func TestDistributorAPI_Register(t *testing.T) {
    // Test implementation
}
```

### Frontend Testing

```typescript
// Component tests
describe('DistributorCard', () => {
  it('renders distributor information', () => {
    // Test implementation
  });
});

// Integration tests
describe('DistributorList', () => {
  it('fetches and displays distributors', () => {
    // Test implementation
  });
});
```

## Deployment

### Backend Deployment

1. Build binary: `go build -o mlm-api cmd/server/main.go`
2. Deploy to server (AWS, GCP, DigitalOcean)
3. Use process manager (systemd, PM2)
4. Configure reverse proxy (Nginx)
5. Enable HTTPS (Let's Encrypt)

### Frontend Deployment

1. Build: `npm run build`
2. Deploy to:
   - **Next.js**: Vercel, AWS Amplify
   - **React Apps**: Netlify, Vercel, AWS S3 + CloudFront
3. Configure environment variables
4. Enable CDN

## Monitoring & Logging

### Backend

- Structured logging (to be implemented)
- Error tracking (Sentry)
- Performance monitoring (New Relic, DataDog)
- Database monitoring

### Frontend

- Error boundary components
- Analytics (Google Analytics, Mixpanel)
- Performance monitoring (Web Vitals)
- User session recording (LogRocket, FullStory)

## Scalability

### Horizontal Scaling

- Stateless backend (JWT auth)
- Load balancer (Nginx, AWS ALB)
- Multiple backend instances
- Database read replicas

### Vertical Scaling

- Optimize queries
- Add caching layer
- Increase server resources
- Database optimization

## Future Enhancements

1. **Payment Integration**: Stripe, PayPal
2. **Email Notifications**: SendGrid, AWS SES
3. **SMS Notifications**: Twilio
4. **Mobile Apps**: React Native
5. **Real-time Updates**: WebSockets
6. **Advanced Analytics**: Custom reporting
7. **Multi-language Support**: i18n
8. **Multi-currency Support**: Currency conversion
9. **Document Management**: File uploads
10. **Training Portal**: Video courses

## Conclusion

This MLM application is built with:
- ✅ Best coding practices
- ✅ Clean architecture
- ✅ Modular design
- ✅ Scalability in mind
- ✅ Security considerations
- ✅ Modern tech stack
- ✅ Comprehensive documentation

The architecture supports all required MLM features including multiple tree types, commission calculations, rank systems, and complete distributor management.
