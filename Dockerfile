FROM golang:1.12.0-alpine3.9
RUN mkdir /app
ADD . /app
ADD . /go/src/github.com/ajdinahmetovic/go-rest
WORKDIR /app
RUN apk update && apk add git && go get -d github.com/gorilla/mux && go get -d github.com/lib/pq && go get github.com/ajdinahmetovic/item-service/db && go get github.com/dgrijalva/jwt-go && go get -d github.com/lib/pq && go get golang.org/x/crypto/bcrypt && go get google.golang.org/grpc && go get go.uber.org/zap
RUN go build -o main .
CMD ["/app/main"]