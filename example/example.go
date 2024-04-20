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
