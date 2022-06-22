FROM golang:1.18-alpine3.16

WORKDIR /go/src

# Install Go modules.
COPY ./it_life_api/go.mod ./
RUN go mod download
