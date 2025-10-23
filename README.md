# IoT Backend Service for ESP32 - Smart Agriculture

A Go-based backend service designed to receive encrypted sensor data from ESP32 devices and forward it to Blynk IoT platform for smart agriculture monitoring and visualization.

## 📋 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [API Documentation](#api-documentation)
- [Project Structure](#project-structure)
- [Usage](#usage)
- [Development](#development)
- [License](#license)

## 🌟 Overview

This backend service acts as a secure middleware between ESP32 IoT devices and the Blynk IoT platform. It receives encrypted sensor data from ESP32 devices, decrypts it, and forwards the agricultural sensor readings to Blynk for real-time monitoring and control.

### Supported Sensors
- **Soil Moisture** - Monitors soil moisture levels
- **Temperature** - Tracks ambient temperature
- **Humidity** - Measures air humidity
- **Light Intensity** - Monitors light levels
- **Air Pollution** - Measures air quality

## ✨ Features

- **🔐 Secure Data Transmission**: AES-CTR encryption for ESP32-to-server communication
- **📡 Blynk Integration**: Automatic data forwarding to Blynk IoT platform
- **🚀 RESTful API**: Clean and simple API endpoints
- **📊 Swagger Documentation**: Auto-generated API documentation
- **🔧 Environment-based Configuration**: Support for development and production environments
- **💾 Database Support**: SQL Server integration via GORM
- **☁️ Cloudinary Integration**: Ready for media file handling
- **🔄 Migration Support**: Database migration utilities

## 🏗️ Architecture

The project follows a clean architecture pattern with clear separation of concerns:

```
┌─────────────┐      ┌──────────────┐      ┌─────────────┐
│   ESP32     │─────▶│  Backend API │─────▶│    Blynk    │
│  (Sensors)  │      │  (Go Server) │      │  Platform   │
└─────────────┘      └──────────────┘      └─────────────┘
   Encrypted            Decrypt +              Display
     JSON              Process Data          Dashboard
```

**Layers:**
- **Controller Layer**: Handles HTTP requests and responses
- **Service Layer**: Business logic implementation
- **Repository Layer**: Data access and external API communication
- **Helper Layer**: Utility functions (encryption, JSON parsing, database)
- **Configuration Layer**: Environment and database setup

## 📦 Prerequisites

- **Go** 1.22.0 or higher
- **SQL Server** (for database operations)
- **Blynk Account** with API token
- **Git** for version control

## 🚀 Installation

1. **Clone the repository:**
```bash
git clone https://github.com/devinluize/IotBackendServiceESP32.git
cd IotBackendServiceESP32
```

2. **Install dependencies:**
```bash
go mod download
```

3. **Create environment configuration files:**

Create `.development/app.env` for development:
```env
# Server Configuration
SERVER_HOSTNAME=localhost
SERVER_PORT=3000

# Database Configuration
DB_DRIVER=sqlserver
DB_USER=your_db_user
DB_PASS=your_db_password
DB_NAME=your_db_name
DB_HOST=localhost
DB_PORT=1433

# Blynk Configuration
BLYNK_API_TOKEN=your_blynk_token
BLYNK_API_URL=https://blynk.cloud/external/api/
DATA_STREAM_SOIL_MOISTURE=V0
DATA_STREAM_TEMPERATURE=V1
DATA_STREAM_LIGHT_INTENSITY=V2
DATA_STREAM_POLLUTION_LEVEL=V3
DATA_STREAM_HUMIDITY=V4

# Other Configuration
CLIENT_ORIGIN=http://localhost:3000
GENERAL_API=/api
```

Create `.production/app.env` for production with similar structure.

## ⚙️ Configuration

### Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `SERVER_HOSTNAME` | Server hostname | `localhost` |
| `SERVER_PORT` | Server port | `3000` |
| `DB_DRIVER` | Database driver | `sqlserver` |
| `DB_USER` | Database username | `sa` |
| `DB_PASS` | Database password | `your_password` |
| `DB_NAME` | Database name | `iot_agriculture` |
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `1433` |
| `BLYNK_API_TOKEN` | Blynk authentication token | `your_blynk_token` |
| `BLYNK_API_URL` | Blynk API endpoint | `https://blynk.cloud/external/api/` |
| `DATA_STREAM_SOIL_MOISTURE` | Blynk virtual pin for soil moisture | `V0` |
| `DATA_STREAM_TEMPERATURE` | Blynk virtual pin for temperature | `V1` |
| `DATA_STREAM_LIGHT_INTENSITY` | Blynk virtual pin for light | `V2` |
| `DATA_STREAM_POLLUTION_LEVEL` | Blynk virtual pin for air quality | `V3` |
| `DATA_STREAM_HUMIDITY` | Blynk virtual pin for humidity | `V4` |

### Encryption Keys

The AES-CTR encryption uses the following default keys (located in `api/helper/encrypt/EncryptHelper.go`):
- **Key**: `ThisIsAESkey1234` (16 bytes)
- **IV**: `ESP32InitVector1` (16 bytes)

⚠️ **Important**: Change these keys in production for security!

## 📖 API Documentation

### Base URL
```
http://localhost:3000
```

### Endpoints

#### Send Sensor Data
```http
POST /api/blynk/
Content-Type: application/json

{
  "blynk_esp_32_request": "base64_encrypted_data"
}
```

**Encrypted Payload Structure** (before encryption):
```json
{
  "soil_moisture": 45.5,
  "light_intensity": 850.2,
  "temperature": 28.3,
  "air_pollution": 150,
  "humidity": 65.8
}
```

**Response (Success):**
```json
{
  "status": "success",
  "message": "success to send data to blynk",
  "data": "success"
}
```

**Response (Error):**
```json
{
  "message": "Failed to decrypt data",
  "error": "error details",
  "success": false
}
```

### Swagger Documentation

Access interactive API documentation at:
```
http://localhost:3000/swagger/index.html
```

## 📂 Project Structure

```
IotBackendServiceESP32/
├── main.go                 # Application entry point
├── mainEncrypt.go         # Encryption utilities
├── go.mod                 # Go module dependencies
├── README.md              # This file
├── request.http           # HTTP request examples
├── api/
│   ├── config/           # Configuration and database connection
│   │   ├── Configuration.go
│   │   └── Connection.go
│   ├── controller/       # HTTP request handlers
│   │   └── blynkController/
│   │       └── BlynkController.go
│   ├── helper/          # Utility functions
│   │   ├── general.go
│   │   ├── json.go
│   │   ├── database/
│   │   │   └── databaseHelper.go
│   │   └── encrypt/
│   │       └── EncryptHelper.go
│   ├── middleware/      # HTTP middlewares
│   │   ├── Cors.go
│   │   └── middleware.go
│   ├── payloads/       # Request/Response structures
│   │   ├── blynk/
│   │   │   └── BlynkPayloads.go
│   │   └── responses/
│   │       ├── error.go
│   │       ├── PaginationPayloads.go
│   │       └── standarResponses.go
│   ├── repositories/   # Data access layer
│   │   └── blynk/
│   │       ├── blynkRepository.go
│   │       └── blynk-repository-impl/
│   │           └── BlynkRepositoryImpl.go
│   ├── route/         # Route definitions
│   │   ├── RegisterRouter.go
│   │   └── Route.go
│   ├── service/       # Business logic layer
│   │   └── blynk/
│   │       ├── BlynkService.go
│   │       └── BlynkServiceImpl/
│   │           └── BlynkServiceImpl.go
│   └── test/         # Test files
│       └── bookmark_test.go
├── docs/             # Swagger documentation
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
└── migrate/          # Database migrations
    └── migrate.go
```

## 🔧 Usage

### Running the Application

**Development Mode:**
```bash
go run main.go
```

**Production Mode:**
```bash
go run main.go prod
```

**Run Migrations:**
```bash
go run main.go migrate
```

### Testing with curl

```bash
# Test endpoint (replace with actual encrypted data)
curl -X POST http://localhost:3000/api/blynk/ \
  -H "Content-Type: application/json" \
  -d '{
    "blynk_esp_32_request": "your_encrypted_base64_data"
  }'
```

### ESP32 Integration Example

```cpp
// Arduino/ESP32 code snippet
#include <HTTPClient.h>
#include <ArduinoJson.h>

void sendDataToServer() {
  HTTPClient http;
  
  // Create JSON payload
  StaticJsonDocument<200> doc;
  doc["soil_moisture"] = soilMoisture;
  doc["light_intensity"] = lightIntensity;
  doc["temperature"] = temperature;
  doc["air_pollution"] = airPollution;
  doc["humidity"] = humidity;
  
  String jsonString;
  serializeJson(doc, jsonString);
  
  // Encrypt the JSON string (implement AES-CTR encryption)
  String encrypted = encryptAESCTR(jsonString);
  
  // Create request payload
  StaticJsonDocument<300> requestDoc;
  requestDoc["blynk_esp_32_request"] = encrypted;
  
  String requestString;
  serializeJson(requestDoc, requestString);
  
  // Send to server
  http.begin("http://your-server:3000/api/blynk/");
  http.addHeader("Content-Type", "application/json");
  int httpCode = http.POST(requestString);
  
  http.end();
}
```

## 👨‍💻 Development

### Generate Swagger Documentation

After modifying API annotations:
```bash
swag init
```

### Adding New Endpoints

1. Create controller in `api/controller/`
2. Create service in `api/service/`
3. Create repository in `api/repositories/`
4. Add route in `api/route/`
5. Update swagger annotations

### Running Tests

```bash
go test ./...
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📧 Contact

**Developer**: Devin Luize  
**Email**: devin@gmail.com  
**Project**: IoT Blynk Smart Agriculture Backend Service

## 📄 License

This project is part of a thesis work. Please contact the author for licensing information.

## 🙏 Acknowledgments

- **Blynk** - IoT platform integration
- **Chi Router** - Fast HTTP router
- **GORM** - ORM library for Go
- **Swagger** - API documentation