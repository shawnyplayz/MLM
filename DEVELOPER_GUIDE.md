# MLM Application - Developer Guide

## 🎯 For Developers

This guide helps developers understand, extend, and customize the MLM application.

---

## 📚 Table of Contents

1. [Getting Started](#getting-started)
2. [Project Structure](#project-structure)
3. [Backend Development](#backend-development)
4. [Frontend Development](#frontend-development)
5. [Adding New Features](#adding-new-features)
6. [Testing](#testing)
7. [Deployment](#deployment)
8. [Best Practices](#best-practices)

---

## 🚀 Getting Started

### Prerequisites Knowledge

**Backend:**
- Go programming language
- Gin web framework
- GORM ORM
- MySQL database
- RESTful API design
- JWT authentication

**Frontend:**
- JavaScript/TypeScript
- React 18
- Next.js 14
- Redux Toolkit
- TailwindCSS
- Axios

### Development Environment Setup

```bash
# 1. Clone/navigate to project
cd F:/Shawn/MLM

# 2. Install backend dependencies
cd backend
go mod download

# 3. Install frontend dependencies
cd ../ecommerce && npm install
cd ../admin-dashboard && npm install
cd ../distributor-portal && npm install

# 4. Set up environment files
# Copy .env.example to .env in each directory
# Update with your credentials
```

---

## 📁 Project Structure

### Backend Structure

```
backend/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go              # Configuration management
│   ├── domain/
│   │   └── models.go              # Business entities
│   ├── repository/
│   │   ├── distributor_repository.go
│   │   ├── order_repository.go
│   │   ├── commission_repository.go
│   │   └── rank_repository.go
│   ├── service/
│   │   ├── distributor_service.go
│   │   ├── tree_service.go
│   │   └── commission_service.go
│   ├── controller/
│   │   └── distributor_controller.go
│   └── middleware/
│       └── auth.go
├── pkg/
│   └── database/
│       └── database.go
├── go.mod
└── .env.example
```

### Frontend Structure

```
[frontend-app]/
├── src/
│   ├── components/          # Reusable UI components
│   ├── features/            # Redux slices
│   ├── pages/               # Route pages
│   ├── services/            # API services
│   ├── store/               # Redux store
│   ├── utils/               # Helper functions
│   └── types/               # TypeScript types
├── public/                  # Static assets
├── package.json
├── tsconfig.json
└── .env.example
```

---

## 🔧 Backend Development

### Adding a New Model

**Step 1: Define the model in `internal/domain/models.go`**

```go
type YourModel struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
    
    // Your fields
    Name      string         `gorm:"size:100;not null" json:"name"`
    Value     float64        `gorm:"type:decimal(15,2)" json:"value"`
    
    // Relationships
    UserID    uint           `gorm:"not null;index" json:"user_id"`
    User      *Distributor   `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
```

**Step 2: Add migration in `pkg/database/database.go`**

```go
func AutoMigrate(db *gorm.DB) error {
    err := db.AutoMigrate(
        &domain.Distributor{},
        &domain.YourModel{},  // Add your model
        // ... other models
    )
    return err
}
```

### Adding a New Repository

**Create `internal/repository/your_repository.go`**

```go
package repository

import (
    "github.com/mlm-app/backend/internal/domain"
    "gorm.io/gorm"
)

type YourRepository interface {
    Create(item *domain.YourModel) error
    FindByID(id uint) (*domain.YourModel, error)
    Update(item *domain.YourModel) error
    Delete(id uint) error
    List(offset, limit int) ([]domain.YourModel, int64, error)
}

type yourRepository struct {
    db *gorm.DB
}

func NewYourRepository(db *gorm.DB) YourRepository {
    return &yourRepository{db: db}
}

func (r *yourRepository) Create(item *domain.YourModel) error {
    return r.db.Create(item).Error
}

func (r *yourRepository) FindByID(id uint) (*domain.YourModel, error) {
    var item domain.YourModel
    err := r.db.First(&item, id).Error
    if err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *yourRepository) Update(item *domain.YourModel) error {
    return r.db.Save(item).Error
}

func (r *yourRepository) Delete(id uint) error {
    return r.db.Delete(&domain.YourModel{}, id).Error
}

func (r *yourRepository) List(offset, limit int) ([]domain.YourModel, int64, error) {
    var items []domain.YourModel
    var total int64
    
    err := r.db.Model(&domain.YourModel{}).Count(&total).Error
    if err != nil {
        return nil, 0, err
    }
    
    err = r.db.Offset(offset).Limit(limit).Find(&items).Error
    return items, total, err
}
```

### Adding a New Service

**Create `internal/service/your_service.go`**

```go
package service

import (
    "github.com/mlm-app/backend/internal/domain"
    "github.com/mlm-app/backend/internal/repository"
)

type YourService interface {
    ProcessItem(id uint) error
    GetItemWithDetails(id uint) (*domain.YourModel, error)
}

type yourService struct {
    yourRepo repository.YourRepository
    // Add other dependencies
}

func NewYourService(yourRepo repository.YourRepository) YourService {
    return &yourService{
        yourRepo: yourRepo,
    }
}

func (s *yourService) ProcessItem(id uint) error {
    item, err := s.yourRepo.FindByID(id)
    if err != nil {
        return err
    }
    
    // Business logic here
    
    return s.yourRepo.Update(item)
}

func (s *yourService) GetItemWithDetails(id uint) (*domain.YourModel, error) {
    return s.yourRepo.FindByID(id)
}
```

### Adding a New Controller

**Create `internal/controller/your_controller.go`**

```go
package controller

import (
    "net/http"
    "strconv"
    
    "github.com/gin-gonic/gin"
    "github.com/mlm-app/backend/internal/service"
)

type YourController struct {
    yourService service.YourService
}

func NewYourController(yourService service.YourService) *YourController {
    return &YourController{
        yourService: yourService,
    }
}

// GetItem godoc
// @Summary Get item by ID
// @Tags your-resource
// @Produce json
// @Security BearerAuth
// @Param id path int true "Item ID"
// @Success 200 {object} domain.YourModel
// @Router /api/v1/your-resource/{id} [get]
func (ctrl *YourController) GetItem(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    
    item, err := ctrl.yourService.GetItemWithDetails(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, item)
}

// CreateItem godoc
// @Summary Create new item
// @Tags your-resource
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param item body CreateRequest true "Item data"
// @Success 201 {object} domain.YourModel
// @Router /api/v1/your-resource [post]
func (ctrl *YourController) CreateItem(c *gin.Context) {
    var req CreateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // Process request
    
    c.JSON(http.StatusCreated, gin.H{"message": "Created successfully"})
}

type CreateRequest struct {
    Name  string  `json:"name" binding:"required"`
    Value float64 `json:"value" binding:"required"`
}
```

### Registering Routes

**Update `cmd/server/main.go`**

```go
// Initialize repository
yourRepo := repository.NewYourRepository(db)

// Initialize service
yourService := service.NewYourService(yourRepo)

// Initialize controller
yourController := controller.NewYourController(yourService)

// Register routes
protected := v1.Group("")
protected.Use(middleware.AuthMiddleware(cfg))
{
    protected.GET("/your-resource/:id", yourController.GetItem)
    protected.POST("/your-resource", yourController.CreateItem)
    protected.PUT("/your-resource/:id", yourController.UpdateItem)
    protected.DELETE("/your-resource/:id", yourController.DeleteItem)
}
```

---

## 🎨 Frontend Development

### Adding a Redux Slice

**Create `src/features/yourFeature/yourSlice.ts`**

```typescript
import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { yourService } from '../../services/yourService';

interface YourState {
  items: YourItem[];
  currentItem: YourItem | null;
  loading: boolean;
  error: string | null;
}

const initialState: YourState = {
  items: [],
  currentItem: null,
  loading: false,
  error: null,
};

// Async thunks
export const fetchItems = createAsyncThunk(
  'your/fetchItems',
  async () => {
    const response = await yourService.getAll();
    return response.data;
  }
);

export const fetchItemById = createAsyncThunk(
  'your/fetchItemById',
  async (id: number) => {
    const response = await yourService.getById(id);
    return response.data;
  }
);

export const createItem = createAsyncThunk(
  'your/createItem',
  async (data: CreateItemData) => {
    const response = await yourService.create(data);
    return response.data;
  }
);

// Slice
const yourSlice = createSlice({
  name: 'your',
  initialState,
  reducers: {
    clearError: (state) => {
      state.error = null;
    },
    setCurrentItem: (state, action: PayloadAction<YourItem>) => {
      state.currentItem = action.payload;
    },
  },
  extraReducers: (builder) => {
    builder
      // Fetch items
      .addCase(fetchItems.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchItems.fulfilled, (state, action) => {
        state.loading = false;
        state.items = action.payload;
      })
      .addCase(fetchItems.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error.message || 'Failed to fetch items';
      })
      // Fetch item by ID
      .addCase(fetchItemById.fulfilled, (state, action) => {
        state.currentItem = action.payload;
      })
      // Create item
      .addCase(createItem.fulfilled, (state, action) => {
        state.items.push(action.payload);
      });
  },
});

export const { clearError, setCurrentItem } = yourSlice.actions;
export default yourSlice.reducer;
```

### Adding an API Service

**Create `src/services/yourService.ts`**

```typescript
import api from '../lib/api';

export interface YourItem {
  id: number;
  name: string;
  value: number;
  created_at: string;
}

export interface CreateItemData {
  name: string;
  value: number;
}

export const yourService = {
  getAll: () => api.get<YourItem[]>('/your-resource'),
  
  getById: (id: number) => api.get<YourItem>(`/your-resource/${id}`),
  
  create: (data: CreateItemData) => api.post<YourItem>('/your-resource', data),
  
  update: (id: number, data: Partial<CreateItemData>) => 
    api.put<YourItem>(`/your-resource/${id}`, data),
  
  delete: (id: number) => api.delete(`/your-resource/${id}`),
};
```

### Creating a Component

**Create `src/components/YourComponent.tsx`**

```typescript
import React, { useEffect } from 'react';
import { useAppDispatch, useAppSelector } from '../store/hooks';
import { fetchItems } from '../features/yourFeature/yourSlice';

export const YourComponent: React.FC = () => {
  const dispatch = useAppDispatch();
  const { items, loading, error } = useAppSelector((state) => state.your);
  
  useEffect(() => {
    dispatch(fetchItems());
  }, [dispatch]);
  
  if (loading) {
    return <div>Loading...</div>;
  }
  
  if (error) {
    return <div className="text-red-500">Error: {error}</div>;
  }
  
  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">Your Items</h1>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        {items.map((item) => (
          <div key={item.id} className="border rounded-lg p-4 shadow">
            <h2 className="text-xl font-semibold">{item.name}</h2>
            <p className="text-gray-600">${item.value}</p>
          </div>
        ))}
      </div>
    </div>
  );
};
```

### Adding a Page (React Router)

**Update `src/App.tsx` or routing file**

```typescript
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { YourComponent } from './components/YourComponent';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Dashboard />} />
        <Route path="/your-page" element={<YourComponent />} />
        {/* Other routes */}
      </Routes>
    </BrowserRouter>
  );
}
```

---

## 🆕 Adding New Features

### Example: Adding Product Reviews

**1. Backend Model**

```go
type ProductReview struct {
    ID            uint      `gorm:"primaryKey" json:"id"`
    CreatedAt     time.Time `json:"created_at"`
    ProductID     uint      `gorm:"not null;index" json:"product_id"`
    DistributorID uint      `gorm:"not null;index" json:"distributor_id"`
    Rating        int       `gorm:"not null" json:"rating"`
    Comment       string    `gorm:"type:text" json:"comment"`
}
```

**2. Repository**

```go
type ProductReviewRepository interface {
    Create(review *domain.ProductReview) error
    FindByProductID(productID uint) ([]domain.ProductReview, error)
    FindByDistributorID(distributorID uint) ([]domain.ProductReview, error)
}
```

**3. Service**

```go
type ProductReviewService interface {
    CreateReview(review *domain.ProductReview) error
    GetProductReviews(productID uint) ([]domain.ProductReview, error)
    GetAverageRating(productID uint) (float64, error)
}
```

**4. Controller**

```go
func (ctrl *ProductReviewController) CreateReview(c *gin.Context) {
    // Implementation
}
```

**5. Frontend Redux Slice**

```typescript
export const createReview = createAsyncThunk(
  'reviews/create',
  async (data: CreateReviewData) => {
    const response = await reviewService.create(data);
    return response.data;
  }
);
```

---

## 🧪 Testing

### Backend Testing

**Unit Test Example**

```go
// internal/service/distributor_service_test.go
package service

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

