######################################
# Prepare npm_builder
######################################
FROM node:10-alpine as npm_builder
WORKDIR /go/src/github.com/checkr/flagr
ADD . .
ARG FLAGR_UI_POSSIBLE_ENTITY_TYPES=null
ENV VUE_APP_FLAGR_UI_POSSIBLE_ENTITY_TYPES ${FLAGR_UI_POSSIBLE_ENTITY_TYPES}
RUN cd ./browser/flagr-ui/ && npm install && npm run build

######################################
# Prepare go_builder
######################################
FROM golang:1.12 as go_builder
WORKDIR /go/src/github.com/checkr/flagr
ADD . .
RUN make build

######################################
# Copy from builder to alpine image
######################################
FROM frolvlad/alpine-glibc:alpine-3.10
RUN apk add --no-cache curl
WORKDIR /go/src/github.com/checkr/flagr
VOLUME ["/data"]

ENV HOST=0.0.0.0
ENV PORT=18000
ENV FLAGR_DB_DBDRIVER=sqlite3
ENV FLAGR_DB_DBCONNECTIONSTR=/data/demo_sqlite3.db
ENV FLAGR_RECORDER_ENABLED=false

COPY --from=go_builder /go/src/github.com/checkr/flagr/flagr ./flagr
COPY --from=npm_builder /go/src/github.com/checkr/flagr/browser/flagr-ui/dist ./browser/flagr-ui/dist
ADD ./docs ./docs/
ADD ./buildscripts ./buildscripts
ADD ./buildscripts/demo_sqlite3.db /data/demo_sqlite3.db

EXPOSE 18000
CMD ./flagr
