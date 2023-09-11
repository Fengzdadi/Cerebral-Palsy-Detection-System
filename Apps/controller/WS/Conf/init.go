package Conf

import (
	"Cerebral-Palsy-Detection-System/Serializer"
	yaml "gopkg.in/yaml.v3"
	"io/ioutil"
)

func LoadLocales(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		return err
	}
	Serializer.Dictionary = &m
	return nil
}
