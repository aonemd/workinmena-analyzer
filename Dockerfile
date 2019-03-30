FROM golang:1.12.1

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 3001

ENTRYPOINT ["/app/workinmena-analyzer"]
