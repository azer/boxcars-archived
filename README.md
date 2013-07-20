Easy-to-configure Static Web/Reverse Proxy Server in Go

## Install

```bash
$ go get github.com/azer/boxcars/boxcars
```

## Usage

Create a configuration file *(Boxcars will be reloading it automatically on any changes)*:

```json
{
  "*": "/home/you/sites/404",
  "foo.com": "/home/you/sites/foo.com",
  "*.foo.com": "localhost:8000",
  "qux.org": "localhost:8081",
  "w.com": "wikipedia.org"
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
