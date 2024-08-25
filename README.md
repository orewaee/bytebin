## About

An application that temporarily stores data as byte arrays.

All data is stored as a byte array.
To store additional information about bin, json meta-files are used.
The application implements a mechanism to clean up files with expired lifetime or missing bin/meta part.


## Usage

To work with the app, you can use the following endpoints:

- `POST /bin` - create a bin.
  The body should contain the binary data that will be saved.
  It is also important to specify the `Content-Type` header.
  If it is not specified, bytebin will try to determine it automatically.
  If this is not possible, `application/octet-stream` will be used.
  A unique id will be returned as a response.

- `GET /bin/{id}` - get bin by id.
  The response will be returned as a byte array and a `Content-Type` header associated with this bin.


## Installation

### Docker

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

### Binary

This variant requires pre-building the project for your system.

To do this, as with Docker, you must clone the repository and navigate to the project directory.
After that, you need to build the project into a single binary using the command below:

```bash
go build -o bytebin -v cmd/bytebin/main.go
```

If necessary, change the `GOOS` environment variable to get a binary that is compatible with your system.

You can then use the `config/example.yaml` configuration file to configure bytebin.
If no configuration file is specified at startup, environment variables will be used.

The startup command:

```bash
./bytebin --config=config/example.yaml
```
