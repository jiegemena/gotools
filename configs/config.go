package configs

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/jiegemena/gotools/filetools"
)

var mconfigjson = make(map[string]interface{}, 4)

//读取配置 config.json
func Init(configpath string) error {
	if configpath == "" {
		configpath = "config.json"
	}
	c1, _ := filetools.PathExists(strings.Replace(configpath, ".json", "debug.json", 1))
	if c1 {
		configpath = strings.Replace(configpath, ".json", "debug.json", 1)
	}

	c, err := filetools.ReadAll(configpath)
	if err != nil {
		fmt.Println("%w", err)
		return err
	}

	err = json.Unmarshal(c, &mconfigjson)
	if err != nil {
		fmt.Println("err = ", err)
		return err
	}

	return nil
}

func isNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

//例子（最多三层）cmds:2:cmd || cmds:te1:cmd
func GetConfig(key string) interface{} {
	s1 := strings.Split(key, ":")

	lens1 := len(s1)
	if lens1 == 0 {
		return nil
	}
	v1 := mconfigjson[s1[0]]
	if lens1 == 1 {
		return v1
	}
	if lens1 == 2 {
		if isNum(s1[1]) {
			sint2, _ := strconv.Atoi(s1[1])
			return (v1.([]interface{}))[sint2].(interface{})
		} else {
			return v1.(map[string]interface{})[s1[1]]
		}
	}

	if lens1 == 3 {
		if isNum(s1[1]) {
			sint2, _ := strconv.Atoi(s1[1])
			v2 := (v1.([]interface{}))[sint2].(interface{})
			return v2.(map[string]interface{})[s1[2]]
		} else {
			return v1.(map[string]interface{})[s1[1]].(map[string]interface{})[s1[2]]
		}
	}

	return mconfigjson[key]
}

func GetJsonList(node interface{}) []interface{} {
	return node.([]interface{})
}

func GetJsonMap(node interface{}, key string) interface{} {
	return node.(map[string]interface{})[key]
}
