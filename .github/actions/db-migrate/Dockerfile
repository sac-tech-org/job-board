FROM golang:1.22.2 as build

ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -C ./db -o migrate

FROM scratch
COPY --from=build /app/db/migrate /migrate

ENTRYPOINT ["/migrate"]
