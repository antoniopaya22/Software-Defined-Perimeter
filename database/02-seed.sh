#!/bin/bash
# Example data for the lab topology in docker-compose.yml.
# Sourced by the MariaDB entrypoint after 01-schema.sql, on first start only.
#
#   SDP ID 1  gateway     -> protects service 1
#   SDP ID 2  controller
#   SDP ID 3  client      -> allowed to use service 1
#   service 1 "protected-web": gateway TCP ${PROTECTED_PORT} is forwarded
#             (DNAT) to ${PROTECTED_SERVICE_IP}:80 on the private network
#
# SPA keys (encrypt_key / hmac_key) are left NULL: the controller generates
# them on first start (controller/entrypoint.sh).

: "${PROTECTED_SERVICE_IP:?PROTECTED_SERVICE_IP must be set}"
: "${PROTECTED_PORT:?PROTECTED_PORT must be set}"

docker_process_sql --database=sdp <<-EOSQL
	INSERT INTO sdpid (sdpid, valid, type, country, state, locality, org, org_unit, alt_name, email, serial)
	VALUES
	  (1, 1, 'gateway',    'ES', 'Asturias', 'Oviedo', 'SDP Lab', 'SDP', 'gateway',    'gateway@sdp.test',    'gw-0001'),
	  (2, 1, 'controller', 'ES', 'Asturias', 'Oviedo', 'SDP Lab', 'SDP', 'controller', 'controller@sdp.test', 'ctrl-0001'),
	  (3, 1, 'client',     'ES', 'Asturias', 'Oviedo', 'SDP Lab', 'SDP', 'client',     'client@sdp.test',     'cli-0001');

	INSERT INTO service (id, name, description)
	VALUES (1, 'protected-web', 'Demo HTTP service hidden behind the SDP gateway');

	INSERT INTO service_gateway (id, service_id, gateway_sdpid, protocol, port, nat_ip, nat_port)
	VALUES (1, 1, 1, 'tcp', ${PROTECTED_PORT}, '${PROTECTED_SERVICE_IP}', 80);

	INSERT INTO sdpid_service (id, sdpid, service_id)
	VALUES (1, 3, 1);
EOSQL
