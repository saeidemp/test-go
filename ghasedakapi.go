package ghasedakapi

//Kavenegar ...
type Ghasedakapi struct {
	Message *MessageService
}
//MessageService ...
type MessageService struct {
	client *Client
}
//New ...
func New(apikey string) *Ghasedakapi {
	client := NewClient(apikey)
	return NewWithClient(client)
}

//NewWithClient ...
func NewWithClient(client *Client) *Ghasedakapi {
	k := &Kavenegar{}
	k.Message = NewMessageService(client)
	return k
}


//NewMessageService ...
func NewMessageService(client *Client) *MessageService {
	m := &MessageService{client: client}
	return m
}

