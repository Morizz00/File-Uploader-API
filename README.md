# 📁 File Uploader API

A lightweight file upload API built with Golang using the Gin framework.

## 🚀 Features

- Upload files to the server via `POST /upload`
- Automatically creates an upload directory (`/uploads`)
- Minimal setup using Gin for routing

---

## 🔧 Installation

```bash
git clone https://github.com/Morizz00/file-uploader-api.git
cd file-uploader-api
go mod tidy

curl -X POST -F "file=@yourfile.txt" http://localhost:8080/upload
Built With
Go and Gin

▶️ Usage
bash
Copy
Edit
go run main.go
📫 Endpoint
http
Copy
Edit
POST http://localhost:8080/upload
🗂 Directory Structure
bash
Copy
Edit
.
├── main.go          # Starts the Gin server and routes
├── handler/
│   └── upload.go    # Handles file uploads
└── uploads/         # Directory where uploaded files are saved
🧪 Example Request (cURL)
bash
Copy
Edit
curl -X POST -F "file=@yourfile.txt" http://localhost:8080/upload
🛠 Built With
Go & Gin
