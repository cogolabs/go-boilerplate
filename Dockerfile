FROM docker.mycompany.com/platform/golang:1.10.3

ADD . /go/src/git.mycompany.com/platform/go-boilerplate.git
WORKDIR /go/src/git.mycompany.com/platform/go-boilerplate.git
RUN go install ./...

CMD ["/go/bin/boilerd"]
