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
