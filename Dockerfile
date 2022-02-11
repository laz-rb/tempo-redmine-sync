FROM golang:1.17-alpine3.15 AS builder
COPY . /build/
WORKDIR /build/
RUN CGO_ENABLED=0 GOOS=linux go build -a -o tempo-redmine-sync ./cmd/tempo-redmine-sync

FROM alpine:3.15
COPY --from=builder /build/tempo-redmine-sync /app/tempo-redmine-sync
WORKDIR /app
ENTRYPOINT [ "./tempo-redmine-sync" ]