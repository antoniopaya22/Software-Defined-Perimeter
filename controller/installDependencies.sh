#!/usr/bin/env bash
echo "Doing intial setup..."
apt-get update -y
apt-get install -y curl git mysql-server openssl openssh-server
curl -fsSL https://deb.nodesource.com/setup_lts.x | sudo -E bash -
apt-get install -y nodejs

echo "Dependencies installed!"