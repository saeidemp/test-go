package ghasedakapi

import (
	"net/url"
	"strings"
	"fmt"
)

type Message struct {
	Message    string            `json:"message"`
}
type MessageResult struct {
	*Return
	Entries []Message `json:"entries"`
}
type Return struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
//Send ...
func (m *MessageService) Send(sender string, receptor []string, message string) ([]Message, error) {
	v := url.Values{}
	v.Set("receptor", ToString(receptor))
	v.Set("message", message)
	return m.CreateSend(v)
}

//CreateSend ...
func (m *MessageService) CreateSend(v url.Values) ([]Message, error) {
	res := new(MessageResult)
	err := m.client.Execute("", v, res)
	return res.Message, err
}

//ToString ...
func ToString(i interface{}) string {
	return strings.Trim(strings.Replace(fmt.Sprint(i), " ", ",", -1), "[]")
}
