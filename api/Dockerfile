FROM golang:1.20-alpine as builder
RUN apk add alpine-sdk
WORKDIR /go/app
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o api -tags musl

FROM alpine:latest as runner
WORKDIR /root/
COPY --from=builder /go/app/api .
ENTRYPOINT /root/api
EXPOSE 8000