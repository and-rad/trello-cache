package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const url = "https://api.trello.com/1/boards"

// Config stores the configuration loaded from a file in the working directory
type Config struct {
	URL   string `json:"-"`
	Key   string `json:"key"`
	Token string `json:"token"`
	Board string `json:"project"`
}

func Load() *Config {
	var cnf Config

	if data, err := ioutil.ReadFile("config.json"); err != nil {
		panic(err)
	} else if err = json.Unmarshal(data, &cnf); err != nil {
		panic(err)
	}

	cnf.URL = fmt.Sprintf("%s/%s?key=%s&token=%s&fields=name&cards=all&lists=all",
		url, cnf.Board, cnf.Key, cnf.Token)

	return &cnf
}
