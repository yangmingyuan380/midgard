# Copyright 2020-2021 Changkun Ou. All rights reserved.
# Use of this source code is governed by a GPL-3.0
# license that can be found in the LICENSE file.

FROM chromedp/headless-shell:latest AS builder-env
WORKDIR /app
COPY . .
RUN apt update && apt install -y wget gcc
RUN mkdir -p /root/goes
ARG GOVERSION=go1.21.5
RUN cd /root/goes && wget https://dl.google.com/go/$GOVERSION.linux-amd64.tar.gz
RUN cd /root/goes && tar xvf $GOVERSION.linux-amd64.tar.gz && rm $GOVERSION.linux-amd64.tar.gz
RUN cd /root/goes && mv /root/goes/go /root/goes/$GOVERSION
RUN cd /root/goes && ln -s /root/goes/$GOVERSION /root/goes/go
RUN cd /root/goes && export GOROOT=~/goes/go
RUN CGO_ENABLED=0 /root/goes/go/bin/go build -mod=vendor

FROM chromedp/headless-shell:latest
RUN apt update && apt install -y dumb-init git
ENTRYPOINT ["dumb-init", "--"]

WORKDIR /app
COPY . .
COPY --from=builder-env /app/midgard /app/mg
EXPOSE 9880
CMD ["/app/mg", "server"]
