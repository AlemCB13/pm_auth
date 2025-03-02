# Authentication Microservice (pm_auth)

This microservice handles user authentication, including registration, login, and logout. It uses Redis as the database to store user credentials and JWT (JSON Web Tokens) for session management.

## Table of Contents
- [Requirements](#requirements)
- [Getting Started](#getting-started)
  - [Running with Docker](#running-with-docker)
  - [Running Locally](#running-locally)
- [API Endpoints](#api-endpoints)
  - [Register a User](#register-a-user)
  - [Login](#login)
  - [Logout](#logout)
  - [Two-Factor Authentication (2FA) (Optional)](#two-factor-authentication-2fa-optional)
- [Testing](#testing)
  - [Using Postman](#using-postman)
  - [Test Cases](#test-cases)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Requirements
- **Docker**: Ensure Docker and Docker Compose are installed on your machine.
- **Postman (optional)**: For testing the API endpoints.

## Getting Started

### Running with Docker

#### Clone the Repository:
```bash
git clone https://github.com/your-username/pm_auth.git
cd pm_auth
```

#### Build and Run the Containers:
```bash
docker-compose up --build
```

#### Access the Microservice:
The microservice will be available at [http://localhost:8080](http://localhost:8080).

### Running Locally

#### Install Go:
Download and install Go from the [official website](https://go.dev/dl/).

#### Install Dependencies:
```bash
go mod download
```

#### Run the Microservice:
```bash
go run src/main.go
```

#### Access the Microservice:
The microservice will be available at [http://localhost:8080](http://localhost:8080).

## API Endpoints

### Register a User
**Method:** `POST`

**URL:** `/register`

**Body (JSON):**
```json
{
  "username": "testuser",
  "password": "testpass",
  "email": "test@example.com"
}
```

### Login
**Method:** `POST`

**URL:** `/login`

**Body (JSON):**
```json
{
  "username": "testuser",
  "password": "testpass"
}
```

### Logout
**Method:** `POST`

**URL:** `/logout`

**Headers:**
```
Authorization: Bearer <token>
```

### Two-Factor Authentication (2FA) (Optional)
**Method:** `POST`

**URL:** `/2fa`

**Headers:**
```
Authorization: Bearer <token>
```

**Body (JSON):**
```json
{
  "token": "<token>",
  "code": "123456"
}
```

## Testing

### Using Postman
1. **Import the Collection:** Import the provided Postman collection to test the endpoints.
2. **Set Environment Variables:** Create an environment in Postman and set the token variable after logging in.

### Test Cases
- **Register a User:** Send a `POST` request to `/register` with valid user data.
- **Login:** Send a `POST` request to `/login` with valid credentials.
- **Logout:** Send a `POST` request to `/logout` with a valid token.
- **Register with Invalid Data:** Send a `POST` request to `/register` with incomplete or invalid data.
- **Login with Invalid Credentials:** Send a `POST` request to `/login` with incorrect credentials.
- **Access Protected Endpoint Without Token:** Try to access `/logout` without sending a token.
- **Access Protected Endpoint with Invalid Token:** Try to access `/logout` with an invalid or expired token.

## Contributing
Contributions are welcome! Please follow these steps:
1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature/YourFeature
   ```
3. Commit your changes:
   ```bash
   git commit -m 'Add some feature'
   ```
4. Push to the branch:
   ```bash
   git push origin feature/YourFeature
   ```
5. Open a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Contact
For any questions or issues, please open an issue on GitHub or contact the maintainer.
