# ---- build: Waverley Labs fwknop fork (SDP-enabled), pinned commit ----
# Same build stage as gateway/Dockerfile so BuildKit can share the cache.
FROM debian:bookworm-slim AS build
RUN apt-get update && apt-get install -y --no-install-recommends \
      ca-certificates git build-essential automake autoconf libtool \
      libssl-dev libjson-c-dev libpcap-dev libuv1-dev iptables \
 && rm -rf /var/lib/apt/lists/*

ARG FWKNOP_REPO=https://github.com/waverleylabs/fwknop.git
ARG FWKNOP_COMMIT=db2f4d07b3bb40f1febd3567cd1a1aa07f26a4e0
RUN git clone "$FWKNOP_REPO" /src && cd /src && git checkout --quiet "$FWKNOP_COMMIT"
WORKDIR /src
RUN sed -i -e 's/\$(SERVER_DIR) \\$/$(SERVER_DIR)/' -e '/^    doc$/d' Makefile.am \
 && libtoolize --force && aclocal && autoheader \
 && automake --force-missing --add-missing && autoconf \
 && ./configure --prefix=/usr --sysconfdir=/etc --localstatedir=/var \
      --disable-server CFLAGS="-O2 -fcommon" \
 && make -j"$(nproc)" \
 && make install DESTDIR=/out

# ---- runtime ----
FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y --no-install-recommends \
      ca-certificates curl iproute2 libssl3 libjson-c5 libuv1 \
 && rm -rf /var/lib/apt/lists/*
COPY --from=build /out/usr/bin/fwknop /usr/bin/fwknop
COPY --from=build /out/usr/lib/libfko.so* /usr/lib/
RUN ldconfig

COPY config/ /etc/sdp/templates/
COPY entrypoint.sh /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["sleep", "infinity"]
