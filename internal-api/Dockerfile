FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o server main.go

FROM scratch
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]
