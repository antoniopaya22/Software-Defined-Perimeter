FROM ubuntu:18.04

WORKDIR /gateway

RUN echo 'debconf debconf/frontend select Noninteractive' | debconf-set-selections

# Install dependencies
RUN apt-get update -y
RUN apt-get install -y make automake curl git net-tools conntrack openssl libssl-dev libjson-c-dev libpcap-dev texinfo libtool autoconf libuv1 libuv1-dev

# Install FWKNOP
RUN git clone https://github.com/waverleylabs/fwknop.git
RUN cd fwknop && libtoolize --force && aclocal && autoheader && automake --force-missing --add-missing && autoconf
RUN cd fwknop && ./configure --disable-server --prefix=/usr --sysconfdir=/etc && make && make install

# Copy configuration files
COPY . /root

RUN chmod 0600 /root/.fwknoprc

# Configure iptables

#RUN iptables -P INPUT DROP
#RUN iptables -P FORWARD DROP
#RUN iptables -S

# Start fwknopd gateway
#CMD ["/usr/sbin/fwknopd", "-n", "service_gate"]
CMD ["/bin/bash", "-c", "sleep infinity"]