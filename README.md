# URL Shortener 

## Features

1. **Shorten URL**: Accepts a URL as input and returns a shortened URL.
2. **Redirection API**: Redirects to the original URL when the shortened URL is visited.
3. **Unique Shortening**: If the same URL is requested again, the same shortened URL will be returned.
4. **Metrics API**: Returns the top 3 most shortened domains.

## Installation

### Prerequisites:
- Golang 1.23 or later (I have used 1.23)
- Docker (for running in a container)

### Steps:

1. Clone the repository:
    ```bash
    git clone https://github.com/your-username/url_shortener.git
    cd url_shortener
    ```

2. Build the Go application:
    ```bash
    go build -o url_proj
    ```

3. Run the application locally:
    ```bash
    ./url_proj
    ```

    The application will start on `http://localhost:8080`.

4. Using Docker, you can build the Docker image and run the container:
    ```bash
    docker build -t url_shortener .
    docker run -p 8080:8080 url_shortener
    ```

    This will start the service on `http://localhost:8080`.

## API Endpoints

### 1. Shorten URL

**POST** `/shorten`

- **Request body**:
    ```json
    {
      "url": "https://example.com"
    }
    ```

- **Response**:
    ```json
    {
      "short_url": "abc123"
    }
    ```

### 2. Redirect to Original URL

**GET** `/:shortURL`

- **Request**: `GET /abc123`
- **Response**: Redirects to the original URL `https://example.com`.

### 3. Metrics

**GET** `/metrics`

- **Response**:
    ```json
    {
      "Udemy": 6,
      "YouTube": 4,
      "Wikipedia": 2
    }
    ```

## Tests

Unit tests are implemented to verify the functionality of the URL shortening, redirection, and metrics APIs. You can run the tests using the following command:

```bash
go test -v ./tests
