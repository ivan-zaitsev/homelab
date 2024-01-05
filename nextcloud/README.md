# nextcloud

## Intall nextcloud
1. Download docker compose
```
wget https://github.com/nextcloud/all-in-one/blob/main/compose.yaml
```

2. Configure reverse proxy
```
Uncomment APACHE_PORT and APACHE_IP_BINDING
```

## Build custom owntracks image
```
docker build -t owntracks-nextcloud-proxy  -f Dockerfile .
```
