boxcars is my attempt to rewrite [door](http://github.com/azer/door) in Go. At this point,
it's only capable of serving static websites. I'll add HTTP proxying soon.

## Usage

Create a configuration file:

```json
{
  "foo.com": "/home/you/sites/foo.com",
  "bar.net": "/home/you/sites/bar.net"
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
