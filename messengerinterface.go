package messenger

import (
	"net/http"
)

type MessengerInterface interface {
	HandleMessage(f MessageHandler)
	HandleDelivery(f DeliveryHandler)
	HandleOptIn(f OptInHandler)
	HandleRead(f ReadHandler)
	HandlePostBack(f PostBackHandler)
	Handler() http.Handler
	ProfileByID(id int64) (Profile, error)
	GreetingSetting(text string) error
	CallToActionsSetting(state string, actions []CallToActionsItem) error
	Response(to int64) ResponseInterface
	Send(to Recipient, message string) error
	SendWithReplies(to Recipient, message string, replies []QuickReply) error
	Attachment(to Recipient, dataType AttachmentType, url string) error
}
