FROM golang:latest as builder
WORKDIR /go/server
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags timetzdata -o ./app ./server.go

FROM alpine:3.15
WORKDIR /go/server
COPY --from=builder /go/server/app ./
ENTRYPOINT ["./app"]