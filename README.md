# Cloud-Native Go App - Production CI/CD Pipeline

A production-style cloud-native backend application built with **Go**, containerized using **Docker**, and deployed automatically to **AWS EC2** using **GitHub Actions CI/CD**.

This project demonstrates real-world DevOps and backend engineering practices including containerization, automated deployments, infrastructure integration, and service reliability.

This project demonstrates real-world backend engineering practices including:

- Clean architecture (handler → service → repository)

- Containerized deployment

- Database integration

- CI/CD pipeline

- Docker image registry

- Automated cloud deployment

- Infrastructure-level reliability patterns
  
---

## Project Overview

This project implements a full backend deployment pipeline:

Git Push → CI Build → Docker Image → DockerHub → AWS EC2 Deployment → Live API

It simulates a production backend workflow used in modern cloud environments.

---

## Architecture

Client Request → HTTP Handler (API Layer) → Service Layer (Business Logic) → Repository Layer → PostgreSQL Database

---

## Features

- REST API built with Go
- Clean backend architecture (Handler → Service → Repository)
- PostgreSQL database integration
- Docker containerization
- Multi-container deployment (App + DB)
- Structured JSON responses
- Health-ready backend design
- GitHub Actions CI/CD pipeline
- DockerHub image push automation
- AWS EC2 auto deployment via SSH
- Database readiness handling
- Production-style retry logic

---

## Tech Stack

- Language: Go

- Database: PostgreSQL

- Containerization: Docker

- CI/CD: GitHub Actions

- Container Registry: DockerHub

- Cloud Provider: AWS EC2

- Deployment: SSH-based auto deployment

- Architecture: Clean architecture pattern

---

## Required GitHub Secrets

- DOCKER_USERNAME
- DOCKER_PASSWORD
- EC2_HOST
- EC2_USER
- EC2_SSH_KEY
  
---

## Local Development Setup

- Clone repo
  - git clone https://github.com/<your-username>/cloud-native-go-app.git
    cd cloud-native-go-app

- Run with Docker Compose
  - docker compose up --build

- Test Api
  - http://localhost:8081/items
    
---

## Learning Outcomes

This project helped build practical experience with:

- Backend system design
- DevOps pipelines
- Cloud deployment workflows
- Docker-based architecture
- Production infrastructure concepts
