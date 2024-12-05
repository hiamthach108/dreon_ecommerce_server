FROM golang:1.22-alpine AS builder

ENV GO111MODULE 'on'
ENV CGO_ENABLED=0
ENV GOPROXY 'https://proxy.golang.org,direct'
ENV GOPRIVATE 'gitlab.com/together-work3/source/*'
ENV GOSUMDB 'off'

WORKDIR /app

RUN apk add upx

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags "-s -w" -o ./build/bin/main ./cmd/main.go
RUN upx -9 /app/build/bin/main

# Create final image
FROM gcr.io/distroless/static:latest
WORKDIR /app
COPY --from=builder /app/build/bin/main main
# COPY --from=builder /app/.env .env
CMD ["/app/main"]