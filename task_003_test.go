package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry003DefaultJitter(t *testing.T){ got:=New().maxJitter; if got != 100*time.Millisecond{t.Fatalf("got %v want %v",got,100*time.Millisecond)} }

func TestTaskRetry003DefaultJitterAdjacent(t *testing.T){
 got:=New(MaxJitter(7*time.Millisecond)).maxJitter
 if got != 7*time.Millisecond{ t.Fatalf("adjacent got %v want %v",got,7*time.Millisecond) }
}
