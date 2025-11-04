Course Review Backend🕹️

Overview

This project is the backend service for a Course Review web application. It provides core APIs for managing users, reviews, classes, professors, and other related entities.

The system consists of:
- A PostgreSQL database running in a Docker container.
- A Go-based API service that interacts with the database to serve course review functionality.

Prerequisites

- Docker: https://www.docker.com/get-started
- Docker Compose: https://docs.docker.com/compose/install/
- Go (optional, for running the API outside Docker)

Setup

1. Clone the repository

   git clone https://github.com/B1gdawg0/course-review-backend
   cd course-review-backend

2. Create a .env file in the root directory with the following variables:

   PORT=8080
   DB_HOST=db
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=yourpassword
   DB_NAME=coursereview
   JWT_SECRET=yoursecret

3. Start the services using the setup script

   ./setup/init.sh

   This script will:
   - Initialize the PostgreSQL database
   - Build and run the Course Review API inside a Docker container

Project Structure

- /cmd          → Entry point for the application
- /setup        → Automation scripts for initialization
- /internal/... → Core application modules (ask me for descriptioin 😘)
