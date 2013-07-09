handy web server with easy static serving and reverse proxy configuration.

## Install

```bash
$ go install github.com/azer/boxcars/boxcars
```

## Usage

Create a configuration file *(Boxcars will be reloading it automatically on any changes)*:

```json
{
  "foo.com": "/home/you/sites/foo.com",
  "bar.net": "localhost:8080",
  "qux.org": "localhost:8081"
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

To enable very verbose mode:

```bash
$ DEBUG=* boxcars config.json
```

Less verbose:

```bash
$ DEBUG=server,sites boxcars config.json
```

![](http://i.cloudup.com/rH_0UwNYg1.jpg)
