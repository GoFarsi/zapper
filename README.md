# zapper [![Go Reference](https://pkg.go.dev/badge/github.com/GoFarsi/zapper.svg)](https://pkg.go.dev/github.com/GoFarsi/zapper)
zapper is zap but customized with multi core and sentry support, zapper make easiest usage with zap logger.

### Cores
- [x] Console Writer
- [x] Sentry Core
- [x] File Writer
- [x] Json Core

## Install

```shell
$ go get -u github.com/GoFarsi/zapper
```

## Example

- Console writer core

```go
package main

import (
	"github.com/GoFarsi/zapper"
	"log"
)

func main() {
	z := zapper.New(false, zapper.WithTimeFormat(zapper.RFC3339NANO))
	if err := z.NewCore(zapper.ConsoleWriterCore(true)); err != nil {
		log.Fatal(err)
	}

	z.Info("test info")
	z.Debug("debug level")
}
```

- Sentry Core

```go
package main

import (
	"github.com/GoFarsi/zapper"
	"log"
	"os"
)

func main() {
	z := zapper.New(false)
	if err := z.NewCore(zapper.SentryCore(os.Getenv("DSN"), "test", zapper.DEVELOPMENT, nil)); err != nil {
		log.Fatal(err)
	}

	err(z)
}

func err(z zapper.Zapper) {
	z.Error("test error new")
}
```

- File Writer Core

```go
package main

import (
	"github.com/GoFarsi/zapper"
	"log"
)

func main() {
	z := zapper.New(true, zapper.WithDebugLevel())
	if err := z.NewCore(zapper.FileWriterCore("./test_data", nil)); err != nil {
		log.Fatal(err)
	}

	z.Debug("debug log")
	z.Info("info log")
	z.Warn("warn log")
	z.Error("error log")
}
```

- Json Core

```go
package main

import (
	"github.com/GoFarsi/zapper"
	"log"
)

func main() {
	z := zapper.New(true, zapper.WithDebugLevel(), zapper.WithServiceDetails(23, "zapper"))
	if err := z.NewCore(zapper.JsonWriterCore("./test_data", ".json", nil)); err != nil {
		log.Fatal(err)
	}

	z.Debug("debug log")
	z.Info("info log")
	z.Warn("warn log")
	z.Error("error log")
}
```

- Multi Core

```go
package main

import (
	"github.com/GoFarsi/zapper"
	"log"
)

func main() {
	z := zapper.New(false, zapper.WithDebugLevel())
	if err := z.NewCore(
		zapper.ConsoleWriterCore(true),
		zapper.FileWriterCore("./test_data", nil),
	); err != nil {
		log.Fatal(err)
	}

	z.Debug("debug log")
	z.Info("info log")
	z.Warn("warn log")
	z.Error("error log")
}
```

## Contributing

1. Fork zapper repository
2. Clone forked project
3. create new branch from main
4. change things in new branch
5. then send Pull Request from your changes in new branch
