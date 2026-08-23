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

func TestRetryAttemptBudgetFenceRegression(t *testing.T) {
	TestRetryAttemptBudgetFence(t)
	TestRetryAttemptBudgetFence(t)
}

func TestRetryDelayAccumulatorLimit(t *testing.T) {
	m := int(^uint(0) >> 1)
	if got := RetryDelayAccumulatorLimit(m, 1); got != m {
		t.Fatalf("got %d", got)
	}
}
