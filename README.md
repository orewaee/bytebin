## About

A small app that works with byte arrays, which allows you to store arbitrary data, not just specific ones.
At the moment there is no limitation, but there are problems with transporting certain types of data.
To correct them, it is necessary to introduce compression and archiving.
Byte arrays are sent directly to memory, allowing them to be quickly retrieved when needed.
It is also possible to configure the maximum size of saved files and their lifetime.


## Usage

To work with beans, two endpoints are implemented:

- `POST /bin` - create a bin
- `GET /bin/{id}` - get data from a bin

There is also `GET /health` to check the life of the app.
