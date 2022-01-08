FROM golang:1.17 AS build-go
COPY . /src
RUN cd /src && go build -x -o server

FROM ubuntu:latest
RUN apt update && apt install firefox -y
COPY --from=build-go /src/server /bin/server

ENTRYPOINT [ "server" ]