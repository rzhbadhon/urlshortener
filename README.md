# 🔗 Go URL Shortener

<p align="center">
  <img src="logo.png" alt="Go URL Shortener Logo" width="120"/>
</p>

<p align="center">
  <b>A lightweight URL shortener service built with Go, SQLx, and PostgreSQL.</b>
</p>

---

## 📖 Overview

This project is a simple yet powerful **URL Shortener** written in Go.  
It allows users to shorten long URLs and automatically sets an expiration time (default: 48 hours).  

The service demonstrates:

- 🌀 Go’s HTTP server capabilities  
- 🗄️ Database integration with **SQLx**  
- ⏳ Expiration handling with `time`  
- 🛠️ Clean handler structure for scalability  

---

## ⚙️ Features

- Shorten any valid URL via a **POST request**  
- Auto‑generate a unique short code using `utils.Shortner`  
- Store original URL, short code, and expiration in the database  
- Return JSON response with short URL and expiration timestamp  
- Built with clean, modular Go code  

---

## 🚀 Getting Started

### Prerequisites
- Go 1.18+  
- PostgreSQL  
- `sqlx` library  

### Installation
Clone the repository:

```bash
git clone https://github.com/rzhbadhon/urlshortener.git
cd urlshortener

urlshortener/
├── handlers/         # HTTP handlers
├── models/           # Request/response models
├── rest/handlers/    # URL Handler
├── rest/middlewares/ # CORS Handler
├── docs/             # HTML Frontend
├── utils/            # Utility function(Shortner)
├── main.go           # Entry point
