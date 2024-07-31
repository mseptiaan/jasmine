FROM golang:1.22-alpine AS builder

WORKDIR /build
COPY . ./
RUN go mod download

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o jasmine-server .

FROM alpine:3.19.1

COPY --from=builder /build/jasmine-server /opt/apps/jasmine-server

WORKDIR /opt/apps
EXPOSE 8080

CMD [ "/opt/apps/jasmine-server" ]