FROM golang:1.20-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
RUN go mod download
COPY . .
RUN go build -o app .

FROM alpine:3.14
WORKDIR /root
COPY --from=build /app/app .
COPY .env.production .env
CMD /root/app

EXPOSE 8080