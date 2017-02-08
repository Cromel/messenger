package messenger

import (
	"image"
	"io"
)

type ResponseMock struct {
}

func (m *ResponseMock) Text(message string) error {
	return nil
}

func (m *ResponseMock) TextWithReplies(message string, replies []QuickReply) error {
	return nil
}

func (m *ResponseMock) Image(im image.Image) error {
	return nil
}

func (m *ResponseMock) Attachment(dataType AttachmentType, uri string) error {
	return nil
}

func (m *ResponseMock) AttachmentData(dataType AttachmentType, filename string, filedata io.Reader) error {
	return nil
}

func (m *ResponseMock) ButtonTemplate(text string, buttons *[]StructuredMessageButton) error {
	return nil
}

func (m *ResponseMock) GenericTemplate(elements *[]StructuredMessageElement) error {
	return nil
}

func (m *ResponseMock) ListTemplate(elements *[]StructuredMessageElement, buttons *[]StructuredMessageButton, style string) error {
	return nil
}
