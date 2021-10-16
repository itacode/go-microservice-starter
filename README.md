# Go microservice starter
It's a lean boilerplate to start developing a microservice in Go.

## Features
- Gin server
- Godotenv to load environment variables
- Makefile for the development tasks
- Air live reload configuration
- Docker configuration
- Docker-compose configuration

The project layout follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout).  
I adopted the [godotenv](https://pkg.go.dev/github.com/joho/godotenv) to load environment variables, following this [convention](https://github.com/bkeepers/dotenv#what-other-env-files-can-i-use).

## Developing
In the Makefile there are the commands to develop and build:
- `make dev`: run the app with live reload (with Air)
- `make`: format, lint, examine and build
- `make start`: start the server in production mode

## Installation
To easily install [GNU Make](https://www.gnu.org/software/make/) in Windows, you may execute `scoop install make` ([Scoop](https://scoop.sh/)) or `choco install make` ([Chocolatey](https://chocolatey.org/)).  
To install globally [Air](https://github.com/cosmtrek/air) for the live reload, execute `go install https://github.com/cosmtrek/air`.
