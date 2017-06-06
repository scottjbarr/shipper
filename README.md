# shipper

A tiny UDP utility that is used to push data to a UDP server.

Can lose messages in the case of loss of connectivity with a server, but
I see that as a preferable option to potentially forever gaining messages
that can never be sent.

The Go UDP connection does a great job of being friendly in the case of not
being able to talk to the server.

You don't want this process to fail

## Usage

Assuming you have another process writing the STDOUT, you pipe the data into
the `shipper` process.

e.g.

    yourprocess | SERVER=127.0.0.1:5000 shipper

## License

The MIT License (MIT)

Copyright (c) 2017 Scott Barr

See [LICENSE.md](LICENSE.md)
