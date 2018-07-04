FROM golang
RUN go get -u google.golang.org/api/gmail/v1
RUN go get -u golang.org/x/oauth2/...
ADD . /go/src/github.com/katsuya94/shizuku/mailstream
WORKDIR /go/src/github.com/katsuya94/shizuku/mailstream
RUN go install github.com/katsuya94/shizuku/mailstream
ENTRYPOINT /go/bin/mailstream
