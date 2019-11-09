package main

import (
	"config_module/config"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(config.Get("Port"), reflect.TypeOf(config.Get("Port")))
	fmt.Println(config.Get("DB"), reflect.TypeOf(config.Get("DB")))
	fmt.Println(config.Get("Test"), reflect.TypeOf(config.Get("Test")))
	fmt.Println(config.Get("isVisible"), reflect.TypeOf(config.Get("isVisible")))
}
