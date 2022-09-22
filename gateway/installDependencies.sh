#!/usr/bin/env bash

echo "Installing dependencies..."
apt-get update -y
apt-get install -y curl git net-tools conntrack openssl libssl-dev libjson-c-dev libpcap-dev texinfo libtool autoconf libuv1 libuv1-dev

echo "Install FWKNOP"
git clone https://github.com/waverleylabs/fwknop.git
cd fwknop
libtoolize --force
aclocal
autoheader
automake --force-missing --add-missing
autoconf
./configure --disable-client --prefix=/usr --sysconfdir=/etc
make
make install