[![Build Status](https://travis-ci.org/Antonito/gfile.svg?branch=master)](https://travis-ci.org/Antonito/gfile)
[![Go Report Card](https://goreportcard.com/badge/github.com/LongMarch7/gfile)](https://goreportcard.com/report/github.com/LongMarch7/gfile)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/5888662aebd54d2681f9a737dfd33913)](https://www.codacy.com/app/Antonito/gfile?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=Antonito/gfile&amp;utm_campaign=Badge_Grade)
[![Coverage Status](https://coveralls.io/repos/github/Antonito/gfile/badge.svg?branch=master)](https://coveralls.io/github/Antonito/gfile?branch=master)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# GFile

gfile is a WebRTC based file exchange software.

It allows to share a file directly between two computers, without the need of a third party.

![ezgif-5-9936f8008e4d](https://user-images.githubusercontent.com/11705040/55694519-686e2d80-5969-11e9-9bc1-f7a59b62732f.gif)

## Note

This project is still in its early stage.

## Usage

### Sender

```bash
gfile send --file filename
```

-   Run the command
-   A base64 encoded SDP will appear, send it to the remote client
-   Follow the instructions to send the client's SDP to your process
-   The file transfer should start

### Receiver

```bash
# SDP being the base64 SDP gotten from the other client
echo "$SDP" | gfile receive -o filename
```

-   Pipe the other client's SDP to gfile
-   A base64 encoded SDP will appear, send it to the remote client
-   The file transfer should start

### Benchmark

`gfile` is able to benchmark the network speed between 2 clients (1 _master_ and 1 _slave_) with the `bench` command.
For detailed instructions, see `Sender` and `Receiver` instructions.

```bash
# Run as 'master'
gfile bench -m

# Run as 'slave'
echo "$SDP" |  gfile bench
```

### Web Interface

A web interface is being developed via WebAssembly. It is currently **not** working.

### Debug

In order to obtain a more verbose output, it is possible to define the logging level via the `GFILE_LOG` environment variable.

> Example: `export GFILE_LOG="TRACE"`
> See function `setupLogger` in  `main.go` for more information

## Contributors

-   Antoine Baché ([https://github.com/LongMarch7](https://github.com/LongMarch7)) **Original author**

Special thanks to [Sean DuBois](https://github.com/Sean-Der) for his help with [pion/webrtc](https://github.com/pion/webrtc) and [Yutaka Takeda](https://github.com/enobufs) for his work on [pion/sctp](https://github.com/pion/sctp)
