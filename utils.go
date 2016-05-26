package utils

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
)

func Dump(params interface{}, exit ...bool) {
	atype := reflect.TypeOf(params)
	fmt.Printf("(%+v)%+v\n", atype, params)
	for _, exit := range exit {
		if exit == true {
			os.Exit(1)
		}
	}
}

func IsIp(ip string) bool {
	reg, err := regexp.Compile("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}")
	if err != nil {
		log.Fatal(err)
	}
	var ips []byte
	ips = []byte(ip)
	ret := reg.Match(ips)
	return ret
}

func IsNumberic(param string) (bool, int) {
	if a, err := strconv.Atoi(param); err == nil {
		return true, a
	}
	return false, 0
}
