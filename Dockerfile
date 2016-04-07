FROM golang:latest
ADD . /code
WORKDIR /code
RUN go build test.go
#CMD go build test.go; ./test 123
CMD /bin/sh
