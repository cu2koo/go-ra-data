openssl genrsa -out root.key 2048
openssl req -x509 -new -nodes -key root.key -sha256 -days 365 -out root.crt
openssl req -new -sha256 -nodes -out server.csr -newkey rsa:2048 -keyout server.key
openssl x509 -req -in server.csr -CA root.crt -CAkey root.key -CAcreateserial -out server.crt -days 365 -sha256