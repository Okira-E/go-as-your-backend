FROM golang:1.23-alpine AS builder

# Install just
RUN apk add --no-cache curl bash
RUN apk add --no-cache git just

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 just build

FROM scratch

COPY --from=builder /app/bin/server /server-app

EXPOSE 3200

ENTRYPOINT ["/server-app"]