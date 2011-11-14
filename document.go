package nlp

import "bufio"
import "bytes"
import "stemmer"
import "unicode"

type Document struct {
	occurances map[NGram]int
}

func NewDocument(in *bufio.Reader) *Document {
	this := &Document{make(map[NGram]int)}
	for sentance, err := in.ReadSlice('.'); err == nil; sentance, err = in.ReadSlice('.') {
		for ngram := range this.tokenize(sentance[:len(sentance)-1]) {
			if _, found := this.occurances[ngram]; !found {
				this.occurances[ngram] = 0
			}
			this.occurances[ngram] += 1
		}
	}
	return this
}

func notletter(rune int) bool {
	return !unicode.IsLetter(rune)
} 

func firstToken( tokens [][]byte) (string,int) {
	for i := 0;i< len(tokens);i++{
		first := stemmer.Stem(bytes.TrimFunc(tokens[i],notletter))
		if !IsStopWord(string(first)) && len(first) >0 {
			return string(first),i
		}
	}
	return "", len(tokens)
}
func (this *Document) tokenize(sentance []byte) <- chan NGram {
	c := make(chan NGram)
	tokens := bytes.Split(sentance, []byte{' '})
	go func() {
		first,i := firstToken(tokens)	
		for j:=i+1;j<len(tokens); j++{
			last := string(stemmer.Stem(bytes.TrimFunc(tokens[j], notletter)))
			if !IsStopWord(last) && len(last) > 0  {
				c <- NewNGram(first, string(last))
				first = string(last)
			}
		}
		close(c)
	}()
	return c
}


func (this *Document) Similarity(other * Document) float64 {
	similarity := 0
	total := 0
	for ngram,count := range this.occurances {
		if match, found :=  other.occurances[ngram]; found {
			similarity += min(count, match)
		}
		total += count
	}
	for _,count := range other.occurances {
		total+= count
	}
	return float64(similarity) / (float64(total) / float64(2))
}