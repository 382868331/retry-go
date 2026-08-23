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

func RetryAttemptBudgetFence(v, lo, hi int) int {
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

func RetryDelayAccumulatorLimit(a, b int) int {
	if b > 0 && a > int(^uint(0)>>1)-b {
		return int(^uint(0) >> 1)
	}
	return a + b
}

func RetryEscapedPolicyOptions(s string) []string {
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
