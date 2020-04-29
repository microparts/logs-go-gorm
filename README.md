This repo is **DEPRECATED**! Use https://github.com/spacetab-io/logs-gorm-go instead.

logs-go-gorm
-----------

Logger fo gorm

## Usage

Initiate new [logger](https://github.com/microparts/logs-go) with prefilled `logs.Config` and use it to initiate new gorm logger

```go
package main

import (
	"time"
	
  "github.com/jinzhu/gorm"
	"github.com/microparts/logs-go"
	"github.com/microparts/logs-go-gorm"
)

func main() {
	// initiate logs-go
	conf := &logs.Config{
		Level:"warn",
		Format: "text",
		Sentry: &logs.SentryConfig{
			Enable: true,
			Stage:"test",
			DSN: "http://dsn.sentry.com",
			ResponseTimeout: 0,
			StackTrace: logs.StackTraceConfig{
				Enable: true,
			},
		},
	}
	
	l, err := logs.NewLogger(conf)
	if err != nil {
		panic(err)
	}
	
	db, _ := gorm.Open("sqlite3", "./db.sqlite")
  db.SetLogger(gormLogger.NewLogger(l))
  db.LogMode(true)
}
```

## Licence

The software is provided under [MIT Licence](LICENCE).
