# Go Notes

## Open Questions & Notes to Review

- Review connection pools in the database.go file...
### Syntax Notes
Order matters - match the function's return signature
:= creates new variables,
= assigns to existing variables,
err reuse is very common in Go - it's idiomatic
## Naming Conventions

### Package Names: lowercase, single word
- package handlers  ✅
- package userauth  ✅
- package database  ✅

- package userAuth  ❌
- package user_auth ❌
- package UserAuth  ❌

### File Names: snake_case

- user_handler.go     ✅
- database_config.go  ✅
- auth_middleware.go  ✅

- userHandler.go      ❌
- UserHandler.go      

### Variables & Functions: camelCase

- var userName string        ✅
- var userCount int         ✅
- func getUserByID() {}     ✅

- var user_name string      ❌
- var UserName string       ❌ (unless exported)

### Exported (Public) Names: PascalCase
These can be used by other packages
- type User struct {}       ✅
- func GetUser() {}         ✅
- var DefaultConfig = {}    ✅

These are private to the package
- type user struct {}       ✅
- func getUser() {}         ✅
- var defaultConfig = {}    ✅


### Constants: Depends on visibility
Exported constants
- const MaxRetries = 3      ✅
- const DefaultTimeout = 30 ✅

Unexported constants
- const maxRetries = 3      ✅
- const defaultTimeout = 30 ✅

Not Go style
- const MAX_RETRIES = 3     ❌



### Proposed Structure

GearMateGo/                    # Your root directory
├── go.mod                     # Module definition (already created)
├── go.sum                     # Dependency checksums (created by go get)
├── main.go                    # Application entry point
├── .env                       # Environment variables
├── .gitignore                 # Git ignore rules
├── README.md                  # Project documentation
├── config/
│   ├── database.go           # Database configuration
│   └── config.go             # Application configuration
├── models/
│   ├── item.go               # Item model (you have this)
│   ├── user.go               # User model
│   └── database.go           # Database connection/migration helpers
├── handlers/
│   ├── items.go              # Item HTTP handlers
│   ├── users.go              # User HTTP handlers
│   ├── auth.go               # Authentication handlers
│   └── health.go             # Health check handler
├── middleware/
│   ├── auth.go               # Authentication middleware
│   ├── cors.go               # CORS middleware
│   ├── logging.go            # Request logging
│   └── validation.go         # Input validation middleware
├── routes/
│   └── routes.go             # Route definitions
├── services/                  # Business logic (optional, for complex apps)
│   ├── item_service.go
│   └── user_service.go
├── utils/
│   ├── response.go           # Standard API response helpers
│   ├── validation.go         # Validation helpers
│   └── helpers.go            # General utility functions
├── migrations/                # Database migrations (if using)
│   ├── 001_create_users.sql
│   ├── 002_create_items.sql
│   └── 003_add_indexes.sql
├── tests/                     # Test files
│   ├── handlers_test.go
│   ├── models_test.go
│   └── integration_test.go
└── docker/                    # Docker configuration (optional)
├── Dockerfile
└── docker-compose.yml