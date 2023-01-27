FROM golang:1.18.3-alpine3.15 as BUILDER

WORKDIR /go/src/app
COPY . .

RUN apk add build-base

RUN go install -v ./...
RUN go build -o ./build/build ./cmd/main.go

FROM alpine:latest
RUN mkdir /app
WORKDIR /app
 
COPY --from=BUILDER /go/src/app/build/build .

CMD ["./build"]