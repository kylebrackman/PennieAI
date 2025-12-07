# PennieAI
## Backend

**AI-Powered Veterinary Document Analysis System**

PennieAI is a production-focused veterinary assistant that streamlines clinical workflows by intelligently analyzing and extracting structured data from unstructured veterinary documents. Using OpenAI's advanced language models, PennieAI can process emails, patient records, lab reports, and clinical notes to automatically extract patient information and segment individual documents from large concatenated files.

---

## Table of Contents

- [Overview](#overview)
- [Current Features](#current-features)
- [Architecture](#architecture)
- [Sliding Window Approach](#sliding-window-approach)
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

## Sliding Window Approach

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

## Getting Started

### Running locally in docker

#### Start docker in Development mode with:

`docker-compose -f docker-compose.dev.yml up`

Add the -d flag to run in detached mode if you would like to keep your terminal free.

#### Ensure that it is running with:

`docker-compose -f docker-compose.dev.yml ps`

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

## API Documentation

### Base URL
```
http://localhost:8080/api/v1
```

### Authentication
Currently no authentication is required. API authentication (JWT) is planned for Phase 7 (Production Readiness).

---


**Rate Limiting:**
- 100 requests per hour per IP address
- Headers returned:
  - `X-RateLimit-Limit`: 100
  - `X-RateLimit-Remaining`: 99
  - `X-RateLimit-Reset`: Unix timestamp

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


---

## Roadmap

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

### Development Principles

- **Learn by Building**: This project prioritizes learning production patterns
- **Incremental Improvement**: Each phase adds new capabilities
- **Best Practices**: Follow Go idioms and industry standards

---

## Built with:
- OpenAI GPT-4 for intelligent document analysis
- Go community for excellent libraries
- Redis Labs for caching and rate limiting patterns
- My wife, for veterinary expertise.

---