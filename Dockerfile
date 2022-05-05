FROM golang:1.18 as builder
WORKDIR /build
COPY . .
RUN go run generator.go

FROM caddy:2.4.6
COPY --from=builder /build/public /usr/share/caddy
COPY --from=builder /build/Caddyfile /etc/caddy/Caddyfile
