package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
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

	err = json.Unmarshal(bytes, &conf)
	if err != nil {
		fmt.Println("settings.conf error")
		return err
	}

	if conf.Primary_ip == conf.Backup_ip {
		port, err := strconv.Atoi(conf.Http_port)
		if err != nil {
			fmt.Println("illegal port number")
			return err
		}
		port = port + 1
		conf.Http_port = strconv.Itoa(port)
	}

	return nil
}
