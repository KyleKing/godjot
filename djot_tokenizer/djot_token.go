package djot_tokenizer

import "fmt"

type DjotToken int

const (
	DivClassKey    = "DivClassKey"
	CodeLangKey    = "CodeLangKey"
	InlineMathKey  = "InlineMathKey"
	DisplayMathKey = "DisplayMathKey"
)

const (
	None          = 0
	DocumentBlock = 2*iota + 1
	HeadingBlock
	QuoteBlock
	ListItemBlock
	CodeBlock
	DivBlock
	PipeTableBlock
	ReferenceDefBlock
	FootnoteDefBlock
	ParagraphBlock
	ThematicBreakToken

	Attribute
	Padding

	VerbatimInline
	EscapedSymbolInline
	SpanInline
	ImageSpanInline
	LinkUrlInline
	LinkReferenceInline
	AutolinkInline
	EmphasisInline
	StrongInline
	HighlightedInline
	SubscriptInline
	SuperscriptInline
	InsertInline
	DeleteInline
	FootnoteReferenceInline
	SymbolsInline
	RawFormatInline
	SmartSymbolInline
)

func (t DjotToken) String() string {
	if t == None {
		return "None"
	}
	if t&1 == 0 {
		return (t ^ 1).String() + "Close"
	}
	switch t {
	case None:
		return "None"
	case DocumentBlock:
		return "DocumentBlock"
	case HeadingBlock:
		return "HeadingBlock"
	case QuoteBlock:
		return "QuoteBlock"
	case ListItemBlock:
		return "ListItemBlock"
	case CodeBlock:
		return "CodeBlock"
	case DivBlock:
		return "DivBlock"
	case PipeTableBlock:
		return "PipeTableBlock"
	case ReferenceDefBlock:
		return "ReferenceDefBlock"
	case FootnoteDefBlock:
		return "FootnoteDefBlock"
	case ParagraphBlock:
		return "ParagraphBlock"
	case ThematicBreakToken:
		return "ThematicBreakToken"
	case Attribute:
		return "Attribute"
	case Padding:
		return "Padding"
	case VerbatimInline:
		return "VerbatimInline"
	case EscapedSymbolInline:
		return "EscapedSymbolInline"
	case SpanInline:
		return "SpanInline"
	case ImageSpanInline:
		return "ImageSpanInline"
	case LinkUrlInline:
		return "LinkUrlInline"
	case LinkReferenceInline:
		return "LinkReferenceInline"
	case AutolinkInline:
		return "AutolinkInline"
	case EmphasisInline:
		return "EmphasisInline"
	case StrongInline:
		return "StrongInline"
	case HighlightedInline:
		return "HighlightedInline"
	case SubscriptInline:
		return "SubscriptInline"
	case SuperscriptInline:
		return "SuperscriptInline"
	case InsertInline:
		return "InsertInline"
	case DeleteInline:
		return "DeleteInline"
	case FootnoteReferenceInline:
		return "FootnoteReferenceInline"
	case SymbolsInline:
		return "SymbolsInline"
	case RawFormatInline:
		return "RawFormatInline"
	case SmartSymbolInline:
		return "SmartSymbolInline"
	default:
		panic(fmt.Errorf("unexpected djot token type: %d", t))
	}
}