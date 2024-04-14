FROM golang:1.22.2-alpine as build

ARG CGO_ENABLED=0
WORKDIR /app

RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -C ./api

FROM scratch
COPY --from=build /app/api/api /api
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/api"]
