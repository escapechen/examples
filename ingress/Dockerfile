FROM golang:alpine
ENV GOPATH /go
RUN apk add --update git
WORKDIR /go
COPY . ~/go/src/github.com/escapechen/examples/ingress
RUN go get -v github.com/escapechen/examples/ingress

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/bin/ingress /usr/local/bin
CMD [ "ingress" ]
