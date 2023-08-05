FROM alpine:3.10.1
ARG version="1.1.1"
LABEL maintainer="Haruai Tamada" \
      peripherals-version=${version} \
      description="Deleting duplicate lines"

RUN    adduser -D nonroot \
    && apk --no-cache add curl=7.66.0-r0 tar=1.32-r0 \
    && curl -s -L -O https://github.com/tamada/peripherals/releases/download/v${version}/peripherals-${version}_linux_amd64.tar.gz \
    && tar xfz peripherals-${version}_linux_amd64.tar.gz \
    && mv peripherals-${version} /opt                    \
    && ln -s /opt/peripherals-${version} /opt/uniq2      \
    && ln -s /opt/peripherals/ /usr/local/bin/uniq2 \
    && rm peripherals-${version}_linux_amd64.tar.gz

ENV HOME="home/nonroot"

WORKDIR /home/nonroot
USER    nonroot

ENTRYPOINT [ "uniq2" ]
