package gompost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// Function to make regular application/x-www-form-urlencoded POST requests
func Make_form_request(address string, request map[string]string) (map[string]interface{}, error) {
	rtn := make(map[string]interface{})
	u, _ := url.ParseRequestURI(address)
	client := &http.Client{}
	data := url.Values{}
	for key, value := range request {
		data.Add(key, value)
	}
	input = data.Encode()
	r, _ := http.NewRequest("POST", fmt.Sprintf("%v", u), bytes.NewBufferString(input))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(input)))
	resp, err := client.Do(r)
	if err != nil {
		rtn["body"] = resp
		return rtn, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	rtn["body"] = string(body)
	return rtn, nil
}

// Function to make application/json requests
func Make_json_request(address string, request map[string]string) (map[string]interface{}, error) {
	rtn := make(map[string]interface{})
	u, _ := url.ParseRequestURI(address)
	client := &http.Client{}
	i, _ := json.Marshal(request)
	input = string(i)
	r, _ := http.NewRequest("POST", fmt.Sprintf("%v", u), bytes.NewBufferString(input))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(input)))
	resp, err := client.Do(r)
	if err != nil {
		rtn["body"] = resp
		return rtn, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var jsonPackage map[string]interface{}
	err2 := json.Unmarshal([]byte(string(body)), &jsonPackage)
	return jsonPackage, err2
}
