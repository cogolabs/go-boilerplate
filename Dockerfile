FROM docker.mycompany.com/platform/golang:1.18.1

# @TODO: specify the correct timezone for your service, or remove
ENV TZ America/New_York

ENV GODEBUG netdns=cgo

ADD . /go/src/git.mycompany.com/platform/go-boilerplate.git
WORKDIR /go/src/git.mycompany.com/platform/go-boilerplate.git
RUN go install ./...

CMD ["/go/bin/boilerd"]
