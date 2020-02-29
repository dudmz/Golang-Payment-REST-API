FROM golang:1.13.8

ARG MONGO_URI
ARG MONGO_DATABASE

ENV GOPATH /go
ENV MONGO_URI $MONGO_URI
ENV MONGO_DATABASE $MONGO_DATABASE

RUN go get -u github.com/golang/dep/...
COPY . /go/src/github.com/carlosdamazio/Stone-REST-API

WORKDIR /go/src/github.com/carlosdamazio/Stone-REST-API
RUN echo "*********DEPENDENCY MANAGEMENT PHASE*****"
RUN dep ensure

RUN echo "*********TESTING PHASE*******************"
RUN go test ./app ./app/model ./app/handler ./app/serializer

RUN echo "*********BUILD PHASE*********************"
RUN go build -o /go/api main.go

EXPOSE 8080
CMD ["/go/api"]

