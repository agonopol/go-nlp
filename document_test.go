package nlp

import "testing"
import "bufio"
import "bytes"
// import "fmt"


func TestTokenize(t * testing.T) {
	line := "Unlike regular variable declarations, a short variable declaration may redeclare variables provided they were originally declared in the same block with the same type, and at least one of the non-blank variables is new. As a consequence, redeclaration can only appear in a multi-variable short declaration. Redeclaration does not introduce a new variable; it just assigns a new value to the original."
	buf := bytes.NewBufferString(line)
	r := bufio.NewReader(buf)
	doc := NewDocument(r)
	compare(t, 1, doc.occurances["unlik/regular"])
}

func TestSimilaritySameDoc(t * testing.T) {
	line := "Unlike regular variable declarations, a short variable declaration may redeclare variables provided they were originally declared in the same block with the same type, and at least one of the non-blank variables is new. As a consequence, redeclaration can only appear in a multi-variable short declaration. Redeclaration does not introduce a new variable; it just assigns a new value to the original."
	buf := bytes.NewBufferString(line)
	r := bufio.NewReader(buf)
	doc := NewDocument(r)
	buf = bytes.NewBufferString(line)
	r = bufio.NewReader(buf)
	same := NewDocument(r)
	compare(t, 1.0, doc.Similarity(same))
}

func TestSimilarityDiffDoc(t * testing.T) {
	line := "Unlike regular variable declarations, a short variable declaration may redeclare variables provided they were originally declared in the same block with the same type, and at least one of the non-blank variables is new. As a consequence, redeclaration can only appear in a multi-variable short declaration. Redeclaration does not introduce a new variable; it just assigns a new value to the original."
	buf := bytes.NewBufferString(line)
	r := bufio.NewReader(buf)
	doc := NewDocument(r)
	diff := "Similarly regular variable declarations, a short variable declaration may redeclare variables provided they were originally declared in the same block with the same type, and at least one of the non-blank variables is new. As a consequence, redeclaration can only appear in a multi-variable short declaration. Redeclaration does not introduce a new variable; it just assigns a new value to the original."
	buf = bytes.NewBufferString(diff)
	r = bufio.NewReader(buf)
	same := NewDocument(r)
	compare(t, float64(0.9677), doc.Similarity(same))
}

func TestFromUrl(t *testing.T) {
	FromUrl("http://localhost:6060/pkg/http/")
	// doc,_ := FromUrl("http://localhost:6060/pkg/http/")
	// for ngram, count := range doc.occurances {
		// fmt.Printf("[%v] -> %d\n", ngram, count)
	// }
}