
###################
# build stage
###################

FROM golang:1.21-alpine AS builder

# RUN apk add --no-cache git
# WORKDIR /go/src/app
# COPY . .

# RUN go get -d -v ./...
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /go/bin/app -v .

ARG CGO_ENABLED=0
ARG GOOS=linux

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -ldflags="-s -w" -o /go/bin/app -v .

###################
# development stage 
###################

FROM alpine:latest AS dev

RUN addgroup -S app && adduser -S app -G app
COPY --from=builder --chown=app /go/bin/app /app
USER app

ENV GIN_ENV development
ENV GIN_PORT 8080
ENV GIN_MODE debug

ENTRYPOINT ["/app"]

###################
# production stage 
###################

FROM alpine:latest AS PROD

RUN addgroup -S app && adduser -S app -G app
COPY --from=builder --chown=app /go/bin/app /app
USER app

ENV GIN_ENV production
ENV GIN_PORT 8080
ENV GIN_MODE release

ENTRYPOINT ["/app"] 