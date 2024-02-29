package shelfit

import (
	"math/rand"
	"strconv"
	"testing"
)

func createRangeInput() string {
	return strconv.Itoa(rand.Intn(1000)) + "-" + strconv.Itoa(rand.Intn(1000))
}

var input = createRangeInput()

func BenchmarkParseRangeInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseRangeInt(input)
	}
}

func TestLjust(t *testing.T) {
	test := "hello world"
	if t1 := ljust(test, 4); t1 != "hell" {
		t.Errorf(`ljust(%s, 4) != "hell"`, test)
	}
	if t2 := ljust(test, 20); len(t2) != 20 {
		t.Errorf(`len(ljust(%s, 20)) != 20`, test)
	}
	if t3 := ljust(test, 0); t3 != "" {
		t.Errorf(`ljust(%s, 0) != ""`, test)
	}
}

func TestAddPrefix(t *testing.T) {
	var (
		test     []string
		prefixes []string
		result   []string
	)
	test = []string{"hello", "world"}
	prefixes = []string{"hello", " ", "#", ""}

	for _, pf := range prefixes {
		result = addPrefix(pf, test)
		for i, v := range result {
			if v != pf+test[i] {
				t.Errorf("mismatch at index %d: %s != %s", i, v, pf+test[i])
			}
		}
	}

}
