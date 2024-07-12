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
