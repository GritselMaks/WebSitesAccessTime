# WebSitesAccessTime
## API Documentation
## Overview
For faster running, you can use "make up" and "make down".

## Endpoints

### `/min`
- Description: Get the website with minimum access time.
- Method: `GET`
- Response:
  - Status: 200 OK
  - Body: JSON object {
    "access_time": int64,
    "url": "string"
}.

### `/max`
- Description: Get the website with maximumaximum access time.
- Method: `GET`
- Response:
  - Status: 200 OK
  - Body: JSON object {
    "access_time": int64,
    "url": "string"
}.

### `/url`
- Description: Get the access time for website.
- Method: `GET`
- Query Parameters:
  - `url` (required)
- Response:
  - Status: 200 OK
  - Body: JSON object {
    "access_time": int64,
    "url": "string"
}.

### `/metrics`
- Description: Get the metrics data for the API.
- Method: `GET`
- Authentication: Basic Authentication required.
- Response:
  - Status: 200 OK
  - Body: JSON object [
    {
        "counter": int64,
        "handler": string
    }    ].

