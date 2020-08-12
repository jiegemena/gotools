package configs

import "testing"

func TestConfig(t *testing.T) {
	Init("")

	v := GetConfig("path")
	if v.(string) != "configdebug.json" {
		t.Fail()
	}

}
