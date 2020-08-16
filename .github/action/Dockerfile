# Pull base image
ARG BASE_IMAGE_TAG
FROM golang:${BASE_IMAGE_TAG:-1.15}

WORKDIR /root

RUN apt update && apt upgrade -y
RUN apt-get -y install libgl1-mesa-dev xorg-dev

ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
