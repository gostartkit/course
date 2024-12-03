# article

a base app for RESTful API

## Gegetting started

### Build from source

```bash
go build -ldflags "-s -w" -buildmode=exe -tags release -o bin/article
```

### Create config

```bash
bin/article config
```

**Note**: Please modify config.json as you need.

### Start service

```bash
bin/article serve
```

**Note**: Default config is use network: `unix`, you can change `network` to `tcp` and `addr` to `127.0.0.1:5000` for test.

### Create models

```bash
gsk model category article tag comment articleTag
```

### Create api code

```bash
gsk code
```

### Build

```bash
gsk build
```