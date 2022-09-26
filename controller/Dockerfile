FROM ubuntu:18.04

EXPOSE 5000

WORKDIR /controller

RUN echo 'debconf debconf/frontend select Noninteractive' | debconf-set-selections

# Install dependencies
RUN apt-get update -y
RUN apt-get install -y curl git libaio1 libaio-dev openssl openssh-server
RUN curl -fsSL https://deb.nodesource.com/setup_lts.x | bash -
RUN apt-get install -y nodejs

# Setup MySQL & Controller
RUN git clone https://github.com/waverleylabs/SDPController.git
COPY config.js /controller/SDPController/config.js
RUN cd SDPController && npm install

# Generate certificates
RUN mkdir certs
RUN cd ~/; openssl rand -writerand .rnd
RUN echo "Generating Root Cert.."
RUN openssl genrsa -des3 -passout pass:antonio -out certs/ca.key 4096
RUN openssl req -new -x509 -days 365 -key certs/ca.key -out certs/ca.crt -passin pass:antonio -subj "/C=ES/ST=PA/L=A/O=Uniovi/OU=SE/CN=PhD/emailAddress=abc@xyz.com"

WORKDIR /controller

COPY . .
#CMD ["node", "./SDPController/sdpController.js"]
#CMD ["/bin/bash", "-c", "sleep infinity"]
CMD ["/bin/bash", "-c", "/controller/run.sh"]