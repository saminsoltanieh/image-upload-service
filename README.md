# Image Upload Service

## Project Overview
This project is a simple backend service written in **Go** that allows users to upload, store, and retrieve image files through HTTP endpoints.

The main goals of this project were to practice backend development concepts such as:
- Building an HTTP server in Go
- Handling multipart file uploads
- Designing REST-style endpoints
- Structuring a Go project using `cmd/` and `internal/`

---

## Technical Stack
- **Language:** Go  
- **Protocol:** HTTP  
- **Storage:** Local filesystem  
- **Project Structure:** Modular (`cmd` + `internal`)

---

## Project Structure
```
image-upload-service/
├── cmd/server          # Application entry point
├── internal/upload     # Upload & storage logic
├── uploads/            # Stored uploaded images
├── go.mod              # Go modules
└── README.md
```

---

## API Endpoints
- `POST /upload` — Upload an image (multipart/form-data)
- `GET /images/{filename}` — Retrieve an uploaded image

> Note: Endpoint paths may vary slightly depending on your router implementation.

---

## How to Run
```bash
git clone https://github.com/saminsoltanieh/image-upload-service
cd image-upload-service
go mod download
go run ./cmd/server
```

---

## Design Decisions
- **Local storage** was selected for simplicity and fast prototyping.
- A **modular structure** separates the app entry point (`cmd/`) from business logic (`internal/`).
- No database is used in this version to keep the service lightweight and focused on core upload functionality.

---

## Future Improvements
- Add authentication/authorization
- Add file validation (type/size limits)
- Store metadata in a database (e.g., PostgreSQL)
- Support cloud storage (e.g., S3)
- Add automated tests and CI

---

## Learning Outcomes
Through this project, I gained hands-on experience with:
- Implementing REST-like APIs in Go
- Handling file uploads and filesystem storage
- Organizing code for maintainability in Go projects
- Managing dependencies using Go modules
