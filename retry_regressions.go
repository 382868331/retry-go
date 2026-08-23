package retry

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

var (
	_ = context.Background
	_ = errors.Is
	_ = fmt.Sprintf
	_ = reflect.DeepEqual
	_ = runtime.Gosched
	_ = strings.TrimSpace
	_ sync.Mutex
	_ = time.Millisecond
	_ = utf8.ValidString
)

func RetryNormalizeBounds(v, lo, hi int) int {
	if lo > hi {
		lo, hi = hi, lo
	}
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func RetrySaturatingAdd(a, b int) int {
	if b > 0 && a > int(^uint(0)>>1)-b {
		return int(^uint(0) >> 1)
	}
	return a + b
}

func RetrySplitEscapedTokens(s string) []string {
	out := []string{}
	cur := []rune{}
	esc := false
	for _, r := range s {
		if esc {
			cur = append(cur, r)
			esc = false
		} else if r == 92 {
			esc = true
		} else if r == 59 {
			out = append(out, string(cur))
			cur = nil
		} else {
			cur = append(cur, r)
		}
	}
	out = append(out, string(cur))
	return out
}

func RetryStableUnique(in []int) []int {
	seen := map[int]bool{}
	out := []int{}
	for _, v := range in {
		if !seen[v] {
			seen[v] = true
			out = append(out, v)
		}
	}
	return out
}

func RetryPartitionValues(in []int, size int) [][]int {
	if size <= 0 {
		return nil
	}
	out := [][]int{}
	for len(in) > 0 {
		n := size
		if n > len(in) {
			n = len(in)
		}
		out = append(out, in[:n])
		in = in[n:]
	}
	return out
}

func RetryTruncateLabel(s string, n int) string {
	r := []rune(s)
	if n < 0 {
		return ""
	}
	if n >= len(r) {
		return s
	}
	return string(r[:n])
}

func RetryParseBooleanOption(s string) (bool, error) {
	v := strings.ToLower(strings.TrimSpace(s))
	if v == "true" {
		return true, nil
	}
	if v == "false" {
		return false, nil
	}
	return false, errors.New("invalid boolean")
}

func RetryBoundedBackoff(base int, attempt int, capPow int) int {
	if attempt < 0 {
		return 0
	}
	if attempt > capPow {
		attempt = capPow
	}
	d := base << attempt
	max := base << capPow
	if d > max {
		return max
	}
	return d
}

func RetrySelectUpperQuantile(v []int, p float64) int {
	if len(v) == 0 {
		return 0
	}
	if p <= 0 {
		return v[0]
	}
	if p >= 1 {
		return v[len(v)-1]
	}
	return v[int(p*float64(len(v)-1))]
}

func RetryCloneNestedState(in map[string]map[string]int) map[string]map[string]int {
	out := map[string]map[string]int{}
	for k, m := range in {
		c := map[string]int{}
		for x, v := range m {
			c[x] = v
		}
		out[k] = c
	}
	return out
}

func RetryReverseUnicodeLabel(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func RetrySlidingWindows(in []int, size int) [][]int {
	if size <= 0 || size > len(in) {
		return nil
	}
	out := [][]int{}
	for i := 0; i+size <= len(in); i++ {
		out = append(out, in[i:i+size])
	}
	return out
}
