## Builder
FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk --no-cache --update add git make

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN make build

## Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --no-cache --update add ca-certificates tzdata && \
    mkdir /app

WORKDIR /app

EXPOSE 10000

COPY --from=builder /app/actions /app/actions

RUN ls -al

CMD ["./actions"]