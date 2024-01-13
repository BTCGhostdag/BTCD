# BTCDctl

BTCDctl is an RPC client for BTCD

## Requirements

Go 1.19 or later.

## Installation

#### Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
```

- Run the following commands to obtain and install BTCD including all dependencies:

```bash
$ git clone https://github.com/BTCGhostdag/BTCD/
$ cd BTCD/cmd/BTCDctl
$ go install .
```

- BTCDctl should now be installed in `$(go env GOPATH)/bin`. If you did not already add the bin directory to your
  system path during Go installation, you are encouraged to do so now.

## Usage

The full kaspctl configuration options can be seen with:

```bash
$ kaspctl --help
```

But the minimum configuration needed to run it is:

```bash
$ BTCDctl <REQUEST_JSON>
```

For example:

```
$ BTCDctl '{"getBlockDagInfoRequest":{}}'
```

For a list of all available requests check out the [RPC documentation](infrastructure/network/netadapter/server/grpcserver/protowire/rpc.md)