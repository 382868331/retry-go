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
