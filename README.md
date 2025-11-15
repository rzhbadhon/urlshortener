# 🔗 URL Shortener

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge)

A lightweight, fast, and efficient URL shortening service built with **Go** and **PostgreSQL**.

[Features](#-features) • [Installation](#-installation) • [Usage](#-usage) • [API Documentation](#-api-documentation) • [Contributing](#-contributing)

</div>

---

## 📋 Table of Contents

- [About](#-about)
- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Installation](#-installation)
- [Configuration](#-configuration)
- [Usage](#-usage)
- [API Documentation](#-api-documentation)
- [Project Structure](#-project-structure)
- [Contributing](#-contributing)
- [License](#-license)

---

## 🚀 About

**URL Shortener** is a modern, high-performance URL shortening service that transforms long URLs into compact, shareable links. Built with Go's blazing-fast performance and PostgreSQL's robust data management, this service provides a reliable solution for link management with automatic expiration handling.

---

## ✨ Features

- 🔗 **Instant URL Shortening** - Transform long URLs into short, memorable links
- ⏰ **Auto-Expiration** - Links automatically expire after 48 hours
- 🔄 **Seamless Redirects** - Fast redirection to original URLs
- 🗄️ **PostgreSQL Backend** - Reliable data persistence
- 🌐 **CORS Support** - Cross-origin resource sharing enabled
- 🎯 **RESTful API** - Clean and intuitive API design
- ⚡ **High Performance** - Built with Go for maximum efficiency
- 🛡️ **Error Handling** - Comprehensive error management

---

## 🛠️ Tech Stack

- **Backend Framework:** Go (Golang)
- **Database:** PostgreSQL
- **HTTP Router:** `net/http` (Standard Library)
- **Database Driver:** `sqlx`, `pq`
- **CORS Middleware:** Custom implementation

---

## 📦 Installation

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher
- Git

### Clone the Repository

```bash
git clone https://github.com/rzhbadhon/urlshortener.git
cd urlshortener
```

### Install Dependencies

```bash
go mod download
```

### Database Setup

1. Create a PostgreSQL database:

```sql
CREATE DATABASE urlshortner;
```

2. Create the required table:

```sql
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    short_code VARCHAR(10) UNIQUE,
    expire_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## ⚙️ Configuration

Update the database connection string in `cmd/serve.go`:

```go
connStr := "user=YOUR_USER password=YOUR_PASSWORD dbname=urlshortner sslmode=disable"
```

**Environment Variables** (recommended):

```bash
export DB_USER="postgres"
export DB_PASSWORD="your_password"
export DB_NAME="urlshortner"
export DB_HOST="localhost"
export DB_PORT="5432"
```

---

## 🎯 Usage

### Start the Server

```bash
go run main.go
```

The server will start on `http://localhost:5000`

### Create a Short URL

```bash
curl -X POST http://localhost:5000/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://www.example.com/very/long/url/that/needs/shortening"}'
```

**Response:**

```json
{
  "short_url": "http://localhost:5000/aB3xY",
  "expire_at": "2025-11-17 15:30:45"
}
```

### Access the Short URL

Simply visit or redirect to:

```
http://localhost:5000/aB3xY
```

---

## 📖 API Documentation

### 1. Create Short URL

**Endpoint:** `POST /shorten`

**Request Body:**

```json
{
  "url": "https://example.com/long-url"
}
```

**Response:**

```json
{
  "short_url": "http://localhost:5000/aB3xY",
  "expire_at": "2025-11-17 15:30:45"
}
```

**Status Codes:**
- `200 OK` - URL shortened successfully
- `400 Bad Request` - Invalid request body
- `405 Method Not Allowed` - Only POST method allowed
- `500 Internal Server Error` - Database error

### 2. Redirect to Original URL

**Endpoint:** `GET /{shortCode}`

**Response:** HTTP 302 redirect to original URL

**Status Codes:**
- `302 Found` - Redirect successful
- `404 Not Found` - Short URL not found
- `410 Gone` - Short URL has expired

---

## 📁 Project Structure

```
urlshortener/
├── cmd/
│   └── serve.go           # Server initialization and routing
├── docs/
│   └── index.html         # Documentation (if any)
├── models/
│   └── url.go             # Data models
├── rest/
│   ├── handlers/
│   │   ├── redirect.go    # Redirect handler
│   │   └── short_url.go   # URL shortening handler
│   └── middlewares/
│       └── cors.go        # CORS middleware
├── utils/
│   └── shortner.go        # URL shortening logic
├── go.mod                 # Go module definition
├── go.sum                 # Dependency checksums
├── main.go               # Application entry point
└── README.md             # Project documentation
```

---

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

1. **Fork** the repository
2. **Create** a feature branch (`git checkout -b feature/AmazingFeature`)
3. **Commit** your changes (`git commit -m 'Add some AmazingFeature'`)
4. **Push** to the branch (`git push origin feature/AmazingFeature`)
5. **Open** a Pull Request

### Development Guidelines

- Follow Go best practices and conventions
- Write clear, descriptive commit messages
- Add tests for new features
- Update documentation as needed

---

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 👨‍💻 Author

**Razibul Hasan Badhon**

- GitHub: [@rzhbadhon](https://github.com/rzhbadhon)
- Repository: [urlshortener](https://github.com/rzhbadhon/urlshortener)

---

## 🙏 Acknowledgments

- Built with ❤️ using Go
- Inspired by popular URL shortening services
- Thanks to the open-source community

---

<div align="center">

**If you find this project useful, please consider giving it a ⭐️**

Made with Go 🚀

</div>