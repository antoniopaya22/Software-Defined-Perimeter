FROM node:22-bookworm-slim

# openssl: used by the `pem` npm package to issue certificates
# mariadb-client: lets the entrypoint check which SDP IDs already have keys
RUN apt-get update \
 && apt-get install -y --no-install-recommends ca-certificates git openssl mariadb-client \
 && rm -rf /var/lib/apt/lists/*

# Waverley Labs SDP Controller, pinned for reproducible builds
ARG SDPCONTROLLER_REPO=https://github.com/waverleylabs/SDPController.git
ARG SDPCONTROLLER_COMMIT=54fbca8e9b3a2f164c1540324ee9f510141b08b4
RUN git clone "$SDPCONTROLLER_REPO" /opt/SDPController \
 && cd /opt/SDPController \
 && git checkout --quiet "$SDPCONTROLLER_COMMIT" \
 && rm -rf .git \
 && npm install --omit=dev --no-audit --no-fund

COPY config.js /opt/SDPController/config.js
COPY entrypoint.sh /usr/local/bin/entrypoint.sh

# /certs: controller-only PKI (CA key included), generated at runtime.
# /bundles: one read-only credential bundle per gateway/client SDP ID.
RUN mkdir -p /certs /bundles && chown node:node /certs /bundles

USER node
WORKDIR /certs
EXPOSE 5000

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
