FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.sum go.mod ./
RUN go mod download

COPY . .
RUN go build -o cmd ./cmd/main.go

FROM alpine

COPY --from=builder . /bin/cmd

ENTRYPOINT ["/bin/app"]