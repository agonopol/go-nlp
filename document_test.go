package nlp

import "testing"
import "bufio"
import "bytes"
import "fmt"


func TestTokenize(t * testing.T) {
	line := "Unlike regular variable declarations, a short variable declaration may redeclare variables provided they were originally declared in the same block with the same type, and at least one of the non-blank variables is new. As a consequence, redeclaration can only appear in a multi-variable short declaration. Redeclaration does not introduce a new variable; it just assigns a new value to the original."
	buf := bytes.NewBufferString(line)
	r := bufio.NewReader(buf)
	doc := NewDocument(r)
	for ngram, count := range doc.occurances {
		fmt.Printf("[%v]-[%v]\n", ngram, count)
	}
	
	
}