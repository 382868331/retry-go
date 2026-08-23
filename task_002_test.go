package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry002DefaultDelay(t *testing.T){ got:=New().delay; if got != 100*time.Millisecond{t.Fatalf("got %v want %v",got,100*time.Millisecond)} }
