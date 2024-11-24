FROM golang:1.23.2 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-expert-cloud-run cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/go-expert-cloud-run .
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT [ "./go-expert-cloud-run" ]