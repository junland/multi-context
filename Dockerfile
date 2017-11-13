FROM golang:1.8.5-alpine3.5

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o multi-context

WORKDIR /app

ENV PORT 8081

EXPOSE 8081

ENTRYPOINT /app/multi-context
