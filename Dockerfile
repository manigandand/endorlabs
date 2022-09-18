FROM alpine:3.14.0 AS builder
# Set Locales to support UTF-8 encoding to support non ASCII characters i.e. emojis
ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8

RUN apk add ca-certificates curl
RUN mkdir /app
WORKDIR /app

COPY endorlabs endorlabs
