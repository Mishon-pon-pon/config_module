package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

//Get - ;;;;
func Get(field string) interface{} {
	var FilePath string = "./config/config.json"
	file, err := os.Open(FilePath)
	if err != nil {
		fmt.Println(err)
	}
	scaner := bufio.NewScanner(file)
	var config interface{}
	var str string
	for scaner.Scan() {
		str += scaner.Text()
	}
	fileByte := []byte(str)
	err = json.Unmarshal(fileByte, &config)
	if err != nil {
		fmt.Println(err)
	}

	conf := config.(map[string]interface{})

	switch reflect.TypeOf(conf[field]) {
	case reflect.TypeOf(new(string)):
		val, is := conf[field].(string)
		if is {
			return val
		}
	case reflect.TypeOf(new(int)):
		val, is := conf[field].(int)
		if is {
			return val
		}
	case reflect.TypeOf(new(float64)):
		val, is := conf[field].(float64)
		if is {
			return val
		}
	case reflect.TypeOf(new(bool)):
		val, is := conf[field].(bool)
		if is {
			return val
		}
	default:
		return conf[field]
	}
	return conf[field]
}
