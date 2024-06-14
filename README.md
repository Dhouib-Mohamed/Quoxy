# Quoxy: API Authenticator Proxy

Quoxy is a project that implements an API authenticator proxy using Go and SQLite. It features a token subscription system managed via Cron jobs for secure token management and storage, along with a reverse proxy component to redirect authenticated requests to the API. The name "Quoxy" is a merge between the words "quote" and "proxy," reflecting its role in securely managing and directing API requests.
## Project Structure

```plaintext
.
├── api
│   └── handler
│       ├── health.go
│       ├── index.go
│       ├── not_found.go
│       ├── subscription.go
│       ├── token.go
│       └── version.go
├── cmd
│   └── token-proxy
├── config.example.yaml
├── deployment
│   └── database
├── internal
│   ├── database
│   ├── frequency_cron
│   ├── models
│   ├── proxy
│   ├── tests
│   └── token_handler
├── scripts
│   ├── debug.sh
│   ├── sql
│   ├── start.sh
│   └── test.sh
├── todo.md
├── util
│   ├── config
│   ├── env
│   ├── error_handler
│   ├── log
│   └── network
└── version.txt
```

Certainly! Here's a brief overview of the content in each folder of your project:

1. **api**: Contains handlers for different API endpoints.
    - **handler**: Contains individual handler files for different endpoints like health checks, subscriptions, tokens, etc.

2. **cmd**: Contains the main application entry point.
    - **token-proxy**: Contains the main Go file for starting the API authenticator proxy.

3. **config.example.yaml**: An example configuration file. Users can copy this file to `config.yaml` and modify it according to their needs.

4. **deployment**: Contains deployment-related files.
    - **database**: Contains a Docker Compose file for deploying the SQLite database in a container.

5. **internal**: Contains internal application logic.
    - **database**: Contains database-related logic and utilities.
    - **frequency_cron**: Contains logic for managing cron jobs.
    - **models**: Contains data models for subscriptions and tokens.
    - **proxy**: Contains logic for the reverse proxy.
    - **tests**: Contains test files and utilities for testing the application.
    - **token_handler**: Contains logic for handling tokens.

6. **scripts**: Contains shell scripts for debugging, starting, and testing the application.
   - **debug.sh**: Shell script for debugging purposes.
   - **start.sh**: Shell script for starting the application.
   - **test.sh**: Shell script for running tests.

7. **util**: Contains utility packages for the application.
   - **config**: Contains utilities for managing configuration settings.
   - **env**: Contains utilities for managing environment settings.
   - **error_handler**: Contains utilities for handling errors.
   - **log**: Contains utilities for logging.
   - **network**: Contains utilities for managing network settings.

8. **version.txt**: A file containing the version information for the project.

This structure organizes the project into logical components, making it easier to maintain and extend in the future.
## Features

- **Token Subscription System:** Utilizes Cron jobs for secure token management and storage.
- **Reverse Proxy:** Redirects authenticated requests to the API.

## Getting Started

### Prerequisites

- Go (1.22)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/api-authenticator-proxy.git
   ```
2. Navigate to the project directory:
   ```sh
   cd api-authenticator-proxy
   ```
3. Install dependencies:
   ```sh
   go mod download
   ```
4. Set up the configuration variables:
   ```sh
   cp config.example.yaml config.yaml
   ```
5. Set up the environment variables:
   ```sh
   cp .env.example .env
   ```

### Running the Application

1. Start the application:
   ```sh
   bash scripts/debug.sh
   ```
2. The application will run on the port specified in the `config.yaml` file.

### Running Tests

1. Run the test suite:
   ```sh
   bash scripts/test.sh
   ```

## Deployment

A `docker-compose.yaml` file is included for deploying the supported databases for testing reasons.

1. Navigate to the deployment directory:
   ```sh
   cd deployment/database
   ```
2. Start the database container:
   ```sh
   docker-compose up -d
   ```

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for any enhancements or bugs.

## License

This project is licensed under the MIT License.

## Contact

For any questions or inquiries, please contact [mdhouib195@gmail.com].
