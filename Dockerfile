FROM golang:1.22.2-alpine3.18 AS build
WORKDIR /src
COPY main.go go.mod go.sum .
RUN go mod download
RUN go mod tidy
RUN go build -o /app/webserver ./main.go

FROM alpine:latest
WORKDIR /app
COPY templates templates/
COPY static static/
COPY langs langs/
COPY config.yaml .
COPY --from=build /app/webserver .
EXPOSE 80
CMD ["./webserver"]
