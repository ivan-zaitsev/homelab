# Self-signed certificate

1. Generate CA key private key and certificate
```
openssl genrsa -out home-ca-private-key.pem 4096
openssl req -new -x509 -days 36500 -key home-ca-private-key.pem -out home-ca-certificate.pem -subj "/CN=Home Trust Center"
```

3. Generate private key and certificate request:
```
openssl genrsa -out home-private-key.pem 4096
openssl req -new -key home-private-key.pem -out home-certificate-request.pem -subj "/CN=*.home.lan" -addext "subjectAltName=DNS:home.lan,DNS:*.home.lan"
```

4. Generate and sign certificate using CA and certificate request
```
openssl x509 -req -in home-certificate-request.pem -CA home-ca-certificate.pem -CAkey home-ca-private-key.pem -CAcreateserial -out home-certificate.pem -days 3650 -extfile <(printf "subjectAltName=DNS:home.lan,DNS:*.home.lan")
```
