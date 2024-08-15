
# Task Management RESTful API

Welcome to the **Task Management RESTful API** built with the Gin framework. This API provides features for user authentication, user management, and task management, following the principles of Clean Architecture.

## 📁 Folder Structure

The project is organized into the following directory structure:

## 🚀 Features

- **Authentication with JWT**: Secure authentication using JSON Web Tokens (JWT).
- **User Management**: Manage user sign-up, login, and user role promotion.
- **Task Management**: Create, retrieve, update, and delete tasks.

## 🏗️ Clean Architecture

This project follows the principles of Clean Architecture to ensure that the code is modular, maintainable, and scalable. The architecture is divided into layers with clearly defined responsibilities:

- **Delivery**: Handles the HTTP requests and responses, including routing and controllers.
- **Domain**: Contains the core business logic and domain models.
- **Infrastructure**: Manages the interaction with external services and systems.
- **Repositories**: Manages data access and persistence.
- **UseCases**: Implements the application's use cases and business rules.

```
bootstrap/
├── app.go
├── database.go
├── env.go

Delivery/
├── controllers/
│   ├── login.controller.go
│   ├── promote_user.controller.go
│   ├── sign_up.controller.go
│   └── task.controller.go
├── main.go
├── routers/
│   ├── login.router.go
│   ├── promote.route.go
│   ├── router.go
│   ├── sign_up.router.go
│   └── task.router.go
└── tmp/
    ├── build-errors.log
    └── main
doc/
└── api_documentation.md

Domain/
├── auth.middleware.go
├── error_response.go
├── jwt_custome.go
├── jwt.service.go
├── login.go
├── password.service.go
├── promote.go
├── sign_up.go
├── task.go
├── user.go
└── validate.go

Infrastructure/
├── auth.middleware.go
├── jwt.service.go
└── password.service.go
mocks/
├── mock_auth_middler.go
├── mock_jwt_service.go
├── mock_login_usecase.go
├── mock_password_service.go
├── mock_promote_usecase.go
├── mock_signup_repository.go
├── mock_task.go
└── mock_user_repository.go

Repositories/
├── task_repository.go
└── user_repository.go

tests/
├── constants/
│   └── user.go
├── controllers/
│   ├── login_controller_test.go
│   ├── promote_controller_test.go
│   ├── signup_controller_test.go
│   └── task_controller_test.go
├── Repositories/
│   ├── task_repository_test.go
│   └── user_repository_test.go
└── UseCases/
    ├── login_usecase_test.go
    ├── promote_user_usecase_test.go
    ├── sign_up_usecase_test.go
    └── task_usecase_test.go

UseCases/
├── login.usecase.go
├── promote.usecase.go
├── sign_up.usecase.go
└── task.usecase.go

go.mod
go.sum
makefile
README.md
```

## 🏁 Getting Started

To get started with the project:

1. **Clone the Repository**:
   ```sh
   git clone https://github.com/solo21-12/task_management_API_gin.git
   cd task_management_API_gin
   ```

2. **Build the Project**:
   ```sh
   make build
   ```

3. **Run the Application**:
   ```sh
   make run
   ```

4. **Run Tests**:
   ```sh
   make test
   ```

5. **Run Tests with Coverage**:
   ```sh
   make test-coverage
   ```

## 🧪 Testing

Testing is crucial for maintaining the quality of the application. To run tests, use the following commands:

- **Run Tests**:
  ```sh
  make test
  ```

- **Run Tests with Coverage**:
  ```sh
  make test-coverage
  ```

## 📜 Documentation

For detailed API documentation, please refer to `doc/api_documentation.md`.

For more information, please visit the [GitHub repository](https://github.com/solo21-12/task_management_API_gin).

---
