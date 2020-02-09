package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// 構造体定義
type DBConfig struct {
	DBMS     string `json:"dbms"`
	User     string `json:"user"`
	Password string `json:"password"`
	Protocol string `json:"protocol"`
	DBName   string `json:"db_name"`
}

func DecodeJson() DBConfig {

	jsonString, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	cf := DBConfig{}
	json.Unmarshal(jsonString, &cf)
	return cf

}
