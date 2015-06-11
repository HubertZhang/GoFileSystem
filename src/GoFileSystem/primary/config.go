package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Conf struct {
	Primary_ip string `json:"primary"`
	Backup_ip  string `json:"backup"`
	Http_port  string `json:"port"`
}

var conf = new(Conf)

func init_config() error {
	bytes, err := ioutil.ReadFile("./conf/settings.conf")
	if err != nil {
		fmt.Println("Error on opening settings.conf")
		return err
	}

	if err := json.Unmarshal(bytes, &conf); err != nil {
		fmt.Println("settings.conf error")
		return err
	}

	return nil
}
