FROM amazonlinux:latest

RUN yum update -y
RUN yum install -y zip go

RUN mkdir -p /lambda /go

ENV GOPATH /lambda/go

WORKDIR /lambda/code

CMD /lambda/code/build.sh
