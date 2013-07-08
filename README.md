boxcars is a handy web server with easy static serving and reverse proxy configuration.

## Usage

Create a configuration file:

```json
{
  "foo.com": "/home/you/sites/foo.com",
  "bar.net": "localhost:8080"
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

![](http://i.cloudup.com/rH_0UwNYg1.jpg)
