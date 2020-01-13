package slack

import (
	"strings"
)

const (
	METText            MessageElementType = "text"
	METUser            MessageElementType = "user"
	METBroadcast       MessageElementType = "broadcast"
	METRichTextSection MessageElementType = "rich_text_section"
)

type RichTextSectionBlockElement struct {
	Type     MessageElementType `json:"type"`
	Elements BlockElements      `json:"elements,omitempty"`
}

// ElementType returns the type of the Element
func (s RichTextSectionBlockElement) ElementType() MessageElementType {
	return s.Type
}

// NewRichTextSectionBlockElement returns an instance of a new RichTextSection Block Element
func NewRichTextSectionBlockElement(ty string, actionID string, elements BlockElements, text MessageElementType, userID string, broadcastRange string) *RichTextSectionBlockElement {
	return &RichTextSectionBlockElement{
		Type:     METRichTextSection,
		Elements: elements,
	}
}

type RichTextComponentBlockElement struct {
	Type   MessageElementType `json:"type"`
	Text   MessageElementType `json:"text,omitempty"`
	UserID string             `json:"user_id,omitempty"`
	Range  string             `json:"range,omitempty"`
}

// ElementType returns the type of the Element
func (s RichTextComponentBlockElement) ElementType() MessageElementType {
	return s.Type
}

// NewRichTextComponentBlockElement returns an instance of a new RichTextComponent Block Element
func NewRichTextComponentBlockElement(ty MessageElementType, actionID string, text MessageElementType, userID string, broadcastRange string) *RichTextComponentBlockElement {
	return &RichTextComponentBlockElement{
		Type:   ty,
		Text:   text,
		UserID: userID,
		Range:  broadcastRange,
	}
}

func (e RichTextSectionBlockElement) FindComponentElements(ty MessageElementType) []*RichTextComponentBlockElement {
	result := make([]*RichTextComponentBlockElement, 0)

	for _, el := range e.Elements.ElementSet {
		if el.ElementType() == ty {
			res := el.(*RichTextComponentBlockElement)
			result = append(result, res)
		}
	}
	return result
}

func (e RichTextSectionBlockElement) GetText() string {
	var buff strings.Builder
	if e.IsMention() || e.IsBroadcast() {
		el := e.FindComponentElements(METText)
		for _, e := range el {
			buff.WriteString(string(e.Text))
		}
	}
	return buff.String()
}

func (e RichTextSectionBlockElement) IsMention() bool {
	return e.HasType(METUser)
}

func (e RichTextSectionBlockElement) IsBroadcast() bool {
	return e.HasType(METBroadcast)
}

func (e RichTextSectionBlockElement) HasType(ty MessageElementType) bool {
	el := e.FindComponentElements(ty)
	return len(el) > 0
}
