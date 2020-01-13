package slack

// RichTextBlock defines data that is used to hold interactive elements.
//
// More Information: https://api.slack.com/reference/messaging/blocks#RichTexts
type RichTextBlock struct {
	Type     MessageBlockType `json:"type"`
	BlockID  string           `json:"block_id,omitempty"`
	Elements BlockElements    `json:"elements"`
}

// BlockType returns the type of the block
func (s RichTextBlock) BlockType() MessageBlockType {
	return s.Type
}

// NewRichTextBlock returns a new instance of an RichText Block
func NewRichTextBlock(blockID string, elements ...BlockElement) *RichTextBlock {
	return &RichTextBlock{
		Type:    MBTRichText,
		BlockID: blockID,
		Elements: BlockElements{
			ElementSet: elements,
		},
	}
}
