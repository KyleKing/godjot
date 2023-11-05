package djot_parser

import (
	"bytes"
	"fmt"
	"github.com/sivukhin/godjot/djot_tokenizer"
	"github.com/sivukhin/godjot/html_writer"
	"github.com/stretchr/testify/require"
	"html"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"testing"
)

func printTextAst(builder *strings.Builder, trees ...Tree[DjotNode]) {
	for _, tree := range trees {
		switch tree.Type {
		case TextNode:
			builder.Write(tree.Text)
		default:
			printTextAst(builder, tree.Children...)
		}
	}
}

func printDjot(text string) string {
	document := []byte(text)
	tokens := djot_tokenizer.BuildDjotTokens(document)
	context := BuildDjotContext(document, tokens)
	ast := buildDjotAst(document, context, tokens, false)
	builder := html_writer.HtmlWriter{}
	ConvertDjotToHtml(&builder, "html", ast...)
	return builder.String()
}

const examplesDir = "examples"

func TestDownloadExample(t *testing.T) {
	//t.Skip("manual test for download of djot examples from official docs")

	normalize := func(line string) string {
		line = strings.Trim(line, "\r\n\t")
		line = strings.TrimPrefix(line, "<pre><code>")
		line = strings.TrimSuffix(line, "</code></pre>")
		return line
	}

	response, err := http.Get("https://raw.githubusercontent.com/jgm/djot/main/doc/syntax.html")
	require.Nil(t, err)
	docBytes, err := io.ReadAll(response.Body)
	require.Nil(t, err)
	var (
		djotStartToken = []byte(`<div class="djot">`)
		htmlStartToken = []byte(`<div class="html">`)
		endToken       = []byte(`</div>`)
	)
	example := 0
	for {
		djotStart := bytes.Index(docBytes, djotStartToken)
		if djotStart == -1 {
			break
		}
		djotEnd := djotStart + bytes.Index(docBytes[djotStart:], endToken)
		djotExample := html.UnescapeString(normalize(string(docBytes[djotStart+len(djotStartToken) : djotEnd])))
		docBytes = docBytes[djotEnd+len(endToken):]

		htmlStart := bytes.Index(docBytes, htmlStartToken)
		require.NotEqual(t, htmlStart, -1)
		htmlEnd := htmlStart + bytes.Index(docBytes[htmlStart:], endToken)
		htmlExample := html.UnescapeString(normalize(string(docBytes[htmlStart+len(htmlStartToken) : htmlEnd])))
		docBytes = docBytes[htmlEnd+len(endToken):]

		require.Nil(t, os.WriteFile(path.Join(examplesDir, fmt.Sprintf("%02d.html", example)), []byte(htmlExample), 0660))
		require.Nil(t, os.WriteFile(path.Join(examplesDir, fmt.Sprintf("%02d.djot", example)), []byte(djotExample), 0660))
		example++
	}
}

func TestDjotDocExample(t *testing.T) {
	dir, err := os.ReadDir(examplesDir)
	require.Nil(t, err)
	for _, entry := range dir {
		name := entry.Name()
		example, ok := strings.CutSuffix(name, ".html")
		if !ok {
			continue
		}
		htmlExample, err := os.ReadFile(path.Join(examplesDir, fmt.Sprintf("%v.html", example)))
		require.Nil(t, err)
		djotExample, err := os.ReadFile(path.Join(examplesDir, fmt.Sprintf("%v.djot", example)))
		require.Nil(t, err)
		t.Run(string(djotExample), func(t *testing.T) {
			result := printDjot(string(djotExample))
			require.Equalf(
				t, string(htmlExample), result,
				"invalid html (%v != %v), djot tokens: %v",
				string(htmlExample), result,
				djot_tokenizer.BuildDjotTokens(djotExample),
			)
		})
	}
}