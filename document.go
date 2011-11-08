package nlp

import "bufio"
import "bytes"
import "stemmer"

type Document struct {
	occurances map[NGram]int
}

func NewDocument(in *bufio.Reader) *Document {
	this := &Document{make(map[NGram]int)}
	for sentance, err := in.ReadSlice('.'); err == nil; sentance, err = in.ReadSlice('.') {
		for _, ngram := range this.tokenize(sentance[:len(sentance)-1]) {
			if _, found := this.occurances[ngram]; !found {
				this.occurances[ngram] = 0
			}
			this.occurances[ngram] += 1
		}
	}
	return this
}

func (this *Document) tokenize(sentance []byte) []NGram {
	tokens := bytes.Split(sentance, []byte{' '})
	ngrams := make([]NGram, len(tokens) - -1)
	first := stemmer.Stem(tokens[0])
	for i := 1; i < len(tokens); i++ {
		stem := stemmer.Stem(tokens[i])
		ngrams[i-1] = NewNGram(string(first), string(stem))
		first = stem
	}
	return ngrams
}


func (this *Document) Distnace(other * Document) float64 {
	similarity := 0
	total := 0
	for ngram,count := range this.occurances {
		if match, found :=  other.occurances[ngram]; found {
			similarity += min(count, match)
			total += count
		}
	}
	for _,count := range other.occurances {
		total+= count
	}
	return float64(similarity) / (float64(total) / float64(2))
}