# Securing the Software-Defined Perimeter framework with automated security configuration deployment systems

## Introduction

A

## Architecture

A

## Installation

A

### SDP Controller

- Install and prepare the SDP Controller:

```bash
$ cd controller && ./installDependencies.sh

$ git clone https://github.com/waverleylabs/SDPController.git && yes | cp -rf config.js SDPController/config.js  && cd SDPController && npm install

$ service mysql start && mysql -u root -e "create database sdp;" && mysql -u root sdp < setup.sql && mysql -u root sdp < data.sql

$ ./Gencert.sh
```

- Start the SDP Controller:

```bash
$ node ./SDPController/sdpController.js
```

### SDP Gateway

- Install and prepare the SDP Gateway:

```bash
$ cd gateway && ./installDependencies.sh

$ cp config/access.conf /etc/fwknop/ && cp config/fwknopd.conf /etc/fwknop/ && cp config/gate_sdp_ctrl_client.conf /etc/fwknop/ && cp config/gate.fwknoprc /etc/fwknop/
```

- Start the SDP Gateway:

```bash
$ fwknopd
```

### SDP Client

- Install and prepare the SDP Client:

```bash
$ cd client && ./installDependencies.sh
$ cp .fwknoprc /home/antonio && cp sdp_ctrl_client.conf /home/antonio
```

- Start the SDP Client:

```bash
$ fwknopd -n service_gate
```