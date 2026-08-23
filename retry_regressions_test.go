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
