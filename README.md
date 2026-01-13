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

---

## API Documentation

### Base URL
```
http://localhost:8080/api/v1
```

### Authentication
Pennie uses Firebase for authentication.

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

## Built with:
- OpenAI GPT-4 for intelligent document analysis
- Go community for excellent libraries
- Redis Labs for caching and rate limiting patterns
- My wife, for veterinary expertise.

---