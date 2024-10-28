FROM golang:1.23-alpine AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app
COPY . .

RUN go mod download && \
  go build -ldflags "-w -s" -o main /app/cmd/main.go

FROM scratch AS final

WORKDIR /app
COPY --from=builder /app/main /app/

ENTRYPOINT [ "/app/main" ]
