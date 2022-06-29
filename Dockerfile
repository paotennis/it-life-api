FROM golang:1.18-alpine3.16 as builder

RUN apk update
RUN apk add curl

WORKDIR /go/src

# Install Go modules.
COPY ./it_life_api/go.mod ./
COPY ./it_life_api/go.sum ./
RUN go mod download

COPY ./it_life_api/ ./

RUN GOOS=linux GOARCH=amd64 go build -o /main ./cmd

FROM alpine:3.16

COPY --from=builder /main .

ENV PORT=${PORT}

ENTRYPOINT ["/main"]
