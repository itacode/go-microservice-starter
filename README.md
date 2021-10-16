# Go microservice starter
It's a lean boilerplate to start developing a microservice in Go.

## Features
- Gin server
- Godotenv to load environment variables
- Makefile for the development and production tasks
- Air live reload configuration
- Docker configuration
- Docker-compose configuration

The project layout follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout).  
I adopted the [godotenv](https://pkg.go.dev/github.com/joho/godotenv) to load environment variables, following this [convention](https://github.com/bkeepers/dotenv#what-other-env-files-can-i-use).

## Development
In the Makefile there are the commands to develop and build:
- `make dev`: run the app with live reload (with Air)
- `make`: format, lint, examine and build
- `make start`: start the server in production mode

## Docker
Build
```sh
docker build -t go_app -f .\build\package\Dockerfile .
```
Docker compose up
```sh
docker-compose -f .\docker-compose.yml up -d --build
```
Docker compose down
```sh
docker-compose -f .\docker-compose.yml down
```

## Installation
To easily install the latest version of [GNU Make](https://www.gnu.org/software/make/) for Windows, you may execute `scoop install make` ([Scoop](https://scoop.sh/)) or `choco install make` ([Chocolatey](https://chocolatey.org/)).  
To install globally [Air](https://github.com/cosmtrek/air) for the live reload, execute `go install https://github.com/cosmtrek/air`.
