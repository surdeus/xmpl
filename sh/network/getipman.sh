#!/bin/sh

# Setting IP you want.
sudo ip addr add 192.168.0.120/24 dev enp3s0

# Setting up IP of gateaway.
sudo ip route add default via 192.168.0.1

