# Todo-List Golang Backend

This is a simple and secure Todo-List backend API built with Golang, Gin, Gorm, and JWT-Go.

## Table of Contents
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [Contributing](#contributing)
- [License](#license)

## Features
- RESTful API endpoints for managing todo task:
- Create, read, update, and delete todo task
- User authentication and authorization using JWT-Go:
- Secure login and access control for protected endpoints
- Database persistence with PostgreSQL via GORM:
- Persistent storage of todo items
- Clear project structure and documentation:
- Easy to understand and maintain

## Requirements
- [Golang](https://go.dev/) version 1.18 or later
- [PostgreSQL](https://www.postgresql.org/) database 
- [Git](https://www.git-scm.com/) (optional for cloning the repository)

## Installation
#### 1. Clone the repository:
```sh
git clone https://github.com/TQP0403/todo-list.git
```

#### 2. Install dependencies:
```sh
cd todo-list
go mod download
```

#### 3. Create a .env file in the root directory by copy example.env:
```sh
cp -R example.env .env
```

#### 4. Run go Application:
```sh
go run main.go
```

## Contributing
Contributions are welcome! Please follow the standard fork-and-pull request workflow for contributing to this project.

## License
This project is licensed under the MIT License.
