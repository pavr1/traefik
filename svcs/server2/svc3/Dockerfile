#build stage
FROM golang:alpine AS build-env
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build-env /go/bin/app /app
ENTRYPOINT ["/app"]
LABEL Name=traefik Version=0.0.1

EXPOSE 9002
