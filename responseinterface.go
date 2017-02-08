package messenger

import (
	"image"
	"io"
)

type ResponseInterface interface {
	Text(message string) error
	TextWithReplies(message string, replies []QuickReply) error
	Image(im image.Image) error
	Attachment(dataType AttachmentType, uri string) error
	AttachmentData(dataType AttachmentType, filename string, filedata io.Reader) error
	ButtonTemplate(text string, buttons *[]StructuredMessageButton) error
	GenericTemplate(elements *[]StructuredMessageElement) error
	ListTemplate(elements *[]StructuredMessageElement, buttons *[]StructuredMessageButton, style string) error
}
