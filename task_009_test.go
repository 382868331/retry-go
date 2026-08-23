package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry009MaxDelayTarget(t *testing.T){ got:=New(MaxDelay(9*time.Millisecond)).maxDelay; if got != 9*time.Millisecond{t.Fatalf("got %v want %v",got,9*time.Millisecond)} }
