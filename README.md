# Waeb

[![Build Status](https://github.com/tomMoulard/traefik-plugin-waeb/actions/workflows/main.yml/badge.svg)](https://github.com/tomMoulard/traefik-plugin-waeb/actions/workflows/main.yml)

Make Traefik a web server !

This is a plugin for [Traefik](https://traefik.io) to build a **web server** as a middleware.

> Not the plugin we deserved, but the plugin we needed.
> 
> ~ [@mpl](https://github.com/mpl) ([src](https://twitter.com/lejatorn/status/1661750793232617477))

> Just because you can, doesn't mean you should
> 
> ~ [@dtomcej](https://github.com/dtomcej) ([src](https://twitter.com/daniel_tomcej/status/1661746210485723136))

## Usage

### Configuration

Here is an example of a file provider dynamic configuration (given here in
YAML), where the interesting part is the `http.middlewares` section:

```yaml
# Dynamic configuration

http:
  routers:
    my-waeb-router:
      rule: host(`waeb.localhost`)
      service: noop@internal # required
      middlewares:
        - traefik-plugin-waeb

  middlewares:
    traefik-plugin-waeb:
      plugin:
        traefik-plugin-waeb:
          root: "/var/www/html/"
```

#### `root`

The `root` parameter is the root directory of the web server.

### Local test

There is a `docker-compose.yml` file to test the plugin locally:

```bash
docker-compose up -d
```

Then, you can go to [http://waeb.localhost](http://waeb.localhost) to see the
result.
