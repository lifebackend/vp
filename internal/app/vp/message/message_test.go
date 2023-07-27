package message

import (
	"regexp"
	"testing"
)

func Test(t *testing.T) {
	rxp, err := regexp.Compile("([A-Z]+)([0-9]+) ([0-9]{2}):([0-9]{2}) Перевод ([0-9.]+р) от ([0-9а-яА-Я ]+.) Баланс: ([0-9.]+р)")

	if err != nil {
		panic(err)
	}

	if ok := rxp.Match([]byte(`VISA3200 22:08 Перевод 4580р от Екатерина И. Баланс: 20018.96р`)); !ok {
		t.Error(ok)
	}
}
