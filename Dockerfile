FROM golang:alpine AS build

ENV GOOS linux

WORKDIR /build

ADD go.mod .
COPY . .

RUN go build -o bytebin -v cmd/bytebin/main.go

FROM alpine

WORKDIR /bytebin

COPY --from=build /build/bytebin /bytebin/bytebin

CMD ["./bytebin"]
