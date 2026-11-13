#!/bin/sh
# Configures the SDP client (initiating host) from the controller-issued bundle.
set -eu

: "${SDP_ID:?SDP_ID must be set}"
: "${GATEWAY_IP:?GATEWAY_IP must be set}"
: "${SERVICE_IDS:?SERVICE_IDS must be set}"
CTRL_ADDR="${CTRL_ADDR:-sdp-controller}"
CTRL_PORT="${CTRL_PORT:-5000}"
BUNDLE=/bundle
TEMPLATES=/etc/sdp/templates

for f in ca.crt "$SDP_ID.crt" "$SDP_ID.key" "$SDP_ID.spa_keys"; do
    [ -s "$BUNDLE/$f" ] || { echo "[client] Missing $BUNDLE/$f" >&2; exit 1; }
done

# Source address the gateway should authorize: the one used to reach it.
CLIENT_IP="${CLIENT_IP:-$(ip -o route get "$GATEWAY_IP" | sed -n 's/.* src \([0-9.]*\).*/\1/p')}"
[ -n "$CLIENT_IP" ] || { echo "[client] Cannot determine source IP towards $GATEWAY_IP" >&2; exit 1; }

# Only (re)initialize on first start: afterwards the control client owns
# these files and may have rotated the keys in them.
if [ ! -f /root/.fwknoprc ]; then
    install -d -m 0700 /root/certs
    install -m 0600 "$BUNDLE/ca.crt" "$BUNDLE/$SDP_ID.crt" "$BUNDLE/$SDP_ID.key" /root/certs/
    KEY_BASE64="$(awk '$1 == "KEY_BASE64" {print $2}' "$BUNDLE/$SDP_ID.spa_keys")"
    HMAC_KEY_BASE64="$(awk '$1 == "HMAC_KEY_BASE64" {print $2}' "$BUNDLE/$SDP_ID.spa_keys")"
    render() {
        sed -e "s|__SDP_ID__|$SDP_ID|g" \
            -e "s|__CTRL_ADDR__|$CTRL_ADDR|g" \
            -e "s|__CTRL_PORT__|$CTRL_PORT|g" \
            -e "s|__CLIENT_IP__|$CLIENT_IP|g" \
            -e "s|__GATEWAY_IP__|$GATEWAY_IP|g" \
            -e "s|__SERVICE_IDS__|$SERVICE_IDS|g" \
            -e "s|__KEY_BASE64__|$KEY_BASE64|g" \
            -e "s|__HMAC_KEY_BASE64__|$HMAC_KEY_BASE64|g" \
            "$1" > "$2"
        chmod 0600 "$2"
    }
    render "$TEMPLATES/fwknoprc.tmpl" /root/.fwknoprc
    render "$TEMPLATES/sdp_ctrl_client.conf.tmpl" /root/sdp_ctrl_client.conf
fi

echo "[client] SDP ID $SDP_ID ($CLIENT_IP) ready. Request access with: fwknop -n service_gate"
exec "$@"
