FROM golang:1.20-alpine AS build
WORKDIR /src
COPY main.go .
RUN go build -o /bin/webapp main.go

FROM alpine:3.18
COPY --from=build /bin/webapp /usr/local/bin/webapp
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/webapp"]
