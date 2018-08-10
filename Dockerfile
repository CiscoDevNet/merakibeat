FROM alpine:3.5

MAINTAINER Neelesh Pateriya

LABEL Description="Meraki health Beats plugin"

RUN apk update && \
    apk upgrade && \
    apk add \
        bash \
        ca-certificates \
    && rm -rf /var/cache/apk/*

RUN mkdir /plugin
COPY merakibeat.yml /plugin/
COPY fields.yml /plugin/
COPY ./release/linux/amd64/merakibeat /plugin/merakibeat

WORKDIR /plugin

ENTRYPOINT ["/plugin/merakibeat", "-e", "-d", "*"]
