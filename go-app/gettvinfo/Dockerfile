FROM golang:latest as builder
WORKDIR /go/gettvinfo
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags timetzdata -o ./app ./cmd/main.go

FROM alpine:3.15
WORKDIR /go/gettvinfo
COPY --from=builder /go/gettvinfo/app /go/gettvinfo/go.env ./
RUN apk add chromium
ENTRYPOINT ["./app"]