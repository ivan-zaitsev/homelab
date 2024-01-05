# media

## Configure wireguard
Update wireguard client configuration in `data/wireguard-config/wg_confs/wg0.conf`

Interface:
- PrivateKey
- Address
- DNS

Peer:
- PublicKey
- PresharedKey
- Endpoint

## Install additional plugins
Add additional repository to Jellyfin
```
https://raw.githubusercontent.com/nicknsy/jellyscrub/main/manifest.json
```
