FROM golang:1.16-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download
RUN go build cmd/main/main.go

## Deploy
FROM alpine

WORKDIR /
RUN ls -la
COPY --from=build /app/main /main
RUN ls /main
EXPOSE 8000

ENTRYPOINT ["/main"]