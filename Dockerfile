# Dockerfile
# FROM docker.io/centos
FROM docker.io/dockergogolj/centos:7.6.1810-tools
MAINTAINER lijian <lijian678@yeah.net>
COPY test /test
ENTRYPOINT ["/test"]
EXPOSE 8080
#End