# PennieAI

**AI-Powered Veterinary Document Analysis System**

PennieAI is a production-focused veterinary assistant that streamlines clinical workflows by intelligently analyzing and extracting structured data from unstructured veterinary documents. Using OpenAI's advanced language models, PennieAI can process emails, patient records, lab reports, and clinical notes to automatically extract patient information and segment individual documents from large concatenated files.

---

## Table of Contents

- [Overview](#overview)
- [Current Features](#current-features)
- [Architecture](#architecture)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [API Documentation](#api-documentation)
- [Development](#development)
- [Roadmap](#roadmap)
- [Contributing](#contributing)

---

## Overview

### The Problem

Veterinarians often receive large files containing multiple documents - patient histories, lab results, referral emails, and clinical notes all concatenated together. Manually parsing these files is time-consuming and error-prone, taking valuable time away from patient care.

### The Solution

PennieAI automates this process by:

1. **Intelligent Document Segmentation**: Uses a sliding window algorithm to identify document boundaries in large files
2. **Patient Information Extraction**: Automatically extracts and structures patient data (name, species, breed, sex, weight, etc.)
3. **Incremental Processing**: Builds on previously extracted data to avoid duplicates and ensure completeness
4. **Structured Storage**: Stores all extracted information in PostgreSQL for easy retrieval and analysis
5. **Rate-Limited AI Integration**: Implements Redis-backed rate limiting to control OpenAI API costs

---

## Current Features

### Core Functionality

- **Document Analysis**: Upload veterinary documents and receive structured patient data and segmented documents
- **Sliding Window Processing**: Intelligently processes large files in chunks to handle documents of any size
- **Patient Entity Recognition**: Extracts patient demographics, physical characteristics, and medical metadata
- **Document Boundary Detection**: Identifies where individual documents begin and end within concatenated files
- **Duplicate Prevention**: Avoids re-extracting information already found in previous windows

### Infrastructure

- **RESTful API**: Clean, versioned API (v1) with standard HTTP methods
- **PostgreSQL Database**: Robust relational storage for patients, documents, and inference logs
- **Redis Integration**:
  - Rate limiting for OpenAI API calls (100 requests/hour by default)
  - Ready for future caching and background job queues
- **CORS Support**: Configured for cross-origin requests
- **Health Checks**: `/health` endpoint for monitoring
- **Environment-Based Configuration**: Uses `.env` files for secrets and configuration

### API Capabilities

- **CRUD Operations**: Full create, read, update, delete for analyzed documents
- **Document Upload**: Multipart form-data support for document uploads
- **AI Testing Endpoints**: Test AI service connectivity and model versions
- **Rate Limit Headers**: Standard `X-RateLimit-*` headers on all responses

---

## Architecture

### High-Level Flow

```
User Upload (PDF/TXT)
    ↓
[Gin HTTP Handler]
    ↓
[Rate Limit Middleware] (Redis)
    ↓
[Document Segmenter] (Sliding Window)
    ↓
[OpenAI API] (GPT-4 Analysis)
    ↓
[Data Extraction & Deduplication]
    ↓
[PostgreSQL Storage]
    ↓
JSON Response (Patient + Documents)
```

### Sliding Window Algorithm

PennieAI uses a sophisticated sliding window approach to process large documents:

1. **Window Creation**: Document is split into overlapping windows of text
2. **AI Analysis**: Each window is analyzed by OpenAI to identify:
   - Patient information within the window
   - Complete documents (with start/end line numbers)
3. **Incremental Building**: Each subsequent window receives:
   - Previously extracted patient data
   - List of already-identified documents
4. **Deduplication**: New findings are merged with existing data, avoiding duplicates by comparing line numbers

This allows PennieAI to handle documents of unlimited length while maintaining context and avoiding redundant processing.

---

## Tech Stack

### Backend
- **Go 1.24**: High-performance, compiled language
- **Gin Web Framework**: Fast HTTP routing and middleware
- **sqlx**: Enhanced SQL operations with struct mapping
- **go-redis**: Redis client for caching and rate limiting

### Data Storage
- **PostgreSQL**: Primary relational database
  - Patients table
  - Analyzed documents table
  - Unprocessed documents table
  - Inferences table (AI request/response logging)
- **Redis**: In-memory data structure store
  - Rate limiting counters
  - Future: Caching, background job queues

### AI Integration
- **OpenAI GPT-4**: Document analysis and entity extraction
- **openai-go**: Official OpenAI Go client

### Infrastructure (Current)
- **godotenv**: Environment variable management
- **CORS Middleware**: Cross-origin request handling

---

## Project Structure

```
PennieAI/
├── main.go                          # Application entry point
├── config/
│   ├── database.go                  # PostgreSQL connection & pool
│   └── redis.go                     # Redis connection & pool
├── models/
│   ├── patient.go                   # Patient entity
│   ├── analyzed_document.go         # Analyzed document entity
│   ├── unprocessed_document.go      # Unprocessed document entity
│   └── inference.go                 # AI inference logging
├── handlers/
│   ├── analyze_document.go          # Document upload & analysis
│   ├── analyzed_documents.go        # CRUD for analyzed docs
│   ├── unprocessed_documents.go     # Unprocessed doc handlers
│   └── ai_tool.go                   # AI service test endpoints
├── services/
│   ├── ai_analyze.go                # Core AI analysis logic
│   ├── ai_tool.go                   # OpenAI service wrapper
│   └── document_segmenter.go        # Document chunking (future)
├── middleware/
│   ├── cors.go                      # CORS configuration
│   └── rate_limit.go                # Redis-backed rate limiting
├── routes/
│   └── routes.go                    # API route definitions
├── prompts/
│   └── document_analysis.go         # OpenAI prompt templates
├── utils/
│   ├── window_builder.go            # Sliding window implementation
│   └── get_file_lines.go            # File parsing utilities
├── mock_data/                       # Sample veterinary documents
├── learning_resources/              # Documentation & learning notes
└── learning-roadmap-checklist.md    # Development roadmap
```

---

## Getting Started

### Prerequisites

- **Go 1.24+** installed
- **PostgreSQL** running (local or remote)
- **Redis** running (local or Docker)
- **OpenAI API Key** with GPT-4 access

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/PennieAI.git
   cd PennieAI
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**

   Create a `.env` file in the project root:
   ```env
   # Database
   DATABASE_URL=postgres://username:password@localhost:5432/pennie_development?sslmode=disable

   # Redis
   REDIS_URL=localhost:6379
   REDIS_PASSWORD=

   # OpenAI
   OPENAI_API_KEY=sk-...

   # Server
   PORT=8080
   GIN_MODE=debug  # or "release" for production
   ```

4. **Set up the database**

   Create the PostgreSQL database and tables:
   ```sql
   CREATE DATABASE pennie_development;

   -- Connect to the database and create tables
   -- (You'll need to create tables based on the models in models/)
   ```

5. **Start Redis**

   Using Docker:
   ```bash
   docker run -d -p 6379:6379 redis:alpine
   ```

   Or install locally:
   ```bash
   # macOS
   brew install redis
   redis-server

   # Linux
   sudo apt-get install redis-server
   redis-server
   ```

6. **Run the application**
   ```bash
   go run main.go
   ```

   You should see:
   ```
   ✅ Database connected successfully
   ✅ Redis connected successfully
   🐾 PennieAI API starting on port 8080...
   📍 Health check: http://localhost:8080/health
   📍 API docs: http://localhost:8080/api/v1
   ```

---

## API Documentation

### Base URL
```
http://localhost:8080/api/v1
```

### Authentication
Currently no authentication is required. API authentication (JWT) is planned for Phase 7 (Production Readiness).

### Endpoints

#### Health Check
```http
GET /health
```

**Response:**
```json
{
  "status": "healthy",
  "service": "PennieAI",
  "version": "1.0.0",
  "message": "Veterinary AI assistant is running! 🐾"
}
```

---

#### Analyze Unprocessed Document
Upload a document for AI analysis.

```http
POST /api/v1/unprocessed/analyze
Content-Type: multipart/form-data
```

**Request:**
- `document` (file): The document to analyze (TXT, PDF, etc.)

**Response:**
```json
{
  "message": "Document analyzed successfully",
  "count": 3,
  "patient": {
    "id": 0,
    "name": "Bella",
    "possibleSpecies": ["Dog"],
    "possibleBreed": ["Golden Retriever"],
    "sex": "Female",
    "date_of_birth": "2019-05-15",
    "weight": 28.5,
    "height": null,
    "color": "Golden",
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-01-15T10:30:00Z"
  },
  "documents": [
    {
      "id": 0,
      "title": "Initial Consultation Notes",
      "content": "",
      "numberOfLines": 45,
      "patientId": 0,
      "startLine": 1,
      "endLine": 45,
      "unprocessedDocumentId": 0,
      "createdAt": "2025-01-15T10:30:00Z",
      "updatedAt": "2025-01-15T10:30:00Z",
      "windowLines": ["Line 1...", "Line 2..."]
    }
  ]
}
```

**Rate Limiting:**
- 100 requests per hour per IP address
- Headers returned:
  - `X-RateLimit-Limit`: 100
  - `X-RateLimit-Remaining`: 99
  - `X-RateLimit-Reset`: Unix timestamp

---

#### Get All Analyzed Documents
```http
GET /api/v1/documents
```

**Response:**
```json
{
  "documents": [
    {
      "id": 1,
      "title": "Lab Report - Blood Work",
      "content": "...",
      "numberOfLines": 87,
      "patientId": 1,
      "startLine": 1,
      "endLine": 87,
      "unprocessedDocumentId": 1,
      "createdAt": "2025-01-15T10:30:00Z",
      "updatedAt": "2025-01-15T10:30:00Z"
    }
  ]
}
```

---

#### Get Document by ID
```http
GET /api/v1/documents/:id
```

**Response:**
```json
{
  "document": {
    "id": 1,
    "title": "Lab Report - Blood Work",
    "content": "...",
    "numberOfLines": 87,
    "patientId": 1,
    "startLine": 1,
    "endLine": 87,
    "unprocessedDocumentId": 1,
    "createdAt": "2025-01-15T10:30:00Z",
    "updatedAt": "2025-01-15T10:30:00Z"
  }
}
```

---

#### Create Document
```http
POST /api/v1/documents
Content-Type: application/json
```

**Request:**
```json
{
  "title": "Physical Exam Notes",
  "content": "Patient appears healthy...",
  "patientId": 1
}
```

---

#### Delete Document
```http
DELETE /api/v1/documents/:id
```

**Response:**
```json
{
  "message": "Document deleted successfully"
}
```

---

#### Test AI Service
```http
GET /api/v1/ai_tool/test
```

Tests connectivity to OpenAI API.

---

#### Get AI Model Version
```http
GET /api/v1/ai_tool/model_version
```

Returns the current OpenAI model being used.

---

## Development

### Running in Development Mode

```bash
# Install Air for hot reload (optional but recommended)
go install github.com/cosmtrek/air@latest

# Run with hot reload
air

# Or run normally
go run main.go
```

### Testing with Redis

```bash
# Check if Redis is running
redis-cli ping
# Should return: PONG

# Monitor Redis commands in real-time
redis-cli monitor

# Check rate limit keys
redis-cli KEYS "ratelimit:*"
```

### Database Migrations

Currently, database schema is managed manually. Database migrations (using tools like `goose` or `migrate`) are planned for Phase 7.

### Code Style

- Follow standard Go formatting: `go fmt ./...`
- Run linters: `go vet ./...`
- Use descriptive variable names
- Comment complex logic (especially in `ai_analyze.go`)

---

## Roadmap

PennieAI is on a journey from a CRUD application to a production-ready, observable, and scalable system. See `learning-roadmap-checklist.md` for the complete roadmap.

### Phase 1: Redis - Caching & Rate Limiting ✅ (In Progress)

**Completed:**
- ✅ Redis installation and connection
- ✅ Redis client integration (`go-redis`)
- ✅ Rate limiting middleware implementation
- ✅ Rate limit HTTP headers
- ✅ 429 status code handling

**Next Steps:**
- [ ] Implement caching for expensive OpenAI responses
- [ ] Cache patient lookups
- [ ] Add cache invalidation logic
- [ ] Monitor cache hit/miss rates
- [ ] Implement distributed locks for duplicate prevention

---

### Phase 2: Background Jobs & Async Processing 🔜

**Goals:**
- Move OpenAI processing to background jobs (prevents HTTP timeouts)
- Implement job queue with `asynq` (Redis-backed)
- Add job status tracking and progress updates
- Implement retry logic with exponential backoff
- Create dead letter queue for failed jobs
- Build real-time progress updates (polling or WebSockets)

**Benefits:**
- Faster API responses (return immediately, process in background)
- Better error handling and retry capabilities
- Scalable worker pools
- No more timeout issues on large documents

---

### Phase 3: Docker & Containerization 🔜

**Goals:**
- Create multi-stage Dockerfile for Go app
- Set up `docker-compose.yml` for full stack:
  - PennieAI API
  - PostgreSQL
  - Redis
  - (Future: Prometheus, Grafana)
- Implement hot reload in Docker for development
- Document Docker setup and commands

**Benefits:**
- Consistent development environments
- Easy onboarding for new developers
- Production-ready containers
- Simple deployment to any Docker-compatible host

---

### Phase 4: API Documentation & Design 🔜

**Goals:**
- Implement OpenAPI/Swagger documentation
- Add Swagger UI endpoint (`/docs`)
- Document all request/response schemas
- Include examples for every endpoint
- Version API properly (v1, v2 strategy)
- Generate API clients from spec

**Tools:**
- `swaggo/swag` for Go annotations
- Swagger UI for interactive docs

---

### Phase 5: Testing Infrastructure 🔜

**Goals:**
- Unit tests for business logic (70%+ coverage)
- Integration tests with `testcontainers-go`
- Mock OpenAI responses for tests
- CI/CD pipeline (GitHub Actions)
- Pre-commit hooks for running tests

**Focus Areas:**
- Test patient extraction logic
- Test document segmentation algorithm
- Test rate limiting behavior
- Test error handling (API failures, timeouts)

---

### Phase 6: Observability & Monitoring 🔜

**Goals:**
- Replace `log` with structured logging (`zap` or `zerolog`)
- Add Prometheus metrics:
  - Request duration histograms
  - Request counts by status code
  - OpenAI API call costs
  - Document processing times
  - Background job metrics
- Set up Grafana dashboards
- Implement distributed tracing (OpenTelemetry)
- Create cost tracking dashboard for OpenAI usage

**Benefits:**
- Understand system performance in production
- Track API costs in real-time
- Debug issues faster with traces
- Set up alerts for errors and budget overruns

---

### Phase 7: Production Readiness 🔜

**Security:**
- [ ] JWT authentication
- [ ] HTTPS/TLS support
- [ ] Input validation and sanitization
- [ ] SQL injection prevention (already using parameterized queries)
- [ ] Rate limiting (already implemented)
- [ ] Security headers
- [ ] Secrets management (HashiCorp Vault or similar)

**Performance:**
- [ ] Database indexing for common queries
- [ ] Connection pooling optimization
- [ ] Profile with pprof
- [ ] Response compression (gzip)
- [ ] Circuit breakers for OpenAI API

**Deployment:**
- [ ] CI/CD pipeline
- [ ] Blue-green deployments
- [ ] Health checks and readiness probes
- [ ] Graceful shutdown
- [ ] Automated database backups
- [ ] Error tracking (Sentry or Rollbar)

---

### Bonus: Advanced Patterns (Future)

- **Event-Driven Architecture**: Event sourcing, Kafka/NATS
- **CQRS**: Separate read/write models
- **GraphQL**: Alternative to REST
- **gRPC**: For inter-service communication
- **Microservices**: Split into specialized services as needed

---

## Current Progress

### What Works Today

1. **Document Upload & Analysis**: Upload a veterinary document and receive structured patient data and segmented documents
2. **Intelligent Extraction**: AI extracts patient demographics and identifies document boundaries
3. **Rate Limited API**: 100 requests/hour to control OpenAI costs
4. **CRUD Operations**: Full management of analyzed documents
5. **Scalable Architecture**: PostgreSQL + Redis ready for production workloads

### Known Limitations

1. **Synchronous Processing**: Large documents can cause HTTP timeouts (will be fixed in Phase 2 with background jobs)
2. **No Caching**: Every analysis hits OpenAI API (expensive)
3. **No Authentication**: API is open to anyone (security risk)
4. **Manual Deployment**: No CI/CD pipeline yet
5. **Limited Observability**: Basic logging, no metrics or tracing

---

## Contributing

This is a learning project focused on building production-ready systems. Contributions are welcome!

### How to Contribute

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests (when available)
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

### Development Principles

- **Learn by Building**: This project prioritizes learning production patterns
- **Incremental Improvement**: Each phase adds new capabilities
- **Documentation**: Document what you learn for future reference
- **Best Practices**: Follow Go idioms and industry standards

---

## Learning Resources

This project is built following industry best practices. Check `learning_resources/` for:
- Q&A documents explaining complex concepts
- Notes on design decisions
- Links to helpful articles and tutorials

**Recommended Reading:**
- "Designing Data-Intensive Applications" by Martin Kleppmann
- "The Go Programming Language" by Donovan & Kernighan
- "Site Reliability Engineering" by Google

**Online Resources:**
- [Go by Example](https://gobyexample.com)
- [Redis University](https://redis.io/university)
- [OpenAI API Documentation](https://platform.openai.com/docs)

---

## License

[Add your license here]

---

## Acknowledgments

Built with:
- OpenAI GPT-4 for intelligent document analysis
- Go community for excellent libraries
- Redis Labs for caching and rate limiting patterns
- The veterinary community for domain expertise

---

**PennieAI** - Transforming veterinary document analysis, one AI call at a time. 🐾