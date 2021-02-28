FROM gitpod/workspace-full:latest

USER root
# Install custom tools, runtime, etc.
RUN apt-get update && apt-get install -y \
    && apt-get install git-lfs -y \
    && apt-get install --no-install-recommends \
        # apt-utils \
        ca-certificates \
        libgl1-mesa-dev \
        xorg-dev \
        gcc \
        docker \
        docker-compose -y \
    && apt-get clean && rm -rf /var/cache/apt/* && rm -rf /var/lib/apt/lists/* && rm -rf /tmp/*

# Give back control
USER root
