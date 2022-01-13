# creating a Docker build environment for compile.
from golang:alpine as build-env
workdir /app
add . /app

# creating the Docker runtime image, using alpine image.
from alpine
# install CA-certs for HTTPS
run apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
workdir /app
copy --from=build-env /app/goserver /app

EXPOSE 8030

ENTRYPOINT ./goserver

