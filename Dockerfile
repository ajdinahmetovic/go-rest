FROM golang:1.12.0-alpine3.9
RUN mkdir /app
ADD . /app
ADD . /go/src/github.com/ajdinahmetovic/go-rest
WORKDIR /app
RUN apk update && apk add git && go get -d github.com/gorilla/mux && go get -d github.com/lib/pq
RUN go build -o main .
CMD ["/app/main"]