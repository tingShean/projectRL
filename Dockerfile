FROM	golang:alpine AS build-env

RUN	mkdir -p /go/src/github.com/tingShean/projectRL
ADD	. /go/src/github.com/tingShean/projectRL/
WORKDIR /go/src/github.com/tingShean/projectRL/
RUN	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# final
FROM	scratch
COPY	--from=build-env /go/src/github.com/tingShean/projectRL/app /
ENTRYPOINT ["app"]
