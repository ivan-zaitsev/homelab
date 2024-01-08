# reverse-proxy

## Build custom caddy image
```
docker build -t caddy-custom  -f Dockerfile .
```

## Configure caddy
Update caddy configuration in `reverse-proxy/Caddyfile`
- DUCKDNS_TOKEN
- DUCKDNS_SUBDOMAIN
