FROM golang:1.11.4

WORKDIR /go/src/github.com/naoki-kishi/go-api-sample
COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]