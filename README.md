# ShortURL Project (Backend)

ShortURL is a backend service for a URL shortening application, responsible for generating and managing compact links from long URLs. Built with Go and Fiber, this service follows a Clean Architecture approach for maintainability and scalability, and uses PostgreSQL as the database.

## Tech Stack
- **Backend**: Go (Golang)
- **Framework**: Fiber
- **Database**: PostgreSQL
- **Architecture**: Clean Architecture

## Setup / Running Locally
> **Note:** Instructions for setting up and running the backend locally will be added here.
1. Clone the repository 
    ```bash
    git clone https://github.com/zulfahmidev/TECHNICAL_TEST_SHORTURL.git
    ```
2. Navigate to the project directory
    ```bash
    cd TECHNICAL_TEST_SHORTURL
    ```
3. Install dependencies
    ```bash
    go mod tidy
    ```
4. Setup your .env file (copy .env.example and update values)
5. Install dependencies for migrations
- **Make** (if not installed)
  - macOS/Linux: biasanya sudah ada
  - Windows: install via Git Bash atau Chocolatey (`choco install make`)
- **golang-migrate** CLI
  - macOS: `brew install golang-migrate`
  - Linux/Windows: download binary dari https://github.com/golang-migrate/migrate/releases
5. Run database migrations
    ```bash
    make migrate-up
    ```
6. Start the server
    ```bash
    go run .
    ```

## Running Test
```bash
go test ./tests
```

## Live Deployment
The backend API is publicly available at: [https://api-shorturl.zulfahmidev.com](https://api-shorturl.zulfahmidev.com)  

## Design Decisions / Assumptions
* Chose Go + Fiber for backend due to performance, simplicity, and alignment with project requirements.
* Followed Clean Architecture principles to separate concerns, improve maintainability, and simplify future extensions.
* PostgreSQL was selected for reliable relational data storage.
* Assumes URLs are unique for shortening; idempotent creation can return an existing short URL if already present.
* Centralized error handling and logging for easier debugging and monitoring.

## Dependencies
* Go >= 1.24.*
* Fiber
* GORM (for database ORM)
* PostgreSQL
