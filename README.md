## About

An application that temporarily stores data as byte arrays.


## Usage

To work with the app, you can use the following endpoints:

- `POST /bin` - create a bin.
  The body should contain the binary data that will be saved.
  It is also important to specify the `Content-Type` header.
  If it is not specified, bytebin will try to determine it automatically.
  If this is not possible, `application/ octet-stream` will be used.
  A unique id will be returned as a response.

- `GET /bin/{id}` - get bin by id.
  The response will be returned as a byte array and a `Content-Type` header associated with this bin.


## Installation

We suggest you deploy your own bytebin using Docker.

The first step is to clone the project:

```bash
git clone https://github.com/orewaee/bytebin.git
cd bytebin
```

After downloading, take a look at the `deploy/compose.yaml` file.
You can customize the Docker Compose config to your liking, or you can start running it right away.

To run bytebin use the following command:

```bash
docker compose -f ./deploy/compose.yaml -p bytebin up -d
```

Or a `sh` script:

```bash
sh scripts/compose.sh
```
