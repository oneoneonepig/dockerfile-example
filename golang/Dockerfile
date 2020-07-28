FROM golang:1.14 AS build

WORKDIR /go/src/main
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM alpine:latest AS runtime
COPY --from=build /go/src/main/app .
CMD ["./app"]
