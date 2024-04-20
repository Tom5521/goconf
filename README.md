# goconf

This is a preferences API inspired by that of [fyne](https://docs.fyne.io/explore/preferences.html)

## Usage/Examples

```go
package main

import (
 "fmt"

 conf "github.com/Tom5521/goconf"
)

func main() {
 settings, err := conf.New("myapp")
 if err != nil {
  panic(err)
 }
 settings.SetString("c1", "value1")

 fmt.Println(settings.String("c1"))
}
```

## Installation

Install my-project with `go get`

```bash
go get -u github.com/Tom5521/goconf@latest
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
