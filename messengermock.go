package messenger

import (
	"net/http"
)

type MessengerMock struct {
}

func (m *MessengerMock) HandleMessage(f MessageHandler) {
}

func (m *MessengerMock) HandleDelivery(f DeliveryHandler) {
}

func (m *MessengerMock) HandleOptIn(f OptInHandler) {
}

func (m *MessengerMock) HandleRead(f ReadHandler) {
}

func (m *MessengerMock) HandlePostBack(f PostBackHandler) {
}

func (m *MessengerMock) Handler() http.Handler {
	return nil
}

func (m *MessengerMock) ProfileByID(id int64) (Profile, error) {
	return Profile{}, nil
}

func (m *MessengerMock) GreetingSetting(text string) error {
	return nil
}

func (m *MessengerMock) CallToActionsSetting(state string, actions []CallToActionsItem) error {
	return nil
}

func (m *MessengerMock) Response(to int64) ResponseInterface {
	r_mock := &ResponseMock{}
	return r_mock
}

func (m *MessengerMock) Send(to Recipient, messenger string) error {
	return nil
}

func (m *MessengerMock) SendWithReplies(to Recipient, messenger string, replies []QuickReply) error {
	return nil
}

func (m *MessengerMock) Attachment(to Recipient, dataType AttachmentType, url string) error {
	return nil
}
