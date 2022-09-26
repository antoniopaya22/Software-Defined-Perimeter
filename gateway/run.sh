#!/usr/bin/env bash

iptables -P INPUT DROP
iptables -P FORWARD DROP
iptables -S

yes Y  | cp config/access.conf /etc/fwknop/
yes Y  | cp config/fwknopd.conf /etc/fwknop/
yes Y  | cp config/gate_sdp_ctrl_client.conf /etc/fwknop/
yes Y  | cp config/gate.fwknoprc /etc/fwknop/

fwknopd -f

#netcat -l 4444

while true; do
    sleep 100
done