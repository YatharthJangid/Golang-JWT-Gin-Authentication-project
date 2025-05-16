# Golang JWT Gin Authentication Demo

This project exposes a small JWT-based authentication API built with Gin and MongoDB. It is configured to run locally with `.env` and deploy as a Render web service.

## Demo Features

- User signup and login with JWT + refresh token generation
- Bearer token authentication middleware
- Protected endpoints for profile lookup and user listing
- Health endpoint for deployment monitoring
- Render blueprint file for one-click service setup

## Required Environment Variables

Copy `.env.example` to `.env` for local work.

```env
PORT=10000
MONGODB_URL=mongodb://localhost:27017
DB_NAME=go_jwt_demo
SECRET_KEY=replace-with-a-long-random-string
```

## Run Locally

```bash
go run .
```

## API

- `GET /` service info
- `GET /healthz` deployment health check
- `POST /users/signup`
- `POST /users/login`
- `GET /users/me` requires `Authorization: Bearer <token>`
- `GET /users/:user_id` requires auth
- `GET /users` requires an admin token

## Render Deployment

1. Push this repository to GitHub.
2. In Render, create a new Blueprint or Web Service from the repository.
3. Set `MONGODB_URL` and `SECRET_KEY` in Render.
4. Deploy. Render will build with `go build -o app .` and check `GET /healthz`.

For a quick demo, use Render MongoDB-compatible infrastructure or an external MongoDB provider and point `MONGODB_URL` at it.
