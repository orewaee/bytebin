FROM golang:alpine AS build

ENV GOOS linux

WORKDIR /build

ADD go.mod .
COPY . .

RUN go build -o out/main -v cmd/bytebin/main.go

FROM alpine

WORKDIR /bytebin

COPY --from=build /build/out/main /bytebin/main

ENTRYPOINT ["./main"]
