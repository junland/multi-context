FROM golang:1.8.5-alpine3.5

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o multi-context

WORKDIR /app

ENV PORT 8001

EXPOSE 8001

COPY --from=binary /app/multi-context /app

CMD ["/app/multi-context"]
