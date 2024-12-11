FROM golang:1.23 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY .env.example .env

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-mate ./cmd/app/main.go

FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/.env .env

COPY --from=build-stage /go-mate /go-mate

USER nonroot:nonroot

ENTRYPOINT ["/go-mate"]