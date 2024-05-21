FROM golang:1.22.3-alpine as golang

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify
RUN go build cmd/main.go

FROM gcr.io/distroless/static-debian11

COPY --from=golang /app .

EXPOSE 3000

CMD ["/app"]