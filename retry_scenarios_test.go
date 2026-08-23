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

func TestRetryDelayAccumulatorLimitRegression(t *testing.T) {
	TestRetryDelayAccumulatorLimit(t)
	TestRetryDelayAccumulatorLimit(t)
}

func TestRetryEscapedPolicyOptions(t *testing.T) {
	got := RetryEscapedPolicyOptions("a\\;b;c")
	if !reflect.DeepEqual(got, []string{"a;b", "c"}) {
		t.Fatalf("got %v", got)
	}
}

func TestRetryEscapedPolicyOptionsRegression(t *testing.T) {
	TestRetryEscapedPolicyOptions(t)
	TestRetryEscapedPolicyOptions(t)
}

func TestRetryPreferredErrorOrder(t *testing.T) {
	got := RetryPreferredErrorOrder([]int{3, 1, 3, 2, 1})
	if !reflect.DeepEqual(got, []int{3, 1, 2}) {
		t.Fatalf("got %v", got)
	}
}

func TestRetryPreferredErrorOrderRegression(t *testing.T) {
	TestRetryPreferredErrorOrder(t)
	TestRetryPreferredErrorOrder(t)
}

func TestRetryZeroBackoffGuard(t *testing.T) {
	if got := RetryZeroBackoffGuard([]int{1, 2}, 0); got != nil {
		t.Fatalf("got %v", got)
	}
}

func TestRetryZeroBackoffGuardRegression(t *testing.T) {
	TestRetryZeroBackoffGuard(t)
	TestRetryZeroBackoffGuard(t)
}

func TestRetryUnicodeOperationLabel(t *testing.T) {
	got := RetryUnicodeOperationLabel("A界B", 2)
	if got != "A界" || !utf8.ValidString(got) {
		t.Fatalf("got %q", got)
	}
}

func TestRetryUnicodeOperationLabelRegression(t *testing.T) {
	TestRetryUnicodeOperationLabel(t)
	TestRetryUnicodeOperationLabel(t)
}

func TestRetryRetryableFlagWhitespace(t *testing.T) {
	got, err := RetryRetryableFlagWhitespace(" TRUE ")
	if err != nil || !got {
		t.Fatalf("got %v %v", got, err)
	}
}

func TestRetryRetryableFlagWhitespaceRegression(t *testing.T) {
	TestRetryRetryableFlagWhitespace(t)
	TestRetryRetryableFlagWhitespace(t)
}

func TestRetryExponentialShiftCap(t *testing.T) {
	if got := RetryExponentialShiftCap(2, 100, 4); got != 32 {
		t.Fatalf("got %d", got)
	}
}

func TestRetryExponentialShiftCapRegression(t *testing.T) {
	TestRetryExponentialShiftCap(t)
	TestRetryExponentialShiftCap(t)
}

func TestRetryJitterUpperEndpoint(t *testing.T) {
	if got := RetryJitterUpperEndpoint([]int{1, 2, 3}, 1); got != 3 {
		t.Fatalf("got %d", got)
	}
}

func TestRetryJitterUpperEndpointRegression(t *testing.T) {
	TestRetryJitterUpperEndpoint(t)
	TestRetryJitterUpperEndpoint(t)
}

func TestRetryOptionSnapshotIsolation(t *testing.T) {
	in := map[string]map[string]int{"a": {"x": 1}}
	got := RetryOptionSnapshotIsolation(in)
	got["a"]["x"] = 9
	if in["a"]["x"] != 1 {
		t.Fatalf("input mutated")
	}
}

func TestRetryOptionSnapshotIsolationRegression(t *testing.T) {
	TestRetryOptionSnapshotIsolation(t)
	TestRetryOptionSnapshotIsolation(t)
}

func TestRetryReverseErrorHistory(t *testing.T) {
	if got := RetryReverseErrorHistory("A界🙂"); got != "🙂界A" {
		t.Fatalf("got %q", got)
	}
}
