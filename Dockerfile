FROM golang:1.13

RUN apt-get update

RUN mkdir /go/src/app
ADD . /go/src/app

WORKDIR /go/src/app

EXPOSE 3000