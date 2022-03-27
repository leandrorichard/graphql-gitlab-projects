# Dockerfile References: https://docs.docker.com/engine/reference/builder/
FROM golang:1.17 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

FROM alpine:latest

ENV GITLAB_GRAPHQL_API_URL="https://gitlab.com/api/graphql"

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the built binary file from the previous stage. #
COPY --from=builder /app/main .

# Runs the application.
CMD ["./main"]
