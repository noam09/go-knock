# go-knock
A port knocking client written in Go.

## Installation
To install go-knock according to your `$GOPATH`:

```console
$ go get github.com/noam09/go-knock
```

## Usage

```console
$ go-knock --help
Go Knock

Usage:
  knock [-d=<ms>] [-u] <host> <port> <port>...
  knock -h | --help
  knock --version

Options:
  -h --help          Show this screen.
  --version          Show version.
  -d, --delay=<ms>   Delay between ports in milliseconds [default: 300].
  -u, --udp          Send UDP packets instead of TCP [default: false].
```

## TODO

* Add UDP support
