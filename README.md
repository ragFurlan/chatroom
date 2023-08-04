# go-clean-architecture-template

### FOLDER STRUCTURE

- cmd/
  - main.go                 // Application entry point
- internal/
  - adapters/               // Adapters to external frameworks
    - http/
      - handler/             // HTTP request handlers
  - app/
    - usecases/              // Application use cases
  - entities/               // Business entities
  - gateways/               // Interfaces for external services
  - repositories/           // Database access
- web/                      // Static files for the web interface
- migrations/               // Database migration scripts
- config/                   // Configuration files
- scripts/                  // Utility scripts
- tests/                    // Unit tests
- Dockerfile                // Docker configuration
- docker-compose.yml        // Docker Compose configuration
- README.md                 // Project documentation