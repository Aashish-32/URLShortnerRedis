# URL Shortener

A simple URL shortening service implemented in Go using the Fiber web framework and Redis for data storage.

## Features

- Shorten long URLs into a custom short URL.
- Set an expiry time for short URLs.
- Rate limiting based on IP address.
- Redirect to original URL from the short URL.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Aashish-32/URL-Shortener.git 

2. Build the Docker images and start the application:

     ```bash
   docker-compose up -d

## Usage

- To shorten a URL, send a POST request to /api/v1 with a JSON payload containing the URL:

      ```bash
   curl -X POST http://localhost:3000/api/v1 -H "Content-Type: application/json" -d '{"url":"https://example.com"}'

- To access a shortened URL, simply enter it in your browser:
    ```bash
   [text](http://localhost:3000/{short_id})

 

