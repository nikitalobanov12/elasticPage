# ElasticPage

A microservices-based platform that transforms static PDF textbooks into interactive learning experiences using AI. Built with Go backend services deployed as AWS Lambdas and a Next.js frontend.

## 🏗️ Architecture

- **Frontend**: Next.js 15+ with TypeScript, TanStack Query v5, Tailwind CSS, T3 Stack
- **Backend**: Go microservices with Gin framework, GORM, AWS Lambda deployment
- **Database**: PostgreSQL (AWS RDS free tier)
- **Storage**: AWS S3 for PDF files
- **AI**: HuggingFace Transformers API (planned)
- **Infrastructure**: Docker Compose for local development, AWS for production

## 📁 Project Structure

```
elasticPage/
├── backend/
│   ├── lambdas/
│   │   └── content-service/     # Textbook management service
│   ├── shared/                  # Shared Go packages
│   │   ├── models/             # Database models
│   │   ├── utils/              # Utility functions
│   │   └── middleware/         # HTTP middleware
│   └── infrastructure/         # Docker Compose & deployment configs
├── frontend/                   # Next.js T3 Stack application
├── docs/                      # Documentation
└── .github/workflows/         # CI/CD workflows
```

## 🚀 Getting Started

### Prerequisites

- Go 1.24+
- Node.js 18+
- Docker & Docker Compose
- AWS CLI (for production deployment)

### Local Development Setup

1. **Clone and navigate to the project**:
   ```bash
   git clone <repository-url>
   cd elasticPage
   ```

2. **Start the database and services**:
   ```bash
   cd backend/infrastructure
   docker-compose up -d
   ```

3. **Set up the backend**:
   ```bash
   cd ../lambdas/content-service
   cp .env.example .env
   # Edit .env with your configuration
   go mod tidy
   go run cmd/api/main.go
   ```

4. **Set up the frontend**:
   ```bash
   cd ../../../frontend
   npm install
   cp .env.example .env
   # Edit .env with your configuration
   npm run db:push
   npm run dev
   ```

5. **Access the application**:
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080
   - Database: localhost:5432

### Environment Variables

#### Backend (.env)
```bash
# Server Configuration
PORT=8080

# Database Configuration
BLUEPRINT_DB_HOST=localhost
BLUEPRINT_DB_PORT=5432
BLUEPRINT_DB_DATABASE=elasticpage
BLUEPRINT_DB_USERNAME=postgres
BLUEPRINT_DB_PASSWORD=postgres
BLUEPRINT_DB_SCHEMA=public

# AWS Configuration
AWS_REGION=us-east-1
S3_BUCKET_NAME=elasticpage-files

# For local development with LocalStack
AWS_ENDPOINT_URL=http://localhost:4566
AWS_ACCESS_KEY_ID=test
AWS_SECRET_ACCESS_KEY=test
```

#### Frontend (.env)
```bash
# Database
DATABASE_URL="postgresql://postgres:postgres@localhost:5432/elasticpage"

# NextAuth.js
NEXTAUTH_SECRET="your-secret-key"
NEXTAUTH_URL="http://localhost:3000"

# API Configuration
NEXT_PUBLIC_API_URL="http://localhost:8080"
```

## 🛠️ Development

### Backend Development

The backend uses a microservices architecture with the following services:

- **content-service**: Handles textbook CRUD operations and file uploads
- **ai-service**: (Planned) AI processing for summaries and quizzes
- **quiz-service**: (Planned) Interactive quiz generation and management
- **user-service**: (Planned) User management and authentication

#### Key Features Implemented:
- ✅ GORM-based database models (User, Textbook)
- ✅ RESTful API endpoints for textbook management
- ✅ File upload handling with S3 integration
- ✅ AWS Lambda deployment configuration
- ✅ Docker Compose for local development

#### API Endpoints:
```
GET    /api/v1/textbooks          # List textbooks
POST   /api/v1/textbooks          # Create textbook
GET    /api/v1/textbooks/:id      # Get textbook by ID
PUT    /api/v1/textbooks/:id      # Update textbook
DELETE /api/v1/textbooks/:id      # Delete textbook
POST   /api/v1/upload             # Upload textbook PDF
```

### Frontend Development

The frontend is built with the T3 Stack and includes:

- ✅ Next.js 15 with App Router
- ✅ TypeScript for type safety
- ✅ TanStack Query v5 for server state management
- ✅ NextAuth.js for authentication
- ✅ Tailwind CSS for styling
- ✅ Custom API client for Go backend integration

#### Key Pages:
- **Home** (`/`): Landing page with platform overview
- **Dashboard** (`/dashboard`): User's textbook library
- **Upload** (`/upload`): PDF upload form
- **Textbook View** (`/textbooks/[id]`): Individual textbook interface (planned)

## 🧪 Testing

### Backend Testing
```bash
cd backend/lambdas/content-service
go test ./...
```

### Frontend Testing
```bash
cd frontend
npm run typecheck
```

## 🚀 Deployment

### AWS Lambda Deployment

The backend services are designed to run as AWS Lambda functions:

1. **Build for Lambda**:
   ```bash
   cd backend/lambdas/content-service
   GOOS=linux GOARCH=amd64 go build -o bootstrap cmd/lambda/main.go
   zip lambda-deployment.zip bootstrap
   ```

2. **Deploy with AWS CLI**:
   ```bash
   aws lambda create-function \
     --function-name content-service \
     --runtime provided.al2 \
     --role arn:aws:iam::ACCOUNT:role/lambda-execution-role \
     --handler bootstrap \
     --zip-file fileb://lambda-deployment.zip
   ```

### Frontend Deployment

Deploy to Vercel or any Next.js-compatible platform:

```bash
cd frontend
npm run build
```

## 🔮 Planned Features

### Phase 2: AI Integration
- [ ] HuggingFace API integration for text processing
- [ ] Automatic chapter/section detection
- [ ] AI-generated summaries and key points
- [ ] Intelligent quiz generation

### Phase 3: Interactive Features
- [ ] Real-time collaborative annotations
- [ ] Progress tracking and analytics
- [ ] WebSocket support for live features
- [ ] Mobile-responsive reading interface

### Phase 4: Advanced Features
- [ ] Multi-language support
- [ ] Advanced search and indexing
- [ ] Integration with learning management systems
- [ ] Offline reading capabilities



**Note**: This is a learning project designed to demonstrate competence with modern web development tools and cloud architecture patterns. The platform is designed to stay within AWS free tier limits for cost-effective learning and demonstration.