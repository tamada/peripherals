FROM alpine:3.10.1
LABEL maintainer="Haruai Tamada" \
      uniq2-version="1.0.2" \
      description="Deleting duplicate lines"

RUN    adduser -D uniq2 \
    && apk --no-cache add curl=7.66.0-r0 tar=1.32-r0 \
    && curl -s -L -O https://github.com/tamada/uniq2/releases/download/v1.0.2/uniq2-1.0.2_linux_amd64.tar.gz \
    && tar xfz uniq2-1.0.2_linux_amd64.tar.gz      \
    && mv uniq2-1.0.2 /opt                         \
    && ln -s /opt/uniq2-1.0.2 /opt/uniq2           \
    && ln -s /opt/uniq2/uniq2 /usr/local/bin/uniq2 \
    && rm uniq2-1.0.2_linux_amd64.tar.gz

ENV HOME="home/uniq2"

WORKDIR /home/uniq2
USER    uniq2

ENTRYPOINT [ "uniq2" ]
