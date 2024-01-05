#!/bin/bash

ip route del default
ip route add default via 172.20.0.10
ip route add 10.10.10.0/24 via 172.20.0.1
