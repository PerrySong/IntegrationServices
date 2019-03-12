# Two-stage build:
#    first  FROM prepares a binary file in full environment ~780MB
#    second FROM takes only binary file ~10MB

FROM golang:1.9 AS builder

RUN go version

COPY . "/go/src/github.com/PerrySong/OAuth"
WORKDIR "/go/src/github.com/PerrySong/OAuth"

RUN go get -v -t  .


#RUN set -x && \
#    #go get github.com/2tvenom/go-test-teamcity && \
#    go get github.com/golang/dep/cmd/dep && \
#    dep ensure -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /your-app

CMD ["/your-app"]

EXPOSE 8080



#########
# second stage to obtain a very small image
FROM scratch

COPY --from=builder /your-app .

EXPOSE 8080

CMD ["/your-app"]