type MockDistributorRepository struct {
    mock.Mock
}

func (m *MockDistributorRepository) FindByID(id uint) (*domain.Distributor, error) {
    args := m.Called(id)
    return args.Get(0).(*domain.Distributor), args.Error(1)
}

func TestDistributorService_GetByID(t *testing.T) {
    mockRepo := new(MockDistributorRepository)
    service := NewDistributorService(mockRepo, nil, nil)
    
    expectedDist := &domain.Distributor{ID: 1, FirstName: "John"}
    mockRepo.On("FindByID", uint(1)).Return(expectedDist, nil)
    
    result, err := service.GetByID(1)
    
    assert.NoError(t, err)
    assert.Equal(t, expectedDist, result)
    mockRepo.AssertExpectations(t)
}
```

### Frontend Testing

**Component Test Example**

```typescript
// src/components/__tests__/YourComponent.test.tsx
import { render, screen } from '@testing-library/react';
import { Provider } from 'react-redux';
import { store } from '../store/store';
import { YourComponent } from '../YourComponent';

describe('YourComponent', () => {
  it('renders items correctly', () => {
    render(
      <Provider store={store}>
        <YourComponent />
      </Provider>
    );
    
    expect(screen.getByText('Your Items')).toBeInTheDocument();
  });
});
```

---

## 🚀 Deployment

### Backend Deployment

**Build Binary**

```bash
cd backend
go build -o mlm-api cmd/server/main.go
```

**Docker Example**

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o mlm-api cmd/server/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/mlm-api .
COPY --from=builder /app/.env .
EXPOSE 8080
CMD ["./mlm-api"]
```

