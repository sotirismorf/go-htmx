FROM node:22-alpine as build-styles

WORKDIR /build

COPY . .

RUN npm install && npm run build

FROM golang:1.22-alpine AS build-go

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

FROM alpine:edge

WORKDIR /app

COPY --from=build-go /build/myapp .

EXPOSE 8080

ENTRYPOINT ["/app/myapp"]
