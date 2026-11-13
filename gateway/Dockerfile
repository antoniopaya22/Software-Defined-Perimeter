# ---- build: Waverley Labs fwknop fork (SDP-enabled), pinned commit ----
FROM debian:bookworm-slim AS build
RUN apt-get update && apt-get install -y --no-install-recommends \
      ca-certificates git build-essential automake autoconf libtool \
      libssl-dev libjson-c-dev libpcap-dev libuv1-dev iptables \
 && rm -rf /var/lib/apt/lists/*

ARG FWKNOP_REPO=https://github.com/waverleylabs/fwknop.git
ARG FWKNOP_COMMIT=db2f4d07b3bb40f1febd3567cd1a1aa07f26a4e0
RUN git clone "$FWKNOP_REPO" /src && cd /src && git checkout --quiet "$FWKNOP_COMMIT"
WORKDIR /src
# The texinfo manual does not build with current texinfo, and it is not needed.
# -fcommon: the code relies on tentative definitions rejected by GCC >= 10.
RUN sed -i -e 's/\$(SERVER_DIR) \\$/$(SERVER_DIR)/' -e '/^    doc$/d' Makefile.am \
 && libtoolize --force && aclocal && autoheader \
 && automake --force-missing --add-missing && autoconf \
 && ./configure --prefix=/usr --sysconfdir=/etc --localstatedir=/var \
      --disable-client --with-iptables=/usr/sbin/iptables CFLAGS="-O2 -fcommon" \
 && make -j"$(nproc)" \
 && make install DESTDIR=/out

# ---- runtime ----
FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y --no-install-recommends \
      ca-certificates iptables iproute2 conntrack libssl3 libjson-c5 libpcap0.8 libuv1 \
 && rm -rf /var/lib/apt/lists/*
COPY --from=build /out/usr/sbin/fwknopd /usr/sbin/fwknopd
COPY --from=build /out/usr/lib/libfko.so* /usr/lib/
COPY --from=build /out/etc/fwknop/access.conf /etc/fwknop/access.conf
RUN ldconfig && mkdir -p /var/run/fwknop

COPY config/ /etc/sdp/templates/
COPY entrypoint.sh /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
