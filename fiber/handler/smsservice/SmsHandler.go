package smsservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (command *SmsCommand) Handle() (bool, error) {

	var (
		url      = Url
		err      error
		request  *http.Request
		response *http.Response
	)

	command.setCredential()

	var body = new(bytes.Buffer)
	err = json.NewEncoder(body).Encode(command)

	if request, err = http.NewRequest("POST", url, body); err != nil {
		return false, err
	}

	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	if response, err = client.Do(request); err != nil {
		return false, err
	}

	var data *string = new(string)
	_ = json.NewDecoder(response.Body).Decode(data)
	fmt.Println(data)

	defer response.Body.Close()

	log.Println(response.StatusCode)

	return true, nil
}

func (command *SmsCommand) setCredential() {
	command.Type = Type
	command.ApiKey = ApiKey
	command.SenderId = SenderId
}
