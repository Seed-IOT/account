
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/account
COPY . .
RUN apk add --no-cache git
RUN go get ./...
RUN go build -o account ./cmd/main.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash
COPY --from=builder /go/src/account/account /account
# COPY --from=builder /go/src/account/config/config.yml /config.yml

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh ./wait-for-it.sh
RUN ["chmod", "+x", "./wait-for-it.sh"]
LABEL Name=account
EXPOSE 8080
