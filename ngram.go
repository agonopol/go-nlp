package nlp

import "strings"

type NGram string 

func NewNGram(words ... string) NGram {
	return NGram(strings.Join(words,"/"))
}

func (this NGram) Each() []string {
	return strings.Split(string(this), "/")
}
