FROM golang:1.6

EXPOSE 8080

ARG package=github.com/thomaso-mirodin/shawty
# ARG package=app 
# ${PWD#$GOPATH/src/}

RUN mkdir -p /go/src/${package}
WORKDIR /go/src/${package}

COPY . /go/src/${package}
RUN go-wrapper download
RUN go-wrapper install


CMD ["go-wrapper", "run"]