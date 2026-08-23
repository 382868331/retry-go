package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry017NilRetryIf(t *testing.T){ got:=New(RetryIf(nil)).retryIf != nil; if got != true{t.Fatalf("got %v want %v",got,true)} }
