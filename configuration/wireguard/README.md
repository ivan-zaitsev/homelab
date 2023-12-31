# wireguard

## Install wireguard server
```
git clone https://github.com/angristan/wireguard-install
cd wireguard-install
chmod +x wireguard-install.sh
./wireguard-install.sh
cat wg0-*.conf
```

## Setup wireguard client

Connect to wireguard server:
```
sudo apt-get install wireguard
sudo nano /etc/wireguard/wg0.conf (Copy the file content from wg0.conf file which was created on wireguard server)
sudo wg-quick up wg0
sudo systemctl enable wg-quick@wg0
```

Add to wireguard client configuration:
```
[Interface]
PostUp = iptables -A FORWARD -p tcp --tcp-flags SYN,RST SYN -j TCPMSS --clamp-mss-to-pmtu
PostDown = iptables -D FORWARD -p tcp --tcp-flags SYN,RST SYN -j TCPMSS --clamp-mss-to-pmtu

[Peer]
PersistentKeepalive = 25
```

Forward server ports to wireguard client:
```
iptables -t nat -A PREROUTING -i eth0 -p tcp --dport 8096 -j DNAT --to-destination CLIENT_IP
```
