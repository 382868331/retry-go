package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry012BackoffAttemptOrigin(t *testing.T){ got:=BackOffDelay(1,nil,New(Delay(2*time.Millisecond)).retrierCore); if got != 2*time.Millisecond{t.Fatalf("got %v want %v",got,2*time.Millisecond)} }
