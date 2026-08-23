package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry010MaxJitterTarget(t *testing.T){ got:=New(MaxJitter(6*time.Millisecond)).maxJitter; if got != 6*time.Millisecond{t.Fatalf("got %v want %v",got,6*time.Millisecond)} }
