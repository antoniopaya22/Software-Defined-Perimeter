#!/bin/sh

set -e

mkdir -p /etc/ansible/roles

apt install ansible -y
apt install -y unzip
apt install sshpass -y

# ==============> Install egida-role-cis and egida-role-setup
ansible-galaxy install antonioalfa22.egida_role_setup,v2.0.1 --roles-path /etc/ansible/roles
ansible-galaxy install antonioalfa22.egida_role_cis,2.0.2 --roles-path /etc/ansible/roles

# Create egida vars location
mkdir /etc/egida
mkdir /etc/egida/vars

# ==============> Download & install egida
wget https://github.com/egida-kassandra/egida/releases/download/2.0.0/egida.zip
unzip egida.zip
cd build
chmod +x egida
chmod 777 ansible.cfg
mv egida /usr/local/bin/egida
mv hostsgroups /etc/egida/hostsgroups
mv vars_template.yml /etc/egida/vars/vars_template.yml
mv ansible.cfg /etc/ansible/ansible.cfg
echo "[local]" > /etc/ansible/hosts
echo "localhost" >> /etc/ansible/hosts
cd ..
rm -rf build
rm egida.zip