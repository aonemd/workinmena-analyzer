FROM golang:1.12.1 as build

WORKDIR /workinmena-analyzer

# cache go modules
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM alpine:3.12.1 as runtime

WORKDIR /app

COPY --from=build /workinmena-analyzer .

EXPOSE 3001

ENTRYPOINT ["/app/workinmena-analyzer"]
