FROM gitpod/workspace-full:build-branch-master

USER root

RUN apt-get update && apt-get upgrade -y
RUN apt-get -y install --no-install-recommends ca-certificates=20210119~20.04.1
RUN apt-get -y install --no-install-recommends libgl1-mesa-dev=20.2.6-0ubuntu0.20.04.1 xorg-dev=1:7.7+19ubuntu14 gcc=4:9.3.0-1ubuntu2

RUN apt-get -y install --no-install-recommends docker=1.5-2 docker-compose=1.25.0-1
RUN apt-get clean && rm -rf /var/cache/apt/* && rm -rf /var/lib/apt/lists/* && rm -rf /tmp/*

USER gitpod
