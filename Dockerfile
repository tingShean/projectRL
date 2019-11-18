FROM	golang:alpine AS build-env

ADD	. .
RUN	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# final
FROM	scratch
COPY	--from=build-env app /
ENTRYPOINT ["app"]
