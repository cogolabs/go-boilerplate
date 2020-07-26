FROM docker.mycompany.com/platform/golang:1.14.6

ENV GODEBUG netdns=cgo
ENV TZ America/New_York

ADD . /go/src/git.mycompany.com/platform/go-boilerplate.git
WORKDIR /go/src/git.mycompany.com/platform/go-boilerplate.git
RUN go install ./...

CMD ["/go/bin/boilerd"]
