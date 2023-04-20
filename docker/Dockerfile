FROM golang:1.20-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN CGO_ENABLED=0 go build -o /api

FROM gcr.io/distroless/base

COPY --from=build --chown=nonroot:nonroot /api /api

USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["/api"]
