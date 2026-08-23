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

func TestRetryNormalizeBounds(t *testing.T) {
	if got := RetryNormalizeBounds(5, 10, 0); got != 5 {
		t.Fatalf("got %d", got)
	}
}

func TestRetryNormalizeBoundsRegression(t *testing.T) {
	TestRetryNormalizeBounds(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestRetryNormalizeBounds(t)
}

func TestRetrySaturatingAdd(t *testing.T) {
	m := int(^uint(0) >> 1)
	if got := RetrySaturatingAdd(m, 1); got != m {
		t.Fatalf("got %d", got)
	}
}

func TestRetrySaturatingAddRegression(t *testing.T) {
	TestRetrySaturatingAdd(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestRetrySaturatingAdd(t)
}

func TestRetrySplitEscapedTokens(t *testing.T) {
	got := RetrySplitEscapedTokens("a\\;b;c")
	if !reflect.DeepEqual(got, []string{"a;b", "c"}) {
		t.Fatalf("got %v", got)
	}
}

func TestRetrySplitEscapedTokensRegression(t *testing.T) {
	TestRetrySplitEscapedTokens(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestRetrySplitEscapedTokens(t)
}

func TestRetryStableUnique(t *testing.T) {
	got := RetryStableUnique([]int{3, 1, 3, 2, 1})
	if !reflect.DeepEqual(got, []int{3, 1, 2}) {
		t.Fatalf("got %v", got)
	}
}

func TestRetryStableUniqueRegression(t *testing.T) {
	TestRetryStableUnique(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestRetryStableUnique(t)
}

func TestRetryPartitionValues(t *testing.T) {
	if got := RetryPartitionValues([]int{1, 2}, 0); got != nil {
		t.Fatalf("got %v", got)
	}
}

func TestRetryPartitionValuesRegression(t *testing.T) {
	TestRetryPartitionValues(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestRetryPartitionValues(t)
}

func TestRetryTruncateLabel(t *testing.T) {
	got := RetryTruncateLabel("A界B", 2)
	if got != "A界" || !utf8.ValidString(got) {
		t.Fatalf("got %q", got)
	}
}

func TestRetryTruncateLabelRegression(t *testing.T) {
	TestRetryTruncateLabel(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestRetryTruncateLabel(t)
}

func TestRetryParseBooleanOption(t *testing.T) {
	got, err := RetryParseBooleanOption(" TRUE ")
	if err != nil || !got {
		t.Fatalf("got %v %v", got, err)
	}
}

func TestRetryParseBooleanOptionRegression(t *testing.T) {
	TestRetryParseBooleanOption(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestRetryParseBooleanOption(t)
}

func TestRetryBoundedBackoff(t *testing.T) {
	if got := RetryBoundedBackoff(2, 100, 4); got != 32 {
		t.Fatalf("got %d", got)
	}
}

func TestRetryBoundedBackoffRegression(t *testing.T) {
	TestRetryBoundedBackoff(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestRetryBoundedBackoff(t)
}

func TestRetrySelectUpperQuantile(t *testing.T) {
	if got := RetrySelectUpperQuantile([]int{1, 2, 3}, 1); got != 3 {
		t.Fatalf("got %d", got)
	}
}

func TestRetrySelectUpperQuantileRegression(t *testing.T) {
	TestRetrySelectUpperQuantile(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestRetrySelectUpperQuantile(t)
}

func TestRetryCloneNestedState(t *testing.T) {
	in := map[string]map[string]int{"a": {"x": 1}}
	got := RetryCloneNestedState(in)
	got["a"]["x"] = 9
	if in["a"]["x"] != 1 {
		t.Fatalf("input mutated")
	}
}

func TestRetryCloneNestedStateRegression(t *testing.T) {
	TestRetryCloneNestedState(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestRetryCloneNestedState(t)
}

func TestRetryReverseUnicodeLabel(t *testing.T) {
	if got := RetryReverseUnicodeLabel("A界🙂"); got != "🙂界A" {
		t.Fatalf("got %q", got)
	}
}

func TestRetryReverseUnicodeLabelRegression(t *testing.T) {
	TestRetryReverseUnicodeLabel(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestRetryReverseUnicodeLabel(t)
}

func TestRetrySlidingWindows(t *testing.T) {
	got := RetrySlidingWindows([]int{1, 2, 3}, 2)
	if len(got) != 2 {
		t.Fatalf("got %v", got)
	}
}

func TestRetrySlidingWindowsRegression(t *testing.T) {
	TestRetrySlidingWindows(t)
	// The public contract remains stable when the regression is exercised repeatedly.
	TestRetrySlidingWindows(t)
}

func TestRetryJoinOptionalParts(t *testing.T) {
	if got := RetryJoinOptionalParts(nil, ","); got != "" {
		t.Fatalf("got %q", got)
	}
}
