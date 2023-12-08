FROM --platform=$BUILDPLATFORM docker.io/golang:1.21-alpine AS build

WORKDIR /src

RUN apk --no-cache add ca-certificates

COPY . .

ARG TARGETOS TARGETARCH

RUN go mod download && \
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /src/gocatgo.bin

FROM scratch as bin

WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/gocatgo.bin .

EXPOSE 8080

CMD ["/app/gocatgo.bin"]
