# xini v0.0.1

> very simple and naive, to be written better
>
> 暂不支持解析行尾注释

## usage

```bash
$ go get -u github.com/binz96/xini
```

- `my.ini`

```ini
; 这是注释
# 这也是注释
; default section
k = v
k1 = v1

; section
[server]
k2 = v2
k3 = v3
```

- `main.go`

```go
package main

import (
	"fmt"
	"log"

	"github.com/binz96/xini"
)

func main() {
    cfg, err := xini.Load("my.ini")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*cfg)
	v := cfg.Section("").Key("k")
	v1 := cfg.Section("").Key("k1")
	v2 := cfg.Section("server").Key("k2")
	v3 := cfg.Section("server").Key("k3")
	fmt.Println(v, v1, v2, v3)
}
```

