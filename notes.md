# Go Notes

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