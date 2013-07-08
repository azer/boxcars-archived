boxcars is a handy web server with easy static serving and reverse proxy configuration.

## Usage

Create a configuration file *(feel free to change later. Boxcars will be reloading it automatically)*:

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

![](http://i.cloudup.com/rH_0UwNYg1.jpg)