### Frontend Deployment

**Build for Production**

```bash
# Next.js
cd ecommerce
npm run build
npm start

# React (Vite)
cd admin-dashboard
npm run build
# Serve dist/ folder
```

---

## ✅ Best Practices

### Backend

1. **Always use interfaces** for repositories and services
2. **Handle errors properly** - don't ignore errors
3. **Use transactions** for multi-step database operations
4. **Validate input** at controller level
5. **Log important operations**
6. **Use constants** for magic strings/numbers
7. **Write tests** for business logic
8. **Document complex functions**

### Frontend

1. **Use TypeScript** for type safety
2. **Keep components small** and focused
3. **Use Redux for global state** only
4. **Memoize expensive computations**
5. **Handle loading and error states**
6. **Use environment variables** for config
7. **Follow naming conventions**
8. **Write reusable components**

### Database

1. **Always use migrations**
2. **Index foreign keys**
3. **Use appropriate data types**
4. **Normalize when appropriate**
5. **Backup regularly**

### Security

1. **Never commit .env files**
2. **Validate all user input**
3. **Use parameterized queries**
4. **Hash passwords properly**
5. **Use HTTPS in production**
6. **Implement rate limiting**
7. **Keep dependencies updated**

---

## 📖 Additional Resources

- **Go Documentation**: https://golang.org/doc/
- **Gin Framework**: https://gin-gonic.com/docs/
- **GORM**: https://gorm.io/docs/
- **React**: https://react.dev/
- **Next.js**: https://nextjs.org/docs
- **Redux Toolkit**: https://redux-toolkit.js.org/
- **TailwindCSS**: https://tailwindcss.com/docs

---

## 🤝 Contributing

When contributing to this project:

1. Follow the existing code structure
2. Write tests for new features
3. Update documentation
4. Follow coding standards
5. Create meaningful commit messages
6. Test thoroughly before submitting

---

**Happy Coding! 🚀**
