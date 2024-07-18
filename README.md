# URL Shortener

URL shortener service written in Go.

## Getting Started

### Prerequisites

- **Go 1.16 or higher**
- **Redis installed and running**.
  You can installed it from their [website](https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/) or running it with Docker with the following command:
  ```bash
  docker run -d --name redis-stack-server -p 6379:6379 redis/redis-stack-server:latest
  ```

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/url-shortener.git
   cd url-shortener
   ```
2. Set up a .env file with the following environment variables in the root of the project:
   ```bash
   HOST=http://localhost:8080/
   REDIS_ADDR=localhost:6379 #Adjust the port if necessary
   REDIS_DB=0
   ```
3. Run the application:
   ```bash
   go run src/cmd/main.go
   ```

## API Endpoints

- **POST /createLink**: Create a short URL.
  - Request Body: `{"url": "http://example.com"}`
  - Response: `{"message": "short url created successfully!", "url": "http://localhost:8080/short123"}`
- **GET /:link**: Redirect to the original URL.
  - Response: Redirects to the original URL.
