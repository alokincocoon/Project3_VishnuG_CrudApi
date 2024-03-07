# syntax=docker/dockerfile:1

FROM golang:1.21.6

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o student-management .

EXPOSE 8080

CMD ["./student-management"]