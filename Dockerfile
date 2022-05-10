FROM golang:1.12-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /xpaymentsorg
RUN cd /xpaymentsorg && make XPS

FROM alpine:latest

WORKDIR /xpaymentsorg

COPY --from=builder /xpaymentsorg/build/bin/XPS /usr/local/bin/XPS

RUN chmod +x /usr/local/bin/XPS

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/usr/local/bin/XPS"]

CMD ["--help"]
