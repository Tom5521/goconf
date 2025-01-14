package conf_test

import (
	"fmt"
	"strconv"
	"testing"

	conf "github.com/Tom5521/goconf"
)

var settings *conf.Preferences

func TestSet(t *testing.T) {
	var err error
	settings, err = conf.New("goconf-test")
	if err != nil {
		t.Log(err)
		t.FailNow()
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
	for i := 0; i < 1000; i++ {
		settings.SetInt(strconv.Itoa(i), i)
	}
}

func TestGet(t *testing.T) {
	expectedValue := 3
	value := settings.Int("lol1")
	if expectedValue != value {
		t.Fail()
	}
}

func TestCreateNewFields(t *testing.T) {
	fields := []conf.Field{
		{"MEOW", ""},
		{"MEOW2", 0},
	}

	settings.CreateNewFields(true, fields...)

	v1 := settings.String("MEOW")
	if v1 != "" {
		t.Fail()
	}
	v2 := settings.Int("MEOW2")
	if v2 != 0 {
		t.Fail()
	}
}
