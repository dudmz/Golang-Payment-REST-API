FROM golang:1.13.8

ENV GOPATH /go

RUN go get -u github.com/golang/dep/...
COPY . /go/src/github.com/carlosdamazio/Stone-REST-API
RUN cd /go/src/github.com/carlosdamazio/Stone-REST-API \
 && dep ensure && go build -o /go/api main.go

EXPOSE 8080

CMD ["/go/api"]

