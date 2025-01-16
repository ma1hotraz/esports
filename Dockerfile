# Stage 1: Build the React app
FROM node:lts-slim AS frontend-builder
WORKDIR /app/js
COPY js/package.json js/pnpm-lock.yaml ./
RUN npm install -g pnpm && pnpm install
COPY js .
RUN pnpm build:prod

# Stage 2: Build the Go app
FROM golang:alpine3.20 AS backend-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /esportsdifference

# Stage 3: Final container with the built React and Go apps
FROM alpine:latest
ENV VIEWS_DIR=/app/views \
    LOGS_DIR=/app/logs \
    PUBLIC_DIR=/app/public
COPY public /app/public 
COPY views /app/views
COPY --from=backend-builder /esportsdifference /app
COPY --from=frontend-builder /app/public/js /app/public/js
EXPOSE 3000
ENTRYPOINT ["/app/esportsdifference"]
