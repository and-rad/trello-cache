package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"trello-cache/conf"
	"trello-cache/trello"
)

func requestData(cnf *conf.Config) ([]byte, error) {
	resp, err := http.Get(cnf.URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func main() {
	cnf := conf.Load()

	data, err := requestData(cnf)
	if err != nil {
		println(err.Error())
		return
	}

	board := trello.Board{}
	if err := json.Unmarshal(data, &board); err != nil {
		println(err.Error())
	} else if pretty, err := json.MarshalIndent(board, "", "    "); err == nil {
		println(string(pretty))
	} else {
		println(string(data))
	}
}
