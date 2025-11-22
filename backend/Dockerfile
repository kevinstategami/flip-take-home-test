FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o server ./cmd/server

FROM debian:bookworm-slim

WORKDIR /app
COPY --from=build /app/server .

EXPOSE 8080
CMD ["./server"]
