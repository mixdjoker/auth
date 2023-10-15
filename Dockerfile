FROM golang:1.21-alpine3.18 AS builder

COPY . /app-src
WORKDIR /app-src

RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOEXPERIMENT="loopvar" go build -ldflags "-w -s" -o ./bin/auth cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app-src/bin/auth ./auth
COPY ./config ./config

ENV AUTH_CONFIG_PATH=./config/auth.yml
CMD [ "./auth" ]
