How to generate a self signed certificate

Assuming that openssl is already installed, run the following command:

openssl req -nodes -new -x509 -keyout server.key -out server.cert