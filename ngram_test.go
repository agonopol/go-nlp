package nlp

import "testing"

func compare(t *testing.T, expected, actual interface{}) {
	switch tp := expected.(type) {
		case []uint8:
			ea := expected.([]uint8)
			aa := actual.([]uint8)
			compare(t, len(ea), len(aa))
			for i := range ea {
				compare(t, ea[i], aa[i])
			}			
		default:
			if expected != actual {
				t.Errorf("value differs. Expected [%v], actual [%v]", expected, actual)
			}
	}

}

func TestNgramNew(t *testing.T) {
	ngram := NewNGram("Hello","World")
	compare(t, "Hello/World", string(ngram))
}

