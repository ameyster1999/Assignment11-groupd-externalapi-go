

```markdown
# Weather API Server

## Overview
This project implements an Weather API server using the Go programming language to provide current weather information for a specified city. It integrates with an external weather API to fetch real-time weather data and returns the results in JSON format. The application is containerized using Docker for easy deployment.

## API Specification

### GET /city
Accepts a city name as a query parameter and returns the current weather in JSON format.

Example request:
```
GET /city?name=Toronto
```

Example response:
```json
{
  "city": "Toronto",
  "temperature": "5°C",
  "weather": "Cloudy"
}
```

### POST /city
Accepts a JSON body with a city name and returns the current weather in JSON format.

Example request:
```
POST /city
Body: {"name": "Toronto"}
```

Example response:
```json
{
  "city": "Toronto",
  "temperature": "5°C",
  "weather": "Cloudy"
}
```

## Docker
The application is containerized using Docker for easy deployment across different environments.

### Dockerfile
The Dockerfile defines the container environment and instructions to build the application image.

### Building and Running
To build the Docker image locally:
```
docker build -t weather-api-server .
```

To run the Docker container:
```
docker run -p 8080:8080 weather-api-server
```

## External Weather API
For this project, we have integrated with [Weatherstack](https://api.openweathermap.org/) API to fetch current weather data.

## Repository Structure
- `main.go`: Contains the main code for the API server.
- `Dockerfile`: Defines the Docker container environment.
- `README.md`: Provides an overview of the project, setup instructions, and other relevant information.

## Setup and Running
1. Clone the repository:
   ```
   git clone https://github.com/ameyster1999/Assignment11-groupd-externalapi-go.git
   ```

2. Build the Docker image:
   ```
   docker build -t ameyster786/weatherapp .
   ```

3. Run the Docker container:
   ```
   docker run -p 8010:8010  ameyster786/weatherapp:latest  
   ```
   

4. Access the API:
    - GET request: Open a web browser or use a tool like Postman and visit `http://localhost:8080/city?name=Toronto`.
    - POST request: Send a POST request to `http://localhost:8080/city` with a JSON body like `{"name": "Toronto"}`.

## Submission
- GitHub Repository: [https://github.com/ameyster1999/Assignment11-groupd-externalapi-go.git]
- Docker Hub Image: [https://hub.docker.com/repository/docker/ameyster786/weatherapp/]
## Deploy to container registry 

 `docker push ameyster786/weatherapp:latest`

```






