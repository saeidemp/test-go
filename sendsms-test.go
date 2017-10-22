package ghasedakapi

import (
	"testing"
)

var (
	api *Ghasedakapi
	apiKey string
	receptor string
)

func setup() {
	apiKey="0"
	receptor=""
}

func TestMessageSend(t *testing.T) {
	setup()
	api := New(apiKey)
	sender := ""                 
	receptor := []string{receptor} 
	message := "Hello Go!"      
	if _, err := api.Message.Send(sender, receptor, message); err != nil {
		//t.Errorf("MessageSend failed: %v", err)
	} else {
		t.Logf("MessageSend OK")
	}
}