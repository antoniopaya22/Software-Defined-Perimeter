#!/usr/bin/env bash

echo "Installing dependencies..."
apt-get update -y
apt-get -y install texinfo libtool autoconf make telnet openssl libssl-dev libjson0 libjson0-dev libpcap-dev git

echo "Install FWKNOP"
git clone https://github.com/waverleylabs/fwknop.git
cd fwknop
libtoolize --force
aclocal
autoheader
automake --force-missing --add-missing
autoconf
./configure -disable-client --prefix=/usr
make
make install