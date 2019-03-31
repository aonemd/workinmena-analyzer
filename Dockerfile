FROM golang:1.12.1 as build

COPY . /workinmena-analyzer

WORKDIR /workinmena-analyzer

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM alpine:3.9.2 as runtime

WORKDIR /app

COPY --from=build /workinmena-analyzer .

EXPOSE 3001

ENTRYPOINT ["/app/workinmena-analyzer"]
