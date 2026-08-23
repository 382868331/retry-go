package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry015CombinedDelaySum(t *testing.T){ got:=CombineDelay(FixedDelay,FixedDelay)(1,nil,New(Delay(3*time.Millisecond)).retrierCore); if got != 6*time.Millisecond{t.Fatalf("got %v want %v",got,6*time.Millisecond)} }
