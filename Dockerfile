FROM golang:latest

ARG GPAY_DIR=/gpay
ENV GPAY_DIR=$GPAY_DIR

RUN apt-get update -y && apt-get upgrade -y \
    && apt install build-essential git -y \
    && mkdir -p /gpay

WORKDIR ${GPAY_DIR}
COPY . .
RUN make gpay-all

ENV SHELL /bin/bash
EXPOSE 8545 8546 8547 30303 30303/udp

ENTRYPOINT ["gpay"]
