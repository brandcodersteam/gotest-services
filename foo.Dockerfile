# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.16-alpine
RUN apk add git

ENV GO111MODULE=on \
  GOPRIVATE=github.com/brandcodersteam/gotest-services


WORKDIR /go/src/gotest-services

# # Download necessary Go modules
# COPY go.mod ./
# # COPY go.sum ./
# RUN go mod download

# # Copy go code
# COPY *.go ./

# # Build user service
# RUN go build -o /user-service

COPY ./services/foo/go.mod .
COPY ./services/foo/go.sum .
RUN go mod download

COPY ./services/foo/*.go .

RUN go build foo.go

EXPOSE 8080

CMD [ "./foo" ]
