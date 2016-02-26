FROM golang:1.6

EXPOSE 8080

ARG path="github.com/thomaso-mirodin/go-shorten"
# ARG package=app
# ${PWD#$GOPATH/src/}

RUN mkdir -p /go/src/${path}
COPY . /go/src/${path}

WORKDIR /go/src/${path}
RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]