package nlp

import "html"
import "io"
import "os"
import "bytes"
import "strings"
import "unicode"



type HTMLScrubbedReader struct {
	node *html.Node
	buf *bytes.Buffer
}

func (this *HTMLScrubbedReader) scrub(data string) string {
	return strings.TrimFunc(data, func(r int) bool {
		return !unicode.IsLetter(r)
	})
}

func (this *HTMLScrubbedReader) parse(node * html.Node, buf *bytes.Buffer) {
	if  node.Type ==  html.TextNode {
		buf.WriteString(node.Data)
	}
	for _,child := range node.Child {
		this.parse(child, buf)
	}
}

func NewHTMLScrubbedReader(reader io.Reader) (*HTMLScrubbedReader,os.Error) {
	node, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBufferString("")
	this := &HTMLScrubbedReader{node,buf}
	this.parse(node, buf)
	return this,nil
}

func (this * HTMLScrubbedReader) Read(p []byte) (int, os.Error) {
	return this.buf.Read(p)
}


