# Makefile for Apod Application

# Variables
APP_NAME := app

# Docker Compose file
DOCKER_COMPOSE_FILE := docker-compose.yml

# Build Docker images
docker-build:
	@echo "Building Docker images..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) build

# Run the application using Docker Compose
docker-run:
	@echo "Running the application using Docker Compose..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) up

# Stop and remove Docker containers
docker-stop:
	@echo "Stopping and removing Docker containers..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) down


