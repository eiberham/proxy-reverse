# Proxy Reverse

Reverse proxy server

## How to generate a self signed certificate ?

Assuming that ***openssl*** is already installed, run the following command:

	foo@bar:$ openssl req -nodes -new -x509 -keyout server.key -out server.cert


## Installation

Clone the repository

```bash
foo@bar:$ git clone https://github.com/wwleak/proxy-reverse.git
```

Run the server

```bash
foo@bar:$ go run main.go
```

That's it, now the reverse proxy should be running on ***https:// localhost:3000/***

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)

{...❤}
