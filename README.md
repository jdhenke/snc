# `snc`

> netcat for tls connections without any verification

## Installation

Ensure `$GOPATH/bin` is on your `$PATH` and get this repo:

```bash
go get github.com/jdhenke/snc
```

## Usage

```bash
snc $host [$port]
```

## Example

```bash
echo -e "GET / HTTP/1.0\n\n" | snc mail.google.com
```
