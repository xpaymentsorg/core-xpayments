FROM golang:1.10-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /XPSchain
RUN cd /XPSchain && make XPS

FROM alpine:latest

LABEL maintainer="anil@xinfin.org"

WORKDIR /XPSchain

COPY --from=builder /XPSchain/build/bin/XPS /usr/local/bin/XPS

RUN chmod +x /usr/local/bin/XPS

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/usr/local/bin/XPS"]

CMD ["--help"]
