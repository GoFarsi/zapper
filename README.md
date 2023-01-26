# zapper
zapper is zap but customized with multi core and sentry support, zapper make easiest usage with zap logger.

### Cores
- [x] Console Writer
- [x] Sentry Core
- [ ] File Writer
- [ ] Json Core

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

## Contributing

1. Fork zapper repository
2. Clone forked project
3. create new branch from main
4. change things in new branch
5. then send Pull Request from your changes in new branch