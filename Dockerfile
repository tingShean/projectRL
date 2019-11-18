FROM	golang:alpine AS build-env

RUN	mkdir -p github.com/tingShean/projectRL
ADD	. github.com/tingShean/projectRL/
RUN	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# final
FROM	scratch
COPY	--from=build-env app /
ENTRYPOINT ["app"]
