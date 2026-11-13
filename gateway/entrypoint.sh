#!/bin/sh
# Configures the SDP gateway: default-drop firewall, credentials from the
# controller-issued bundle, then fwknopd in the foreground.
set -eu

: "${SDP_ID:?SDP_ID must be set}"
: "${PUBLIC_IP:?PUBLIC_IP must be set (gateway address clients send SPA to)}"
CTRL_ADDR="${CTRL_ADDR:-sdp-controller}"
CTRL_PORT="${CTRL_PORT:-5000}"
BUNDLE=/bundle
TEMPLATES=/etc/sdp/templates

for f in ca.crt "$SDP_ID.crt" "$SDP_ID.key" "$SDP_ID.spa_keys"; do
    [ -s "$BUNDLE/$f" ] || { echo "[gateway] Missing $BUNDLE/$f" >&2; exit 1; }
done

# Interface that owns PUBLIC_IP: fwknopd sniffs SPA packets there.
PCAP_INTF="$(ip -o -4 addr show | awk -v ip="$PUBLIC_IP" '{split($4, a, "/")} a[1] == ip {print $2; exit}')"
[ -n "$PCAP_INTF" ] || { echo "[gateway] No interface has address $PUBLIC_IP" >&2; exit 1; }

# ---- firewall: deny by default, fwknopd opens access per authorized client ----
iptables -P INPUT DROP
iptables -P FORWARD DROP
iptables -A INPUT -i lo -j ACCEPT
iptables -A INPUT -m conntrack --ctstate ESTABLISHED,RELATED -j ACCEPT
iptables -A FORWARD -m conntrack --ctstate ESTABLISHED,RELATED -j ACCEPT
# Per-client DNAT and FORWARD rules are added (and expired) by fwknopd.
# Traffic it forwards leaves through the private side: masquerade it so the
# protected service replies through the gateway. (fwknopd's own SNAT support
# is disabled; see ENABLE_IPT_SNAT in fwknopd.conf.tmpl for why.)
for intf in $(ip -o -4 addr show | awk -v pub="$PCAP_INTF" '$2 != "lo" && $2 != pub {print $2}'); do
    iptables -t nat -A POSTROUTING -o "$intf" -j MASQUERADE
done
iptables -S

KEY_BASE64="$(awk '$1 == "KEY_BASE64" {print $2}' "$BUNDLE/$SDP_ID.spa_keys")"
HMAC_KEY_BASE64="$(awk '$1 == "HMAC_KEY_BASE64" {print $2}' "$BUNDLE/$SDP_ID.spa_keys")"

render() {
    sed -e "s|__SDP_ID__|$SDP_ID|g" \
        -e "s|__CTRL_ADDR__|$CTRL_ADDR|g" \
        -e "s|__CTRL_PORT__|$CTRL_PORT|g" \
        -e "s|__PCAP_INTF__|$PCAP_INTF|g" \
        -e "s|__KEY_BASE64__|$KEY_BASE64|g" \
        -e "s|__HMAC_KEY_BASE64__|$HMAC_KEY_BASE64|g" \
        "$1" > "$2"
    chmod 0600 "$2"
}
render "$TEMPLATES/fwknopd.conf.tmpl" /etc/fwknop/fwknopd.conf

# Credentials are copied only on first start: afterwards the control client
# owns them and rewrites them when the controller rotates credentials.
if [ ! -f /etc/fwknop/gate_sdp_ctrl_client.conf ]; then
    install -d -m 0700 /etc/fwknop/certs
    install -m 0600 "$BUNDLE/ca.crt" "$BUNDLE/$SDP_ID.crt" "$BUNDLE/$SDP_ID.key" /etc/fwknop/certs/
    render "$TEMPLATES/sdp_ctrl_client.conf.tmpl" /etc/fwknop/gate_sdp_ctrl_client.conf
fi

echo "[gateway] SDP ID $SDP_ID listening for SPA on $PCAP_INTF ($PUBLIC_IP)"
exec fwknopd -f "$@"
