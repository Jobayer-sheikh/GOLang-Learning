package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type ShortMessageConfig struct {
	ApiUrl   string
	ApiKey   string
	SenderId string
}

type NotificationConfig struct {
	ApiUrl    string
	Serverkey string
	SenderId  string
}

type ConfigurationSettings struct {
	ShortMessageConfig   ShortMessageConfig
	NotificationConfig   NotificationConfig
	IsCacheEnable        bool
	IsUniversalOtpEnable bool
}

var result map[string]interface{}

func ReadConfigFile(env string) {
	var (
		err       error
		byteValue []byte
		jsonFile  *os.File
	)

	if jsonFile, err = os.Open("appsettings." + env + ".json"); err != nil {
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

func main() {

	ReadConfigFile("Local")

	var configuration *ConfigurationSettings = Configure[ConfigurationSettings]("ConfigurationSettings")
	fmt.Println(configuration)

	//var (
	//	err       error
	//	byteValue []byte
	//	jsonFile  *os.File
	//	result    map[string]interface{}
	//)
	//
	//if jsonFile, err = os.Open("appsettings.Local.json"); err != nil {
	//	log.Println(err)
	//}
	//
	//defer jsonFile.Close()
	//
	//if byteValue, err = io.ReadAll(jsonFile); err != nil {
	//	log.Println(err)
	//}
	//
	//if err = json.Unmarshal(byteValue, &result); err != nil {
	//	log.Println(err)
	//}
	//
	//var configJson = result["ConfigurationSettings"].(map[string]interface{})
	//fmt.Println(configJson)
	//
	//byteValue, err = json.Marshal(configJson)
	//
	//var config = ConfigurationSettings{}
	//if err = json.Unmarshal(byteValue, &config); err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(config)
}
