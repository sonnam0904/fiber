package helpers

import (
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
)

func LoadConfiguration(path ...interface{}) interface{} {
	file, _ := ioutil.ReadFile("./conf/config.json")
	decoder := jsoniter.Get(file, path...).ToString()
	return decoder
}
