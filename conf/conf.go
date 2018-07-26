package conf

import (
	"encoding/json"
	"io/ioutil"
)

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

	cnf.URL = "https://api.trello.com/1/boards/" + cnf.Board +
		"?key=" + cnf.Key +
		"&token=" + cnf.Token +
		"&fields=" + "name" +
		"&card_attachments=" + "cover" +
		"&cards=all&lists=all&labels=all"

	return &cnf
}
