FROM alpine:latest

RUN apk --no-cache add curl

WORKDIR /app
COPY serv-1 . 

ENTRYPOINT /app/serv-1