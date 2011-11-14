package nlp

import "testing"
import "math"


func compare(t *testing.T, expected, actual interface{}) {
	switch tp := expected.(type) {
		case []uint8:
			ea := expected.([]uint8)
			aa := actual.([]uint8)
			compare(t, len(ea), len(aa))
			for i := range ea {
				compare(t, ea[i], aa[i])
			}
		case float64:
			ef := expected.(float64)
			af := actual.(float64)
			if math.Fabs(ef - af) > .001 {
				t.Errorf("value differs. Expected [%v], actual [%v]", expected, actual)
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

