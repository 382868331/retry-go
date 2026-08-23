package retry

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"unicode/utf8"
)

var (
	_ = context.Background
	_ = errors.Is
	_ = reflect.DeepEqual
	_ = utf8.ValidString
)

func TestRetryAttemptBudgetFence(t *testing.T) {
	if got := RetryAttemptBudgetFence(5, 10, 0); got != 5 {
		t.Fatalf("got %d", got)
	}
}
