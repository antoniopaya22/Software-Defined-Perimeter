# Software-Defined Perimeter

![GitHub LICENSE](https://img.shields.io/github/license/antonioalfa22/Software-Defined-Perimeter?color=green&label=license&logoColor=green)
![Version](https://img.shields.io/badge/version-1.0.0-green)
![Environment](https://img.shields.io/badge/environment-docker-blue)

## Description

This repository contains the code and resources necessary to deploy an architecture based on Software-Defined Perimeters using the open-source implementation made by [WaverleyLabs](https://www.waverleylabs.com/services/software-defined-perimeter-panther-sdp-implementation/). The deployment can be performed both on real machines and devices as well as on a virtualized network architecture using Docker and Docker-Compose to test the functionality of the SDPs.

## Architecture & Components

![SDP Architecture](docs/BasicSDP.svg)

- **SDP Controller**: The SDP Controller uses the [WaverleyLabs - SDP Controller](https://github.com/waverleylabs/SDPController.git) repository and contains a set of scripts and resources to automate its configuration and deployment. For more information on SDP Controller, see the following sites:
  - [http://www.waverleylabs.com/services/software-defined-perimeter](http://www.waverleylabs.com/services/software-defined-perimeter/)
  - [https://cloudsecurityalliance.org/group/software-defined-perimeter/](https://cloudsecurityalliance.org/group/software-defined-perimeter/)

- **SDP Gateway**: The SDP Gateway uses the open source software [fwknop](https://github.com/mrash/fwknop) which implements an authorization scheme known as Single Packet Authorization (SPA) for strong service concealment. SPA requires only a single packet which is encrypted, non-replayable, and authenticated via an HMAC in order to communicate desired access to a service that is hidden behind a firewall in a default-drop filtering stance. The SDP Gateway uses the iptables firewall to deny all requests to the Accepting Hosts (AH) by default, authorizing only those that have been authorized by the SDP Controller.

- **SDP Client (IH)**: The SDP client or initiating host is responsible for requesting the necessary authorization from the SDP Controller to access the Accepting Hosts (AH) protected by the SDP Gateway.

![SPA](docs/rsz_spa.png)

## Installation and deployment

To deploy in a Docker test environment you must first have Docker and Docker-Compose installed on your system. The next steps are to define the network architecture you want to set up and which services you want to protect with the gateways.

### Configure SDP-Controller

The SDP Controller uses a MySQL database to store the users, services, permissions, etc. In the database folder, there is included a file called `setup.sql` with the function of creating all the necessary tables and relations in the database.

> This file includes in lines 693-704 some example data added:
```sql
INSERT INTO sdpid(sdpid,valid,type,country,state,locality,org,org_unit,alt_name,email,serial) VALUES (1,1,'gateway','ES', 'PA', 'A', 'Uniovi', 'SE', 'PhD', 'abc@xyz.com', 'abc123');
INSERT INTO sdpid(sdpid,valid,type,country,state,locality,org,org_unit,alt_name,email,serial) VALUES (2,1,'controller','ES', 'PA', 'A', 'Uniovi', 'SE', 'PhD', 'abc@xyz.com', 'abc123');
INSERT INTO sdpid(sdpid,valid,type,country,state,locality,org,org_unit,alt_name,email,serial) VALUES (3,1,'client','ES', 'PA', 'A', 'Uniovi', 'SE', 'PhD', 'abc@xyz.com', 'abc123');
INSERT INTO service VALUES(1,'controller', 'controller' );
INSERT INTO service VALUES(2,'tcp', 'tcp' );
INSERT INTO service VALUES(3,'udp', 'udp' );
INSERT INTO service_gateway (id,service_id, gateway_sdpid, protocol, port, nat_ip, nat_port) VALUES(1, 1, 1 ,'tcp',5000, 'controller', 5000);
INSERT INTO service_gateway (id,service_id, gateway_sdpid, protocol, port, nat_ip, nat_port) VALUES(2, 2, 1 ,'tcp',4444, 'gateway', 4444);
INSERT INTO sdpid_service VALUES (1, 1, 1);
INSERT INTO sdpid_service VALUES (2, 1, 2);
INSERT INTO sdpid_service VALUES(3, 3, 1);
INSERT INTO sdpid_service VALUES(4, 3, 2);
```

> In the example data, a controller listening on port 5000 and a gateway allowing authorized connections on port 4444 is added.

In the folder named `controller` you will find both the Dockerfile for the SDP Controller deployment and the following files:

- `config.js`: Contains the SDP Controller configuration variables such as port, database connection or certificate locations.
- `run.sh`: Generates the credentials and launches the SDP Controller process in listening mode.

### Configure SDP-Gateway

The `gateway` folder contains both the Dockerfile for the SDP Gateway deployment and a folder called `config` with the following configuration files that must be edited to adapt to the designed architecture:

- `access.conf`: This file defines how fwknopd will modify firewall access controls for specific IPs/networks.

- `fwknopd.conf`: This is the configuration file for fwknopd, the Firewall Knock Operator daemon (SPA).

- `gate_sdp_ctrl_client.conf`: This file defines the locations of the certificates and how to connect to the SDP Controller.

- `gate.fwknoprc`: fwknoprc stanzas configuration file

### Configure SDP-Client

The `client` folder contains a Dockerfile for the SDP Client deployment and the following configuration files for communication and authorization with the SDP Controller:

- `.fwknoprc`: This file defines the services protected by the SDP Gateway that the client wants to access.

- `sdp_ctrl_client.conf`: This file defines the locations of the certificates and how to connect to the SDP Controller.

### Start

1. Modify the `docker-compose.yml` file to deploy the desired network architecture and run Docker-Compose. 

**Important**: The SDP Controller is the one that generates the certificates, you must copy the necessary certificates to the SDP Client and SDP Gateway (in the example case they have a shared volume).

> The content of the `docker-compose.yml` file contains a basic architecture with a single client, controller & gateway.

```bash
$ docker-compose up -d
```

2. Connecting the SDP Gateway and the SDP Controller

```bash
$ docker exec -it sdp-gateway /bin/bash
$ fwknopd -f

---
(sdp_com.c:423) Setting CA cert for peer cert verification.
(sdp_com.c:622) Starting connection attempt 1
(sdp_com.c:371) Connected with TLS_AES_256_GCM_SHA384 encryption
(sdp_com.c:735) Server certificates:
(sdp_com.c:737) Subject: /C=ES/ST=PA/L=A/O=Uniovi/OU=SE/CN=2/emailAddress=abc@xyz.com
(sdp_com.c:740) Issuer: /C=ES/ST=PA/L=A/O=Uniovi/OU=SE/CN=PhD/emailAddress=abc@xyz.com
(sdp_ctrl_client.c:627) Credentials-good message received
(sdp_message.c:258) Received credential update message
(sdp_ctrl_client.c:637) Credential update received
(sdp_ctrl_client.c:1960) All new credentials stored successfully
(sdp_message.c:272) Received service or access data message
(sdp_ctrl_client.c:675) Access data update received
```

3. Request certificates to the SDP Controller from the SDP Client

```bash
$ docker exec -it sdp-client /bin/bash
$ /usr/bin/fwknop -n service_gate

---
(sdp_com.c:423) Setting CA cert for peer cert verification.
(sdp_com.c:622) Starting connection attempt 1
(sdp_com.c:371) Connected with TLS_AES_256_GCM_SHA384 encryption
(sdp_com.c:735) Server certificates:
(sdp_com.c:737) Subject: /C=ES/ST=PA/L=A/O=Uniovi/OU=SE/CN=2/emailAddress=abc@xyz.com
(sdp_com.c:740) Issuer: /C=ES/ST=PA/L=A/O=Uniovi/OU=SE/CN=PhD/emailAddress=abc@xyz.com
(sdp_ctrl_client.c:627) Credentials-good message received
(sdp_message.c:258) Received credential update message
(sdp_ctrl_client.c:637) Credential update received
(sdp_ctrl_client.c:1960) All new credentials stored successfully
```