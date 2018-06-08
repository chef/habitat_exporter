package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func JsonHttpGet(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP code: %d: %s", resp.StatusCode, body)
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}

func HttpGetStatus(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	resp.Body.Close()
	return resp.StatusCode, nil
}
