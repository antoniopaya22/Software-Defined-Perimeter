#!/bin/sh
# Bootstraps the SDP PKI and starts the controller.
#
# Idempotent: the CA and each member's credentials are only generated when
# missing, so restarting the container keeps the keys the gateway and client
# already use. `docker compose down -v` wipes everything for a clean start.
set -eu

CERT_DIR="${CERT_DIR:-/certs}"
BUNDLE_DIR="${BUNDLE_DIR:-/bundles}"
SDP_IDS="${SDP_IDS:-1 2 3}"
CONTROLLER_SDP_ID="${CONTROLLER_SDP_ID:-2}"
CA_SUBJECT="${CA_SUBJECT:-/C=ES/ST=Asturias/L=Oviedo/O=SDP Lab/OU=SDP/CN=SDP Lab CA}"
CA_DAYS="${CA_DAYS:-3650}"
SDPC=/opt/SDPController

: "${CA_KEY_PASSWORD:?CA_KEY_PASSWORD must be set}"
: "${DB_PASSWORD:?DB_PASSWORD must be set}"

cd "$CERT_DIR"
rm -f .ready
umask 077

if [ ! -s ca.key ] || [ ! -s ca.crt ]; then
    echo "[entrypoint] Creating SDP certificate authority"
    openssl genrsa -aes256 -passout env:CA_KEY_PASSWORD -out ca.key 4096
    openssl req -new -x509 -sha256 -days "$CA_DAYS" -key ca.key \
        -passin env:CA_KEY_PASSWORD -subj "$CA_SUBJECT" -out ca.crt
fi

# ask the database whether an SDP ID already has SPA keys stored
has_keys() {
    MYSQL_PWD="$DB_PASSWORD" mariadb --skip-ssl -h "${DB_HOST:-sdp-db}" -u "${DB_USER:-sdp}" \
        -N -B -e "SELECT COUNT(*) FROM sdpid WHERE sdpid = $1 AND encrypt_key IS NOT NULL" \
        "${DB_NAME:-sdp}"
}

for id in $SDP_IDS; do
    case "$id" in
        *[!0-9]*|'') echo "[entrypoint] Invalid SDP ID: $id" >&2; exit 1 ;;
    esac
    if [ -s "$id.crt" ] && [ -s "$id.key" ] && [ -s "$id.spa_keys" ] && [ "$(has_keys "$id")" = "1" ]; then
        echo "[entrypoint] Credentials for SDP ID $id already exist"
        continue
    fi
    echo "[entrypoint] Generating credentials for SDP ID $id"
    rm -f "$id.crt" "$id.key" "$id.spa_keys"
    node "$SDPC/genCredentials.js" "$id"
    if [ ! -s "$id.crt" ] || [ ! -s "$id.key" ] || [ ! -s "$id.spa_keys" ]; then
        echo "[entrypoint] genCredentials.js did not produce files for SDP ID $id" >&2
        exit 1
    fi
done

# Export one bundle per member (cert, key, SPA keys, CA cert) into the shared
# bundles volume. Each gateway/client mounts only its own subdirectory, so it
# never sees other members' keys or the CA private key.
for id in $SDP_IDS; do
    [ "$id" = "$CONTROLLER_SDP_ID" ] && continue
    mkdir -p "$BUNDLE_DIR/$id"
    cp "$id.crt" "$id.key" "$id.spa_keys" ca.crt "$BUNDLE_DIR/$id/"
    chmod 0755 "$BUNDLE_DIR/$id"
    chmod 0644 "$BUNDLE_DIR/$id/"*
done
touch .ready

echo "[entrypoint] Starting SDP Controller"
exec node "$SDPC/sdpController.js"
