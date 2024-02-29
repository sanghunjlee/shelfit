package shelfit

import (
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func newUUID() string {
	return uuid.NewString()
}

func contain(val string, arr []string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func addPrefix(prefix string, arr []string) []string {
	if prefix == "" {
		return arr
	}
	var newArr []string
	for _, v := range arr {
		newArr = append(newArr, prefix+v)
	}
	return newArr
}

func ljust(s string, l int) string {
	if l <= 0 {
		return ""
	}
	if l < len(s) {
		return string(s[:l])
	} else {
		return s + strings.Repeat(" ", l-len(s))
	}
}

func correctFieldName(field string, anyStruct interface{}) (bool, string) {
	val := reflect.ValueOf(anyStruct)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		fType := val.Type().Field(i)
		if strings.EqualFold(field, fType.Name) {
			return true, fType.Name
		}
	}
	return false, field
}

func compareStrings(left string, right string) bool {
	if left == right {
		return true
	}
	l := strings.ToLower(left)
	r := strings.ToLower(right)
	return l == r
	/* gr := 0
	i := 0
	for j := 0; j < len(r); j++ {
		frag := r[i:j]
		matched, _ := regexp.MatchString(frag, l[i:])
		if !matched || j == len(r)-1 {
			gr += len(frag)
			i = j
		}
	}
	fmt.Println(gr) */
}

func timeStamp(t time.Time) time.Time {
	yr, mt, dt := t.Date()
	h, m, s := t.Clock()

	return time.Date(yr, mt, dt, h, m, s, 0, t.Location())
}

func parseRangeInt(s string) (ints []int) {
	splits := strings.Split(s, "-")
	if len(splits) == 2 {
		start, _ := strconv.Atoi(splits[0])
		end, _ := strconv.Atoi(splits[1])
		if start > end {
			start, end = end, start
		}
		for i := start; i <= end; i++ {
			ints = append(ints, i)
		}
	}
	return ints
}
