package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

var result map[string]interface{}

func ReadConfigFile(env string) {
	var (
		err       error
		byteValue []byte
		jsonFile  *os.File
	)

	if jsonFile, err = os.Open("env/appsettings." + env + ".json"); err != nil {
		log.Println(err)
	}

	defer jsonFile.Close()

	if byteValue, err = io.ReadAll(jsonFile); err != nil {
		log.Println(err)
	}

	if err = json.Unmarshal(byteValue, &result); err != nil {
		log.Println(err)
	}
}

func Configure[T comparable](name string) *T {
	var configJson = result[name].(map[string]interface{})
	var byteValue, _ = json.Marshal(configJson)

	var config = new(T)
	if err := json.Unmarshal(byteValue, config); err != nil {
		fmt.Println(err)
	}

	return config
}
