FROM golang:1.16 AS build

WORKDIR /app

COPY . .
RUN go build -o app

FROM alpine:3.14
WORKDIR /app

COPY --from=build /app/app .
CMD ["./app"]