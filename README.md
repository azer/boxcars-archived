Easy-to-configure Static Web/Reverse Proxy Server in Go

## Install

```bash
$ go get github.com/azer/boxcars/boxcars
```

## Usage

Create a configuration file *(it'll be auto-loading changes once you start)*  like the below example;

```json
{
  "foo.com": "/home/you/sites/foo.com",
  "*.bar.net": "localhost:8080",
  "qux.org": {
    "/favicon.ico": "sites/qux.org/static/favicon.ico",
    "/logo.jpg": "sites/qux.org/static/logo.jpg",
    "/static": "sites/qux.org/static",
    "*": "localhost:3000"
  },
  "singlepage.com": {
    "/api": "localhost:1337",
    "*": "sites/singlepage.com"
  },
  "*": "/home/you/404.html"
}
```

And start the server:

```bash
$ boxcars config.json
```

To specify the port:

```bash
$ boxcars config.json -port=8001
```

## Logging

To enable logging for specific modules: 

```bash
$ DEBUG=server,sites boxcars config.json
```

To enable very verbose mode (not recommended):

```bash
$ DEBUG=* boxcars config.json
```

## Benchmarks

* [Nginx VS Boxcars](https://gist.github.com/azer/5955772)

![](http://i.cloudup.com/rH_0UwNYg1.jpg)
