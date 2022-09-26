FROM ubuntu:18.04

EXPOSE 4444

WORKDIR /gateway

RUN echo 'debconf debconf/frontend select Noninteractive' | debconf-set-selections

USER root

# Install dependencies
RUN apt-get update -y
RUN apt-get install -y make automake curl git net-tools conntrack openssl libssl-dev libjson-c-dev libpcap-dev texinfo libtool autoconf libuv1 libuv1-dev
RUN apt-get install -y iptables

# Install FWKNOP
RUN git clone https://github.com/waverleylabs/fwknop.git
RUN cd fwknop && libtoolize --force && aclocal && autoheader && automake --force-missing --add-missing && autoconf
RUN cd fwknop && ./configure --disable-client --prefix=/usr --sysconfdir=/etc && make && make install

# Copy configuration files
COPY . .

CMD ["/bin/bash","-c", " ./run.sh $SERVER_HOSTNAME"]