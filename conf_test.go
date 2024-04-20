package conf_test

import (
	"fmt"
	"strconv"
	"testing"

	conf "github.com/Tom5521/goconf"
)

func TestSet(t *testing.T) {
	settings, err := conf.New("goconf-test")
	if err != nil {
		t.Fail()
	}
	value := 12341243
	settings.SetInt("lol1", value)
	gvalue := settings.Int("lol1")
	fmt.Println(gvalue)
	if value != gvalue {
		t.Fail()
	}
	settings.SetInt("lol1", 3)
	fmt.Println(settings.Int("lol1"))
	for i := range 1000 {
		settings.SetInt(strconv.Itoa(i), i)
	}
}
