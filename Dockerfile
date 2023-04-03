FROM golang:1.20-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
RUN go mod download
COPY . .
RUN go test -coverprofile=coverage.out ./...
RUN go build -o app .
RUN go tool cover -func=coverage.out | awk '{print $3}' | awk -F. '{print $1}' | tail -n1 | awk '{if ($1 < 10) {exit 1}}'

FROM alpine:3.14
WORKDIR /root
COPY --from=build /app/app .
COPY .env.production .env
CMD /root/app

EXPOSE 8080