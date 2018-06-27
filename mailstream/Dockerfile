FROM golang
ADD . /go/src/github.com/katsuya94/shizuku/mailstream
RUN go get -u google.golang.org/api/gmail/v1
RUN go get -u golang.org/x/oauth2/...
RUN go install github.com/katsuya94/shizuku/mailstream
ENTRYPOINT /go/bin/mailstream
